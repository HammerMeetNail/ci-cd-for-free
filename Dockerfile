FROM golang:1.13.7 AS Test

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get golang.org/x/tools/cmd/cover &&\
    go get github.com/mattn/goveralls

RUN go test ./... -v -covermode=count -coverprofile=coverage.out
RUN go install

FROM alpine:3.9 as Deploy
COPY --from=Test /go/bin/app /app
ENTRYPOINT ["./app"]
