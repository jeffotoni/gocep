FROM golang:1.20 as builder
WORKDIR /go/src/gocep
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o gocep main.go
RUN cp gocep /go/bin/gocep

FROM scratch
ENV TZ=America/Sao_Paulo
COPY --from=builder /go/bin/gocep .
ENTRYPOINT ["./gocep"]