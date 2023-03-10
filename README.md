# gopenfusion

A toy implementation of the [Fusionfall Packet Protocol](https://openpunk.com/pages/fusionfall-openfusion/) written in Go.

## Login Sever

An example login server implementation exists in `server/`. This implementation should be compatible with existing OpenFusion databases, however this only exists as an example and doesn't direct clients to a shard server (they're softlocked after the tutorial, or during character selection).

## Generating structures

Dump and decompile the `Assembly - CSharp.dll` assembly from the fusionfall main.unity3d, using a tool like [ilspycmd](https://www.nuget.org/packages/ilspycmd/). The full output source can then be passed to `genstructs.py` script located in `tools/`, which will handle scraping constants and calculating structure padding. See the script for details on usage.