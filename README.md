# gopenfusion

A toy implementation of the [Fusionfall Packet Protocol](https://openpunk.com/pages/fusionfall-openfusion/) written in Go.

## Landwalker demo

An implementation of a landwalker server is located in `login/` && `shard/`. This includes a functional login server and a dummy shard (supporting the minimum amount of packets necessary). The DB implementation in `core/db/` matches the OpenFusion 1.4 SQLite tables, which the login server uses. There's no support for NPCs nor other players, and is liable to softlock the client.

## Generating structures

Dump and decompile the `Assembly - CSharp.dll` assembly from the fusionfall main.unity3d, using a tool like [ilspycmd](https://www.nuget.org/packages/ilspycmd/). The full output source can then be passed to `genstructs.py` script located in `tools/`, which will handle scraping constants and calculating structure padding. See the script for details on usage.