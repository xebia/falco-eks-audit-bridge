# Go related commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test ./...
GOGET=$(GOCMD) get -u -v
IMAGE=xebia/falco-eks-audit-bridge

# Detect the os so that we can build proper statically linked binary
OS := $(shell uname -s | awk '{print tolower($$0)}')

# Get tag or a short hash of the git for building images.
TAG = $$(git describe --tags --always)

# Name of actual binary to create
BINARY = falco-eks-audit-bridge

# GOARCH tells go build which arch. to use while building a statically linked executable
GOARCH = amd64

# Setup the -ldflags option for go build here.
# While statically linking we want to inject version related information into the binary
LDFLAGS = -ldflags='-extldflags "-static"'

.PHONY: run
run: bin #this will cause "bin" target to be build first
	./$(BINARY)-$(OS)-$(GOARCH) # Execute the binary

# bin creates a platform specific statically linked binary. Platform sepcific because if you are on
# OS-X; linux binary will not work.
.PHONY: bin
bin:
	env CGO_ENABLED=0 GOOS=$(OS) GOARCH=${GOARCH} go build -a -installsuffix cgo ${LDFLAGS} -o ${BINARY}-$(OS)-${GOARCH} . ;

# Create a docker image with the binary embedded
.PHONY: docker
docker:
	docker build -t $(IMAGE):$(TAG) .

# Push pushes the image to the docker repository.
.PHONY: push
push: docker
	docker push $(IMAGE):$(TAG)

# Remove the binary.
.SILENT: clean
.PHONY: clean
clean:
	$(GOCLEAN)
	@rm -f ${BINARY}-$(OS)-${GOARCH}
