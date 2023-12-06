# gopenfusion
![Unit Tests](https://github.com/CPunch/gopenfusion/actions/workflows/tests.yaml/badge.svg)

A toy implementation of the [Fusionfall Packet Protocol](https://openpunk.com/pages/fusionfall-openfusion/) (see: `cnet/`) and accompanying services, written in Go.

## Landwalker demo

An implementation of a landwalker server is located in `login/` && `shard/`. This includes a functional login server and a dummy shard (supporting the minimum amount of packets necessary). There is minimal support for NPCs, and minimal support for player interaction (chat & player movement being mostly it).

Startup the environment using

```sh
$ chmod +x ./build.sh && ./build.sh
$ docker compose up
```

The environment consists of a shard service, login service, redis && postgres containers. redis is used to pass login metadata between the login and shard services, while postgres is just used to store player accounts and characters.

login server is hosted at `127.0.0.1:23000`, just join from your [favorite client](https://github.com/OpenFusionProject/OpenFusion/releases/latest)

## Generating structures

Dump and decompile the `Assembly - CSharp.dll` assembly from the fusionfall main.unity3d, using a tool like [ilspycmd](https://www.nuget.org/packages/ilspycmd/). The full output source can then be passed to `genstructs.py` script located in `tools/`, which will handle scraping constants and calculating structure padding. See the script for details on usage.
