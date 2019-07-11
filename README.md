# Overview

Cyberark offer an API with PAM to request hosts data in the safe (host type, user, pwd).
To request informations you need a credential and a unique idconnexion (Number between 1 and 100).
The request lasts around 1 to 5 seconds, during this time any request with the same credential and idconnexion locks the account for 15 minutes.

The goal of this projet is to handle unique idconnection number for each request.
The request to the cyberark's api is launch from a pod that run on Jenkins inside our OpenShift.
Each instance of Jenkins use it's own credential.
From one Jenkins multiple requests can be send simultaneously, the idconnexion provided for each request needs to be unique.
We can't run more than 10 simultaneous pod and PAM connexions on a Jenkins.

# API Entrypoints

- GET /id/{jenkins}

Get the value to use as idconnexion for the Jenkins named :{jenkins}

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

- GET /ready

readiness probe

- GET /metrics

prometheus metrics


# Build

To build the app you need to install go and define your $GOPATH ($GOPATH/src/idcyberark should be the path to the source code)

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

    clean       Remove $GOPATH/bin/idcyberark
    build       Build app in $GOPATH/bin
    run         Run app on port: 8000 - check $(PORT)
    test        Launch test
    container   Build the docker image
    drun        Run docker image on binded port: 8000 - check $(PORT)
```

# Usage

You can configure the run with two variables PORT and MAXCOUNTER.
Default value for PORT is 8000 and MAXCOUNTER is set to 100 by default at runtime.

MAXCOUNTER define the maximum value of the counter's range (1 ..MAXCOUNTER) 

```bash
# Exemple
$ docker run -p 8080:8080 -e PORT=8080 -e MAXCOUNTER=10 -d jacintod/idcyberark:0.0.1  
```
