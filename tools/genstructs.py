#!/bin/python
'''
    genstructs.py - gopenfusion

    Takes raw structures from a decompiled 'Assembly - CSharp.dll' from a main.unity3d fusionfall beta client,
    and transpiles them to gopenfusion's custom packet structure & tags. Some useful constants are also grabbed
    and transpiled. This requires a C compiler installed, since struct field padding is grabbed via the `offsetof()`
    C macro. Some manual rearranging of structures from the disassembled source might be needed, however the script
    should 'just werk' off of ilspycmd output. This script can also be modified to generate c-style structures
    (because it already does!)

    usage: ./genstructs.py [IN.cs] > structs.go

    It's recommended to run structs.go through `gofmt`.
'''
from distutils.ccompiler import new_compiler
import subprocess
import sys
import os

PACK_ALIGN = 4

def sanitizeName(name: str) -> str:
    # all exported fields in go must start capitalized
    return name[0:1].upper() + name[1:]

def writeToFile(source: str, filePath: str) -> None:
    with open(filePath, "w") as out:
        out.write(source)

WARN_INVALID = False

class StructTranspiler:
    class StructField:
        def __init__(self, name: str, type: str, marshal: str) -> None:
            self.marshal = marshal
            self.type = type
            self.ctype = type # for transpilation to c
            self.tags = ""
            self.size = 0
            self.padding = 0
            self.name = sanitizeName(name)
            self.cname = self.name
            self.needsPatching = False

            if type == "byte":
                self.type = "uint8"
                self.ctype = "char"
                self.size = 1
            elif type == "sbyte":
                self.type = "int8"
                self.ctype = "char"
                self.size = 1
            elif type == "short":
                self.type = "int16"
                self.ctype = "short"
                self.size = 2
            elif type == "int":
                self.type = "int32"
                self.ctype = "int"
                self.size = 4
            elif type == "uint":
                self.type = "uint32"
                self.ctype = "int"
                self.size = 4
            elif type == "float":
                self.type = "float32"
                self.ctype = "float"
                self.size = 4
            elif type == "long":
                self.type = "int64"
                self.ctype = "long"
                self.size = 8
            elif type == "ulong":
                self.type = "uint64"
                self.ctype = "long"
                self.size = 8
            elif type == "byte[]":
                self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                self.type = "[%d]byte" % self.size
                self.ctype = "char"
                self.cname += "[%d]" % self.size
            elif type == "short[]":
                self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                self.type = "[%d]int16" % self.size
                self.ctype = "short"
                self.cname += "[%d]" % self.size
                self.size *= 2
            elif type == "string":
                # all strings in fusionfall are utf16, in a uint16 array
                self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                self.type = "string"
                self.addTag("size:\"%d\"" % self.size) # special tag to tell our decoder/encoder the size of the uint16[] array
                self.ctype = "short"
                self.cname += "[%d]" % self.size
                self.size *= 2
            elif type == "int[]":
                self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                self.type = "[%d]int32" % self.size
                self.ctype = "int"
                self.cname += "[%d]" % self.size
                self.size *= 4
            elif type == "float[]":
                self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                self.type = "[%d]float32" % self.size
                self.ctype = "float"
                self.cname += "[%d]" % self.size
                self.size *= 4
            elif type == "long[]":
                self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                self.type = "[%d]int64" % self.size
                self.ctype = "long"
                self.cname += "[%d]" % self.size
                self.size *= 8
            else:
                # assume it's a structure that will be defined later
                if type.find("[]") != -1: # it's an array!
                    type = type.replace("[]", "")
                    self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                    self.cname = self.name + "[%d]" % self.size
                else:
                    self.cname = self.name
                    self.size = 1
                self.type = sanitizeName(type)
                self.ctype = sanitizeName(type)
                self.needsPatching = True

        def addTag(self, tag: str) -> None:
            if len(self.tags) > 0: # if there's already a tag defined, make sure there's a space separating them
                self.tags += " "

            self.tags += tag

    def __init__(self, lines: list[str]) -> None:
        self.name = "UnknownStruct"
        self.size = 0
        self.fields: list[self.StructField] = []
        self.lines = lines
        self.cursor = 0

    def getLine(self) -> str:
        return self.lines[self.cursor]

    def getNextLine(self) -> str:
        self.cursor += 1
        return self.getLine()

    def needsPatching(self) -> int:
        for i in range(len(self.fields)):
            if self.fields[i].needsPatching:
                return i
        return -1

    # skip lines until we run into a '[StructLayout('
    def searchStructLayout(self) -> None:
        line = self.getLine()
        while True:
            if line.find("[StructLayout(") != -1 and line.find("Size = ") != -1:
                self.size = int(line[(line.find("Size = ") + len("Size = ")):line.find(")]")])
                line = self.getNextLine()
                self.name = sanitizeName(line[(line.find("public struct ") + len("public struct ")):len(line)-1])
                self.getNextLine()
                return
            line = self.getNextLine()

    # skip lines until we find a struct field
    def searchMarshal(self) -> StructField:
        line = self.getLine()
        while line.find("}") == -1:
            if line.find("[MarshalAs(") != -1:
                marshal = line
                line = self.getNextLine()

                typeStart = line.find("public ") + len("public ")
                typeEnd = line.find(" ", typeStart)
                type = line[typeStart:typeEnd]

                name = line[typeEnd+1:line.find(";", typeEnd)]
                return self.StructField(name, type, marshal)
            line = self.getNextLine()

        return None

    def toCStyle(self) -> str:
        source = "typedef struct {\n"
        for field in self.fields:
            source += "\t%s %s;\n" % (field.ctype, field.cname)
        source += "} %s;\n" % self.name

        for field in self.fields:
            source += "printf(\"%d \""
            source += ", offsetof(%s, %s));\n" % (self.name, field.name)

        source += "printf(\"\\n\");\n"
        return source

    def populatePadding(self, offsets: list[str]) -> None:
        currentSize = 0
        for i in range(len(self.fields)):
            currentSize += self.fields[i].size
            if i < len(self.fields)-1 and currentSize != int(offsets[i+1]):
                self.fields[i].padding = int(offsets[i+1]) - currentSize
                currentSize += self.fields[i].padding

        if currentSize < self.size:
            self.fields[len(self.fields)-1].padding = self.size - currentSize

    def parseStructure(self) -> None:
        self.searchStructLayout()

        line = self.getLine()
        while line.find("}") == -1:
            field = self.searchMarshal()
            if field != None:
                self.fields.append(field)

            line = self.getLine()

    def toGoStyle(self) -> str:
        global WARN_INVALID

        source = "type " + self.name + " struct {"
        currentSize = 0
        for field in self.fields:
            currentSize += field.size

            if field.padding > 0:
                currentSize += field.padding
                field.addTag("pad:\"%d\"" % field.padding)

            source += "\n\t" + field.name + " " + field.type
            if len(field.tags) > 0:
                source += " `" + field.tags + "`"

        if currentSize != self.size:
            source += "\n// WARNING: computed size is %d, needs to be %d!!" % (currentSize, self.size)
            WARN_INVALID = True
        else:
            source += "\n// SIZE: %d" % self.size
        source += "\n}\n"
        return source

