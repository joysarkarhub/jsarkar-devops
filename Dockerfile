FROM alpine:latest
RUN mkdir -p /go/bin
ADD jsarkar-devops /go/bin
ADD script.sh /go/bin
WORKDIR /go/bin
ENTRYPOINT /go/bin/jsarkar-devops