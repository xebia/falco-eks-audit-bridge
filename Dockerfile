FROM golang:alpine AS builder

WORKDIR /app

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

COPY . .

RUN go mod download
    # Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o falco-eks-audit-bridge -v

FROM gcr.io/distroless/base

COPY --from=builder /app/falco-eks-audit-bridge /app/falco-eks-audit-bridge

# Run the falco-eks-audit-bridge binary.
ENTRYPOINT ["/app/falco-eks-audit-bridge"]

