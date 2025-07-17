package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	unico := make([]int, 0)
	rep := make([]Pair, 0)
	for c := range vet {
		if !slices.Contains(unico, int(math.Abs(float64(vet[c])))) {
			unico = append(unico, int(math.Abs(float64(vet[c]))))
		}
	}
	slices.Sort(unico)
	for c := range unico {
		var r int
		for c2 := range vet {
			if int(math.Abs(float64(vet[c2]))) == unico[c] {
				r++
			}
		}
		rep = append(rep, Pair{unico[c], r})
	}
	return rep
}

func teams(vet []int) []Pair {
	time := make([]Pair, 0)
	r := 1
	for c := range vet {
		if c < len(vet)-1 && vet[c] == vet[c+1] {
			r++
		} else {
			time = append(time, Pair{vet[c], r})
			r = 1
		}
	}
	return time
}

func mnext(vet []int) []int {
	mnext := make([]int, len(vet))
	for c := range vet {
		if vet[c] > 0 && ((c > 0 && vet[c-1] < 0) || (c < len(vet)-1 && vet[c+1] < 0)) {
			mnext[c] = 1
		} else {
			mnext[c] = 0
		}
	}
	return mnext
}

func alone(vet []int) []int {
	alone := make([]int, len(vet))
	for c := range vet {
		if vet[c] < 0 {
			alone[c] = 0
		} else if vet[c] > 0 && ((c > 0 && vet[c-1] < 0) || (c < len(vet)-1 && vet[c+1] < 0)) {
			alone[c] = 0
		} else {
			alone[c] = 1
		}
	}
	return alone
}

func couple(vet []int) int {
	h := make([]int, 0)
	m := make([]int, 0)
	var casal int
	for c := range vet {
		if vet[c] > 0 {
			h = append(h, vet[c])
		} else {
			m = append(m, vet[c])
		}
	}
	for c := range h {
		for c2 := range m {
			if h[c] == int(math.Abs(float64(m[c2]))) {
				casal++
				m[c2] = 0
				break
			}
		}
	}
	return casal
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	for c := range len(vet)-len(seq)+1 {
		if slices.Equal(vet[c:c+len(seq)], seq) {
			return c
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	nvet := make([]int, 0)
	for c := range vet {
		for c2 := range posList {
			if c == posList[c2] {
				break
			} else if c2 == len(posList)-1 {
				nvet = append(nvet, vet[c])
			}
		}
	}
	return nvet
}

func clear(vet []int, value int) []int {
	nvet := make([]int, 0)
	for c := range vet {
		if vet[c] != value {
			nvet = append(nvet, vet[c])
		}
	}
	return nvet
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
