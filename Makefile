NAME=$(lastword $(subst /, ,$(abspath .)))
VERSION=$(shell git.exe describe --tags 2>nul || echo v0.0.0)
GOOPTC=-ldflags "-s -w -X main.version=$(VERSION)"
GOOPTG=-ldflags "-s -w -X main.version=$(VERSION) -H windowsgui"

ifeq ($(OS),Windows_NT)
    SHELL=CMD.EXE
    SET=SET
else
    SET=export
endif

all:
	go fmt $(foreach X,$(wildcard internal/*),&& cd $(X) && go fmt && cd ../..)
	cd cmd/msiver  && go fmt && $(SET) "CGO_ENABLED=0" && go build -o ../../msiver.exe $(GOOPTC)
	cd cmd/gmsiver && go fmt && $(SET) "CGO_ENABLED=0" && go build -o ../../gmsiver.exe $(GOOPTG)

_package:
	$(MAKE) all
	zip $(NAME)-$(VERSION)-$(GOOS)-$(GOARCH).zip msiver.exe gmsiver.exe

package:
	$(SET) "GOOS=windows" && $(SET) "GOARCH=386"   && $(MAKE) _package
	$(SET) "GOOS=windows" && $(SET) "GOARCH=amd64" && $(MAKE) _package

manifest:
	make-scoop-manifest *-windows-*.zip > $(NAME).json
