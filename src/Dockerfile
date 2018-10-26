FROM golang:alpine as builder

RUN apk add --no-cache git
RUN go get -u github.com/gorilla/mux

ENV SRC=/go/src/github.com/aystream/time-rest-service/src/
RUN mkdir -p '${SRC}'
ADD . ${SRC}


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/time-rest-service github.com/aystream/time-rest-service/src


FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/time-rest-service /opt/time-rest-service

EXPOSE 8080

ENV NAME "Alexander Yakovlev"
ENTRYPOINT /opt/time-rest-service