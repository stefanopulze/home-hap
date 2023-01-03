# Home-hap
Bridge for my home device to homekit world 

## Installation
You can use `config.sample.yaml` as template for the real `config.yml`

## Configuration
Configuration file `config.yml` is loaded from the current directory.

If you want change the configuration file in a custom folder you can
use the cli option `config`

```bash
./homekit-bridge -config=<path-to-config>
```

## Build
**Local**
Compile the application in a single binary inside a `dist` folder
```bash
./scripts/build.sh
```

**Docker**
Use docker to compile the application and create a busy image only with the application binary
as self-contained

```bash
./scripts/docker-build.sh
```