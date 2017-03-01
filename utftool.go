// This is a quick little tool for playing with valid and invalid UTF8
// strings, and for studying incompatibilities between string
// libraries.  See README.md.

package main

import (
	//"errors"
	"flag"
	"fmt"
	"strings"
	//"time"
	"unicode"
	//"unicode/utf8"
	"strconv"
)

var todiagram = flag.String("diagram", "", "Diagram a UTF8 string.")
var torunes = flag.String("runes", "", "Runes of the string.")

//! Prints the string runes.
func runes(foo []byte) string {
	for _, c := range string(foo) {
		fmt.Printf("Rune %q (%04x):\n", c, c)
		if unicode.IsControl(c) {
			fmt.Println("\tis control rune")
		}
		if unicode.IsDigit(c) {
			fmt.Println("\tis digit rune")
		}
		if unicode.IsGraphic(c) {
			fmt.Println("\tis graphic rune")
		}
		if unicode.IsLetter(c) {
			fmt.Println("\tis letter rune")
		}
		if unicode.IsLower(c) {
			fmt.Println("\tis lower case rune")
		}
		if unicode.IsMark(c) {
			fmt.Println("\tis mark rune")
		}
		if unicode.IsNumber(c) {
			fmt.Println("\tis number rune")
		}
		if unicode.IsPrint(c) {
			fmt.Println("\tis printable rune")
		}
		if !unicode.IsPrint(c) {
			fmt.Println("\tis not printable rune")
		}
		if unicode.IsPunct(c) {
			fmt.Println("\tis punct rune")
		}
		if unicode.IsSpace(c) {
			fmt.Println("\tis space rune")
		}
		if unicode.IsSymbol(c) {
			fmt.Println("\tis symbol rune")
		}
		if unicode.IsTitle(c) {
			fmt.Println("\tis title case rune")
		}
		if unicode.IsUpper(c) {
			fmt.Println("\tis upper case rune")
		}
		if c < 0x80 {
			fmt.Println("\tis an ASCII rune")
		}
	}
	return ""
}

//! Diagrams a UTF8 string.
func diagram(foo []byte) string {
	var slice [6]string

	i := 0   //input byte index.
	row := 0 //Row of the output string.
	for i < len(foo) {
		if (foo[i]&0x80 == 0) || (foo[i]&0xC0 == 0xC0) {
			//Reset the row on first-byte.
			row = 0
		} else {
			//Increment the row.
			row++
		}

		//Draw the byte.
		if row > 0 {
			for len(slice[row]) < len(slice[0])-2 {
				slice[row] = slice[row] + " "
			}
		}
		slice[row] = slice[row] + fmt.Sprintf("%02x ", foo[i])
		i++
	}
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n",
		slice[0], slice[1], slice[2], slice[3], slice[4], slice[5])
}

//! Unquotes a potentially damaged string, but only if it's
func unquote(foo string) string {
	t, err := strconv.Unquote(foo)

	if err != nil {
		//fmt.Printf("Unquote(%#v): %v\n", foo, err);

		// When there is a syntax error, that's probably
		// because the string isn't quoted.  We'll use the
		// literal parameter instead.
		return foo
	}

	//Return the parsed string if we can parse it.
	return t
}

//! Main method.
func main() {
	//Parses the command-line runes.
	flag.Parse()

	//Handle the runes here.
	if strings.Compare(*todiagram, "") != 0 {
		fmt.Println("Diagram:\n", diagram([]byte(unquote(*todiagram))))
	} else if strings.Compare(*torunes, "") != 0 {
		fmt.Println("Runes:")
		runes([]byte(unquote(*torunes)))
	} else {
		fmt.Println("Try --help.")
	}
}
