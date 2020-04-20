MAIN=main

all: $(MAIN)

%: %.go
	go build $<

run: $(MAIN)
	./$<
