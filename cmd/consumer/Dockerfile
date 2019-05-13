FROM alpine:latest
RUN apk add --no-cache \
        libc6-compat
COPY . .
CMD ["./web-consumer"]
EXPOSE 8083