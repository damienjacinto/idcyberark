GONAME?=$(shell basename "$(PWD)")
# app
PORT?=8080
# prometheus
PPORT?=9090
# grafana
GPORT?=3000
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
GOBIN?=$(GOPATH)/bin
CONTAINER_IMAGE?=jacintod/${GONAME}

.PHONY: help clean test

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
	@ echo '    compose     Run docker-compose (idcyberark on port $(PORT) / prometheus on port $(PPORT) / grafan on port $(GPORT))'
	@ echo ''

clean:
	rm -f ${GOBIN}/${GONAME}

build: clean
	go build \
		-ldflags '-s -w -X "${GONAME}/version.Release=${RELEASE}" -X "${GONAME}/version.Commit=${COMMIT}"' \
		-o ${GOBIN}/${GONAME}

run: build
	PORT=${PORT} ${GOBIN}/${GONAME}

test:
	go test -v -race ./...

container:
	docker stop $(CONTAINER_IMAGE):$(RELEASE) || true && \
	docker build \
	--build-arg VERSION=$(RELEASE) \
	--build-arg COMMIT=$(COMMIT) \
	--build-arg APP=$(GONAME) \
	 -t $(CONTAINER_IMAGE):$(RELEASE) .

drun: container
	docker run --name ${GONAME} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(CONTAINER_IMAGE):$(RELEASE)

compose: container
	echo "PORT=${PORT}\nPROMETHEUS_PORT=${PPORT}\nGRAFANA_PORT=${GPORT}\nCONTAINER_IMAGE=$(CONTAINER_IMAGE)\nRELEASE=$(RELEASE)" > .env
	sed -e "s/PORT/${PORT}/" $(PWD)/prometheus/prometheus_template.yml > $(PWD)/prometheus/prometheus.yml
	sed -e "s/PORT/${PPORT}/" $(PWD)/grafana/datasource_template.yml > $(PWD)/grafana/provisioning/datasources/datasource.yml
	docker-compose up