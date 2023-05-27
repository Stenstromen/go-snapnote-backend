#FROM golang:1.20-alpine as build
#WORKDIR /
#COPY *.go ./
#COPY *.mod ./
#COPY *.sum ./
#RUN go build -o /go-readthenburn-backend
#
#FROM alpine:latest
#COPY --from=build /go-readthenburn-backend /
#EXPOSE 8080
#CMD [ "/go-readthenburn-backend" ]

FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o /go-readthenburn-backend

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go-readthenburn-backend /app/
CMD ["/app/go-readthenburn-backend"]
