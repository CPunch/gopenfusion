# gopenfusion

A toy implementation of the [Fusionfall Packet Protocol](https://openpunk.com/pages/fusionfall-openfusion/) written in Go.

## Generating structures

Dump and decompile the `Assembly - CSharp.dll` assembly from the fusionfall main.unity3d, using a tool like [ilspycmd](https://www.nuget.org/packages/ilspycmd/). The full output source can then be passed to `genstructs.py` script located in `tools/`, which will handle scraping constants and calculating structure padding. See the script for details on usage.