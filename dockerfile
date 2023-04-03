FROM golang:1.19-alpine

ARG github_token

RUN go install moul.io/depviz/v3/cmd/depviz@latest
RUN /go/bin/depviz fetch -github-token ${github_token} gnolang/roadmap

CMD /go/bin/depviz gen json gnolang/roadmap
