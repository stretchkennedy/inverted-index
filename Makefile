BINS      := client server

PROTOS    := $(wildcard proto/*.proto)
PROTOGENS := $(PROTOS:proto/%.proto=gen/%.pb.go)

all: $(BINS)
$(BINS): $(PROTOGENS)
	go build -o $@ cmd/$@/*.go

gen/%.pb.go: proto/%.proto gen
	protoc -I proto --go_out=gen/ $<

gen:
	mkdir gen

.PHONY: clean
clean:
	rm -rf $(BINS) gen
