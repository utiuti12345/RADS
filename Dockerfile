FROM golang:1.13.1 as gobuild
COPY . /work
WORKDIR /work
RUN go build -o rads .

FROM alpine:latest

RUN mkdir -p /work/
COPY --from=gobuild /work/rads /work
WORKDIR /work
CMD ["/work/rads"]