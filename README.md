Howdy y'all,

This is a quick little tool for playing around with UTF-8 strings.
Intended for use in security research, this tool takes care to support
illegal runes and awkward regions.  The tool is primarily written in
Golang, but test cases are included in other languages for quickly
testing the legality of strings.

Since I don't expect many folks to use it, I've taken the liberty of a
short cut or two.  Sorry about that.

Examples:

```
./utftool -diagram "cat: кот"
Diagram:
63 61 74 3a 20 d0 d0 d1
               ba be 82
```

--Travis
