FROM golang:alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o falco-eks-audit-bridge .

FROM scratch

COPY --from=builder /build/falco-eks-audit-bridge /app/
WORKDIR /app
CMD ["./falco-eks-audit-bridge"]
