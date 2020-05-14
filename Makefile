DIR=.

all: $(DIR)/main

%: %.go
	go build -o $@ $<

run: $(DIR)/main
	./$<

clean:
	find . -type f -name main -delete
