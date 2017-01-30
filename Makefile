
TARGETS = utftool

all: $(TARGETS)

clean:
	rm -f $(TARGETS)

utftool: *.go
	go build

run: utftool
	./utftool -diagram écoutez!
fmt:
	go fmt
