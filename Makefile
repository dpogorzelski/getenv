VERSION=$(shell git describe)
DIST=dist

all: linux-amd64

linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $(DIST)/amd64/getenv
	zip $(DIST)/getenv-$(VERSION)-linux-amd64.zip $(DIST)/amd64/getenv