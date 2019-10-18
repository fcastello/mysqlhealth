# Go parameters
VERSION=0.0.1
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mysqlhealth

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...

pkg-linux: build-linux
	mkdir -p tmp/mysqlhealth
	cp README.md tmp/mysqlhealth
	cp LICENSE tmp/mysqlhealth
	cp mysqlhealth tmp/mysqlhealth
	cd tmp && COPYFILE_DISABLE=1 tar zcvf mysqlhealth_$(VERSION)_amd64.tar.gz mysqlhealth && mv mysqlhealth_$(VERSION)_amd64.tar.gz ../ && cd ..
	rm -rf tmp

pkg-darwin: build-darwin
	mkdir -p tmp/mysqlhealth
	cp README.md tmp/mysqlhealth
	cp LICENSE tmp/mysqlhealth
	cp mysqlhealth tmp/mysqlhealth
	cd tmp && COPYFILE_DISABLE=1 tar zcvf mysqlhealth_$(VERSION)_darwin_amd64.tar.gz mysqlhealth && mv mysqlhealth_$(VERSION)_darwin_amd64.tar.gz ../ && cd ..
	rm -rf tmp

release: pkg-linux pkg-darwin
	echo '$(VERSION) release built'

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf mysqlhealth_*.tar.gz

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v

# Cross compilation
build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
