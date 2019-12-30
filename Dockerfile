FROM golang:1.13.1 as gobuild
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
COPY . /work
WORKDIR /work
RUN go build -o rads .

FROM alpine:latest

RUN mkdir -p /work /secret
COPY --from=gobuild /work/rads /work
WORKDIR /work
CMD ["/work/rads", "-f" ,"/secret"]