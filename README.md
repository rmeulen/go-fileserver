[![Go Report Card](https://goreportcard.com/badge/github.com/rmeulen/go-fileserver)](https://goreportcard.com/report/github.com/rmeulen/go-fileserver)
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
