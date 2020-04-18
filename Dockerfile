# tart by building the application.
# Build em gocep com distroless
FROM golang:1.14.1 as builder

WORKDIR /go/src/gocep

COPY gocep .

ENV GO111MODULE=on

#RUN go install -v ./...
#RUN GOOS=linux go  build -ldflags="-s -w" -o gocep main.go
RUN cp gocep /go/bin/gocep

RUN ls -lh

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/gocep /
CMD ["/gocep"]