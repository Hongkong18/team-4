EXECS=$(shell ls cmd)
MAIN=$(patsubst %, ./cmd/%/main.go, $(EXECS))

default: all

dep:
	go get -d -v ./...

%: ./cmd/%/main.go
	go build -v -o $@ $<

all: dep fmt vet $(EXECS)

tests:
	go test ./... -cover -v

tests_report:
	go test ./... -coverprofile=c.out
	go tool cover -html=c.out

clean:
	rm -f $(EXECS)
	rm -f c.out

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golint -set_exit_status ./...

cfg: cfggen
	./cfggen

.PHONY: $(MAIN) clean


