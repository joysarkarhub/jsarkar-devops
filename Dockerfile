FROM alpine:latest
# RUN adduser -S -D -H -h /app appuser
# USER appuser
# COPY --from=builder /build/main /app/
# WORKDIR /app
RUN mkdir -p /go/bin
ADD jsarkar-devops /go/bin
WORKDIR /go/bin
ENTRYPOINT /go/bin/jsarkar-devops