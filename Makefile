BINS      := client server
DIRS      := bin gen

PROTOS    := $(wildcard proto/*.proto)
PROTOGENS := $(PROTOS:proto/%.proto=gen/%.pb.go)

all: $(BINS)
$(BINS): $(PROTOGENS) bin
	go build -o bin/$@ src/cmd/$@/*.go

gen/%.pb.go: proto/%.proto gen
	protoc -I proto --go_out=gen/ $<

$(DIRS):
	mkdir $@

.PHONY: clean
clean:
	rm -rf $(DIRS)
