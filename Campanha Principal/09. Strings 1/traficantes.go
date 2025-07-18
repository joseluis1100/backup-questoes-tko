package main

import "fmt"

func main() {
	frase := make([]rune, 0)
	substituido := make([]rune, 0)
	substituto := make([]rune, 0)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		frase = append(frase, x)
	}
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		substituido = append(substituido, x)
	}
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		substituto = append(substituto, x)
	}
	nfrase := make([]string, 0)
	c := 0
	for {
		if string(frase[c:c+len(substituido)]) == string(substituido) {
			nfrase = append(nfrase, string(substituto))
			c += len(substituido)
		} else {

			nfrase = append(nfrase, string(frase[c]))
			c++
		}
		if c >= len(frase)-len(substituido)+1 {
			nfrase = append(nfrase, string(frase[len(frase)-1]))
			break
		}
	}
	for c := range nfrase {
		fmt.Printf("%s", nfrase[c])
	}
	fmt.Println()
}
