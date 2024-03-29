FROM golang:1.11-alpine as build

ARG VERSION
ARG COMMIT
ARG APP

RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

COPY Gopkg.lock Gopkg.toml /go/src/${APP}/
WORKDIR /go/src/${APP}/

RUN dep ensure -vendor-only

COPY . /go/src/${APP}/
RUN CGO_ENABLED=0 go build \
    -ldflags "-s -w -X ${APP}/version.Release=${VERSION} -X ${APP}/version.Commit=${COMMIT}" \
    -o /bin/app

FROM scratch
ENV PORT 8000
ENV MAXCOUNTER 0
EXPOSE $PORT
COPY --from=build /bin/app /app

CMD ["/app"]