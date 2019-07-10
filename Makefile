GONAME?=$(shell basename "$(PWD)")
PORT?=8000
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GOBIN?=$(GOPATH)/bin
CONTAINER_IMAGE?=docker.io/damienjacinto/${GONAME}

.PHONY: clean test

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