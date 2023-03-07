#!/bin/python
'''
    genstructs.py - gopenfusion

    Takes raw structures from a decompiled 'Assembly - CSharp.dll' from a main.unity3d fusionfall beta client,
    and transpiles them to gopenfusion's custom packet structure & tags. This requires a compiler installed,
    since struct field padding is grabbed via the `offsetof()` C macro. Some manual rearranging of structures
    from the disassembled source might be needed. This script can also be used to generate c-style structures
    (because it already does!)

    usage: ./genstructs.py [IN.cs] > out.go
'''
from distutils.ccompiler import new_compiler
import subprocess
import sys
import os

PACK_ALIGN = 4

def sanitizeName(name: str) -> str:
    # all exported fields in go must start capitalized
    return name[0:2].upper() + name[2:]

def writeToFile(source: str, filePath: str) -> None:
    with open(filePath, "w") as out:
        out.write(source)

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
            elif type == "int[]":
                self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                self.type = "[%d]int32" % self.size
                self.ctype = "int"
                self.cname += "[%d]" % self.size
                self.size *= 4
            elif type == "short[]":
                self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                self.type = "[%d]int16" % self.size
                self.ctype = "short"
                self.cname += "[%d]" % self.size
                self.size *= 2
            elif type == "long[]":
                self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                self.type = "[%d]int64" % self.size
                self.ctype = "long"
                self.cname += "[%d]" % self.size
                self.size *= 8
            elif type == "float[]":
                self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                self.type = "[%d]float32" % self.size
                self.ctype = "float"
                self.cname += "[%d]" % self.size
                self.size *= 4
            elif type == "string":
                # all strings in fusionfall are utf16, in a uint16 array
                self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                self.type = "string"
                self.tags += "size:\"%d\"" % self.size
                self.ctype = "short"
                self.cname += "[%d]" % self.size
                self.size *= 2
            else:
                # assume it's a structure that will be defined later
                if type.find("[]") != -1:
                    type = type.replace("[]", "")
                    self.size = int(marshal[(marshal.find("SizeConst = ") + len("SizeConst = ")):marshal.find(")]")])
                    self.cname = name + "[%d]" % self.size
                else:
                    self.cname = name
                    self.size = 1
                self.name = name
                self.type = sanitizeName(type)
                self.ctype = sanitizeName(type)
                self.needsPatching = True

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
        return self.lines[self.cursor]

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
        source = "type " + self.name + " struct {"
        currentSize = 0
        for field in self.fields:
            currentSize += field.size

            if field.padding > 0:
                currentSize += field.padding
                field.tags += " pad:\"%d\"" % field.padding

            source += "\n\t" + field.name + " " + field.type
            if len(field.tags) > 0:
                source += " `" + field.tags + "`"

        if currentSize != self.size:
            source += "\n// WARNING: computed size is %d, needs to be %d!!" % (currentSize, self.size)
        else:
            source += "\n// SIZE: %d" % self.size
        source += "\n}\n"
        return source


if __name__ == '__main__':
    inFilePath = sys.argv[1]

    iFile = open(inFilePath, "r")
    lines = iFile.readlines()
    structs: list[StructTranspiler] = []
    while True:
        try:
            struct = StructTranspiler(lines)
            struct.parseStructure()
            structs.append(struct)
            lines = struct.lines[struct.cursor:]
        except:
            break

    for i in range(len(structs)):
        struct = structs[i]
        f = struct.needsPatching()
        while f != -1:
            name = struct.fields[f].type
            for s in structs:
                if s.name == name:
                    if struct.fields[f].size > 1:
                        structs[i].fields[f].type = ("[%d]" % struct.fields[f].size) + struct.fields[f].type
                    structs[i].fields[f].size *= s.size
                    structs[i].fields[f].needsPatching = False

            f = struct.needsPatching()

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

    lines = subprocess.getoutput(['./tmp']).splitlines()
    for i in range(len(structs)):
        structs[i].populatePadding(lines[i].split(" "))
    
    for struct in structs:
        print(struct.toGoStyle())

    os.remove("tmp")
    os.remove("tmp.c")
    os.remove("tmp.o")


