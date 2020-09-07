# tart by building the application.
# Build em gocep com distroless
FROM golang:1.14.1 as builder
WORKDIR /go/src/gocep
COPY . .
ENV GO111MODULE=on
RUN GOOS=linux go build -trimpath -ldflags="-s -w" -o gocep main.go
RUN cp gocep /go/bin/gocep

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/gocep /
CMD ["/gocep"]