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

func getMen(vet []int) []int {
	h := make([]int, 0)
	for c := range vet {
		if vet[c] > 0 {
			h = append(h, vet[c])
		}
	}
	return h
}

func getCalmWomen(vet []int) []int {
	m := make([]int, 0)
	for c := range vet {
		if vet[c] > -10 && vet[c] < 0 {
			m = append(m, vet[c])
		}
	}
	return m
}

func sortVet(vet []int) []int {
	slices.Sort(vet)
	return vet
}

func sortStress(vet []int) []int {
	c := 1
	for {
		if int(math.Abs(float64(vet[c]))) < int(math.Abs(float64(vet[c-1]))) {
			vet[c], vet[c-1] = vet[c-1], vet[c]
			c = 1
		} else {
			c++
		}
		if c == len(vet) {
			break
		}
	}
	return vet
}

func reverse(vet []int) []int {
	inverso := vet
	for c, f := 0, len(vet)-1; c < f; c, f = c+1, f-1 {
		inverso[c], inverso[f] = inverso[f], inverso[c]
	}
	return inverso
}

func reverseInplace(vet []int) {
	_ = vet
}
func unique(vet []int) []int {
	unico := make([]int, 0)
	for c := range vet {
		if !slices.Contains(unico, vet[c]) {
			unico = append(unico, vet[c])
		}
	}
	return unico
}

func repeated(vet []int) []int {
	unico := make([]int, 0)
	rep := make([]int, 0)
	for c := range vet {
		if !slices.Contains(unico, vet[c]) {
			unico = append(unico, vet[c])
		} else {
			rep = append(rep, vet[c])
		}
	}
	return rep
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			printVec(reverse(str2vet(args[1])))
		case "reverse_inplace":
			vet := str2vet(args[1])
			reverseInplace(vet)
			printVec(vet)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

