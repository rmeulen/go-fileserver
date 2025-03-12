[![Go Report Card](https://goreportcard.com/badge/github.com/rmeulen/go-fileserver)](https://goreportcard.com/report/github.com/rmeulen/go-fileserver)
[![Build and Push Docker Image](https://github.com/rmeulen/go-fileserver/actions/workflows/docker-image.yml/badge.svg)](https://github.com/rmeulen/go-fileserver/actions/workflows/docker-image.yml)
# go-fileserver
A simple fileserver written in Go (1.12)

## Building source
```
go build -o fileserver main.go
```

## Configuration
|Name     |Default|Description                                 |
|---------|-------|--------------------------------------------|
|PORT     |`8080` |Port the fileserver is listening on         |
|FILE_ROOT|`./` |File root of the fileserver on the machine  |

### Viper Configuration
The application uses Viper to manage configuration settings. Viper allows you to read configuration from environment variables, configuration files, and command-line flags.

#### Environment Variables
Viper automatically reads environment variables. You can set the following environment variables to override the default configuration:
```
export PORT=9090
export FILE_ROOT=/path/to/files
```

#### Configuration File
You can also use a configuration file in JSON, TOML, YAML, HCL, or Java properties format. Create a configuration file (e.g., `config.yaml`) with the following content:
```yaml
PORT: 9090
FILE_ROOT: /path/to/files
```
Then, set the `VIPER_CONFIG` environment variable to the path of the configuration file:
```
export VIPER_CONFIG=/path/to/config.yaml
```

## Usage: Application
You can run the application as-is. No flags or parameters are required.
```
./fileserver
```
## Usage: Docker
* Build the docker image:
```
docker build -t <tag> .
```
* Run the container. Use a `volume` to share data on your machine.

```
docker run --rm -p 8080:8080 -v ${PWD}:/root <tag>
```
