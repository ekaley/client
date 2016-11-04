ORGANIZATION = josephgorse
PROJECT = ipam-client
BINARYNAME = ipam-client
GOOUT = ./bin
MAIN = cmd/main.go

default: deps build

deps:
	go get ./...

build:
		go build -o $(GOOUT)/$(BINARYNAME) $(MAIN)

clean:
	rm -f $(GOOUT)/$(BINARYNAME)