def transpileStructs(lines: list[str]) -> list[StructTranspiler]:
    structs: list[StructTranspiler] = []

    def tryPatching(struct: StructTranspiler) -> StructTranspiler:
        f = struct.needsPatching()
        lastf = None
        while f != -1 and lastf != f:
            # search for existing struct
            name = struct.fields[f].type
            for s in structs:
                if s.name == name:
                    if struct.fields[f].size > 1: # was it an array?
                        struct.fields[f].type = ("[%d]" % struct.fields[f].size) + struct.fields[f].type
                    struct.fields[f].size *= s.size # field's size was set to 1 even if it wasn't an array
                    struct.fields[f].needsPatching = False # mark done
                    break

            lastf = f
            f = struct.needsPatching()

        return struct

    while True:
        try:
            struct = StructTranspiler(lines)
            struct.parseStructure()

            # try to resolve patches right here
            struct = tryPatching(struct)

            structs.append(struct)
            lines = struct.lines[struct.cursor:]
        except:
            break

    # move all structures that need patching to the end (so hopefully they'll be in the
    # right order for our C source generation)
    def sortStruct(s: StructTranspiler):
        if s.needsPatching() != -1:
            return 1
        return 0
    structs.sort(key=sortStruct)

    # check for undefined types in structures and patch the sizes
    for i in range(len(structs)):
        structs[i] = tryPatching(structs[i])

    # we compile a small c program to grab the exact offsets and alignments
    source = "#include <stdio.h>\n#include <stddef.h>\n"
    source += "int main() {\n"
    source += "#pragma pack(push)\n#pragma pack(%d)\n" % PACK_ALIGN
    for struct in structs:
        source += struct.toCStyle()
    source += "#pragma pack(pop)\nreturn 0;\n}\n"

    writeToFile(source, "tmp.c")
    compiler = new_compiler()
    compiler.compile(['tmp.c'])
    compiler.link_executable(['tmp.o'], 'tmp')

    # patch padding bytes
    lines = subprocess.getoutput(['./tmp']).splitlines()
    for i in range(len(structs)):
        structs[i].populatePadding(lines[i].split(" "))

    return structs

class ConstTranspiler:
    def __init__(self, lines: list[str]) -> None:
        self.lines = lines
        self.cursor = 0
        self.value = 0
        self.name = 0

    def getLine(self) -> str:
        return self.lines[self.cursor]

    def getNextLine(self) -> str:
        self.cursor += 1
        return self.getLine()

    # skip lines until we run into a 'public const uint '
    def searchForConst(self) -> None:
        line = self.getLine()
        while True:
            if line.find("public const uint ") != -1:
                nameStart = (line.find("public const uint ") + len("public const uint "))
                self.value = int(line[(line.find("= ", nameStart) + len("= ")):line.find("u;")])
                self.name = line[nameStart:line.find(" = ", nameStart)]
                self.getNextLine()
                return
            line = self.getNextLine()

def transpileConsts(lines: list[str]) -> list[ConstTranspiler]:
    consts: list[ConstTranspiler] = []

    while True:
        try:
            const = ConstTranspiler(lines)
            const.searchForConst()

            consts.append(const)
            lines = const.lines[const.cursor:]
        except:
            break

    return consts

if __name__ == '__main__':
    inFilePath = sys.argv[1]

    iFile = open(inFilePath, "r")
    lines = iFile.readlines()
    structs = transpileStructs(lines)
    consts = transpileConsts(lines)

    # emit structures
    source = "// generated via genstructs.py"
    if WARN_INVALID:
        source += " - WARN!! Not all structures are valid, grep 'WARNING'\n"
    else:
        source += " - All structure padding and member alignment verified\n"

    source += "package protocol\n\nconst (\n"
    for const in consts:
        source += "\t%s = 0x%x\n" % (const.name, const.value)
    source += ")\n\n"

    for struct in structs:
        source += struct.toGoStyle() + "\n"
    print(source)

    os.remove("tmp")
    os.remove("tmp.c")
    os.remove("tmp.o")

