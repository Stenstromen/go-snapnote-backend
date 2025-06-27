FROM golang:1.24-alpine as builder
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags='-w -s -extldflags "-static"' -o /go-snapnote-backend

FROM scratch
COPY --from=builder /go-snapnote-backend /
USER 65534:65534
EXPOSE 8080
CMD ["/go-snapnote-backend"]