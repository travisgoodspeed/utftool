
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
	./utftool -diagram '"\xc3\xa9coutez!"'
	./utftool -runes aəiő
	./utftool -runes '"a\u0259i\u0151"'
fmt:
	go fmt
