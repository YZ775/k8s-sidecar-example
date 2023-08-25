FROM ghcr.io/cybozu/golang:1.20-jammy AS builder

WORKDIR /work
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o sidecar main.go


FROM scratch
COPY --from=builder /work/sidecar /sidecar
ENTRYPOINT ["/sidecar"]
