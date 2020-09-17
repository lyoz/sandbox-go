DIR=.

all: $(DIR)/main

%: %.go
	go build -o $@ $<

run:
	go run main.go constant.go

clean:
	find . -type f -name main -delete
