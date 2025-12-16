FROM golang:1.24-alpine as builder
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags='-w -s -extldflags "-static"' -o /go-snapnote-backend

FROM alpine:latest
RUN apk add --no-cache strace
COPY --from=builder /go-snapnote-backend /go-snapnote-backend
EXPOSE 8080
ENTRYPOINT ["strace", "-f", "-e", "trace=all", "-o", "/tmp/strace.log", "/go-snapnote-backend"]

