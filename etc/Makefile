# makes #
# https://gist.github.com/azer/7c83d0b59de8328355ad #
GOFILES=$(wildcard *.go)
GOTITLE=$(shell basename "$(PWD)")
GOREPORTS=./doc
GOTESTS=./src/...
GOBUILD=./build

# command #
.PHONY : run test cover build clean

run :
	@echo "MAKES » run"
	go run -v $(GOFILES)

test :
	@echo "MAKES » test"
	go test -cover -v $(GOTESTS)

cover :
	@echo "MAKES » cover"
	go test -coverprofile $(GOREPORTS)/cover.out -cover -v $(GOTESTS)
	go tool cover -html $(GOREPORTS)/cover.out -o $(GOREPORTS)/cover.html
	http-server $(GOREPORTS)

build : clean
	@echo "MAKES » build"
	gox -output "$(GOBUILD)/$(GOTITLE).{{.OS}}.{{.Arch}}" -verbose

clean :
	@echo "MAKES » clean"
	go clean -v
	@rm -rf $(GOBUILD)/*
