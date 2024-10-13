FROM golang:1.23-alpine AS build-env
ARG version
ADD . /src/server
WORKDIR /src/server
RUN go mod vendor
RUN go build -ldflags "-s -w -X 'main.version=${version}'" -o server ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY .env /app/
COPY --from=build-env /src/server/server /app/

ENTRYPOINT ["/app/proxy-venture"]
