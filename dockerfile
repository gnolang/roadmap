FROM golang:1.19-alpine

RUN go install moul.io/depviz/v3/cmd/depviz@latest

ENTRYPOINT ["/go/bin/depviz"]
