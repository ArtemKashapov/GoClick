# syntax=docker/dockerfile:1
FROM golang:onbuild

RUN go mod init
RUN go mod tidy
