# gopenfusion

A toy fusionfall packet protocol implementation written in Go.

## Generating structures

Dump and decompile the `Assembly - CSharp.dll` assembly from the fusionfall main.unity3d, using a tool like [ilspycmd](https://www.nuget.org/packages/ilspycmd/). The structures can then be stripped from the source and passed to the `genstructs.py` script located in `tools/`. See the script for details on usage.