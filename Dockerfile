FROM golang:alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o falco-eks-audit-bridge . && \
    apk add -U --no-cache ca-certificates

FROM scratch

COPY --from=builder /build/falco-eks-audit-bridge /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

CMD ["./falco-eks-audit-bridge"]
