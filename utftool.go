// This is a quick little tool for playing with UTF8 strings, and for
// studying incompatibilities between string libraries.  See
// README.md.

package main

import (
	//"errors"
	"flag"
	"fmt"
	"strings"
	//"time"
)

var todiagram = flag.String("diagram", "", "Diagram a UTF8 string.");

//! Diagrams a UTF8 string.
func diagram(foo []byte) string{
	var slice [6]string;

	i:=0;      //input byte index.
	row := 0;  //Row of the output string.
	for i<len(foo){
		if (foo[i]&0x80==0) || (foo[i]&0xC0==0xC0) {
			//Reset the row on first-byte.
			row=0;
		}else{
			//Increment the row.
			row++;
		}

		//Draw the byte.
		if row>0 {
			for len(slice[row])<len(slice[0])-2{
				slice[row]=slice[row]+" ";
			}
		}
		slice[row]=slice[row]+fmt.Sprintf("%02x ",foo[i]);
		i++;
	}
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n",
		slice[0],slice[1],slice[2],slice[3],slice[4],slice[5]);
}

func main() {
	//Parses the command-line flags.
	flag.Parse()

	//Handle the flags here.
	if(strings.Compare(*todiagram,"")!=0){
		fmt.Println("Diagram:\n",diagram([]byte(*todiagram)));
	}else{
		fmt.Println("Try --help.");
	}
}
