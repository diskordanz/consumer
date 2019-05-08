FROM alpine:latest
RUN apk add --no-cache \
        libc6-compat
COPY . .
CMD ["./consumer"]
EXPOSE 8080