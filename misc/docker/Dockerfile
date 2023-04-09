FROM golang:1.20-alpine
RUN apk add make
ADD ./ ./
RUN cd ./misc/deps; make install
