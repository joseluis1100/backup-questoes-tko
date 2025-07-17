package main

import "fmt"

func main() {
	enigma := make([]rune, 0)
	key := make([]rune, 0)
	var x rune
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		enigma = append(enigma, x)
	}
	for {
		fmt.Scanf("%c", &x)
		if x == '\n' {
			break
		}
		x -= '0'
		key = append(key, x)
	}
	pos := 0
	mapa_enigma := make([]rune, len(enigma))
	for i := range enigma {
		mapa_enigma[i] = key[pos%len(key)]
		pos++
	}
	resolvido := make([]rune, len(enigma))
	for i := range enigma {
		resolvido[i] = enigma[i] ^ rune(mapa_enigma[i])
	}
	fmt.Println(string(resolvido))
}
