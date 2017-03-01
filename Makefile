
TARGETS = utftool

all: $(TARGETS)

clean:
	rm -f $(TARGETS)

utftool: *.go
	go build

install: utftool
	cp utftool /usr/local/bin/
run: utftool
	./utftool -diagram écoutez!
	./utftool -runes aəiő
fmt:
	go fmt
