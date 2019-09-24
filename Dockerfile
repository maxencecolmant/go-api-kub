
FROM golang:latest as build-env
USER root
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v ./...
RUN go build -o /go/bin/goapi


FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/goapi /
EXPOSE 8080
CMD ["/goapi"]