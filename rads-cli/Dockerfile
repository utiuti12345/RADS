FROM golang:1.13.1 as gobuild
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
COPY . /work
WORKDIR /work
RUN go build -o rads-cli .

FROM busybox:latest

#RUN apk --no-cache add tzdata && \
#    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
#    apk del tzdata

RUN mkdir -p /work /secret
COPY --from=gobuild /work/rads-cli /work
COPY /shell /work
WORKDIR /work
CMD ["/work/rads-cli","-h"]