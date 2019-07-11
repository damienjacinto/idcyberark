GONAME?=$(shell basename "$(PWD)")
PORT?=8000
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GOBIN?=$(GOPATH)/bin
CONTAINER_IMAGE?=jacintod/${GONAME}

.PHONY: clean test help

help:
	@ echo
	@ echo '  Usage:'
	@ echo ''
	@ echo '    make <target>'
	@ echo ''
	@ echo '  Targets:'
	@ echo ''
	@ echo '    clean       Remove $(GOBIN)/$(GONAME)'
	@ echo '    build       Build app in $(GOBIN)'
	@ echo '    run         Run app on port: $(PORT) - check $$(PORT)'
	@ echo '    test        Launch test'
	@ echo '    container   Build the docker image'
	@ echo '    drun        Run docker image on binded port: $(PORT) - check $$(PORT)'
	@ echo ''

clean:
	rm -f ${GOBIN}/${GONAME}

build: clean
	go build \
		-ldflags '-s -w -X "${GONAME}/version.Release=${RELEASE}" -X "${GONAME}/version.Commit=${COMMIT}" -X "${GONAME}/version.BuildTime=${BUILD_TIME}"' \
		-o ${GOBIN}/${GONAME}

run: build
	PORT=${PORT} ${GOBIN}/${GONAME}

test:
	go test -v -race ./...

container:
	docker stop $(CONTAINER_IMAGE):$(RELEASE) || true && \
	docker rm $(CONTAINER_IMAGE):$(RELEASE) || true && \
	docker build \
	--build-arg VERSION=$(RELEASE) \
	--build-arg APP=$(GONAME) \
	 -t $(CONTAINER_IMAGE):$(RELEASE) .

drun: container
	docker run --name ${GONAME} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(CONTAINER_IMAGE):$(RELEASE)