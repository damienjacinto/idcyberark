# Overview

Projet to handle the idconnection of the request between a pod and the api PAM.

Each connexion to PAM needs a unique Id by credential.
Each Jenkins has his own credential.

The goal of this api is to handle a counter by Jenkins. From an instance of Jenkins you can have multiple requests simultaneously and the Id provided needs to be unique.

# API

- GET /id/{jenkins}

Get the increment value of the unique id for the Jenkins {jenkins}

``` bash
$ curl localhost:8080/id/pic-eul
{"id":1}

$ curl localhost:8080/id/pic-eul
{"id":2}

$ curl localhost:8080/id/pic-dosn
{"id":1}

$ curl localhost:8080/id/pic-eul
{"id":3}
```

- GET /health

healthcheck probe

- GEt /ready

readyness probe

- GET /metrics

prometheus metrics


# Build

To build the app you need to install go and define your $GOPATH ($GOPATH/src/idcyberark should to be valid path).

To manage the dependencies this project use [dep](https://github.com/golang/dep)

```bash
# to install dep on debian plateforms
$ sudo apt-get install go-dep

# to download the dependencies of the project
$ cd $GOPATH/src/idcyberark
$ dep ensure -vendor-only
```

Check usage to build and run idcyberark

## Makefile

Run *make help* to see build/test options

```
$ make help

Usage:

    make <target>

  Targets:

    clean       Remove $GOPATH/idcyberark
    build       Build app in $GOPATH/bin
    run         Run app on port: 8000 - check $(PORT)
    test        Launch test
    container   Build the docker image
    drun        Run docker image on binded port: 8000 - check $(PORT)
```

# Usage

You can configure the run with two variables PORT and MAXCOUNTER.
Default value for PORT is 8000 and MAXCOUNTER is set to 100 by default at runtime.

MAXCOUNTER define the maximum value of the range for the counter (1 ..MAXCOUNTER) 

```bash
# Exemple
$ docker run -p 8080:8080 -e PORT=8080 -e MAXCOUNTER=10 -d jacintod/idcyberark:0.0.1  
```
