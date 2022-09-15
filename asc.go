//Mostra alguns caracteres e seus
//codigos AscII

package main

 import "fmt"

 func main() {
 for	s := 33; s <= 255; s++ {
		if s >= 127 && s <= 160 { s = 161
		}
		fmt.Println("Caractere ", s, " ->", string(rune(s)))
	}
}
