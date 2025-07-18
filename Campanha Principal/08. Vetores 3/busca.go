package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func in(vet []int, x int) bool {
	return slices.Contains(vet, x)
}

func indexOf(vet []int, x int) int {
	for c := range vet {
		if vet[c] == x {
			return c
		}
	}
	return -1
}

func findIf(vet []int) int {
	for c := range vet {
		if vet[c] > 0 {
			return c
		}
	}
	return -1
}

func minElement(vet []int) int {
	var menor int
	for c := range vet {
		menor = vet[c]
		break
	}
	for c := range vet {
		if vet[c] < menor {
			menor = vet[c]
		}
	}
	for c := range vet {
		if vet[c] == menor {
			return c
		}
	}
	return -1
}

func findMinIf(vet []int) int {
	var menor int
	for c := range vet {
		if vet[c] > 0 {
			menor = vet[c]
			break
		}
	}
	for c := range vet {
		if vet[c] > 0 && vet[c] < menor {
			menor = vet[c]
		}
	}
	for c := range vet {
		if vet[c] == menor {
			return c
		}
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("$" + line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]

		switch cmd {
		case "end":
			return
		case "in":
			if len(parts) >= 3 {
				vet := str2vet(parts[1])
				val, _ := strconv.Atoi(parts[2])
				if in(vet, val) {
					fmt.Println("true")
				} else {
					fmt.Println("false")
				}
			}
		case "index_of":
			if len(parts) >= 3 {
				vet := str2vet(parts[1])
				val, _ := strconv.Atoi(parts[2])
				fmt.Println(indexOf(vet, val))
			}
		case "find_if":
			if len(parts) >= 2 {
				vet := str2vet(parts[1])
				fmt.Println(findIf(vet))
			}
		case "min_element":
			if len(parts) >= 2 {
				vet := str2vet(parts[1])
				fmt.Println(minElement(vet))
			}
		case "find_min_if":
			if len(parts) >= 2 {
				vet := str2vet(parts[1])
				fmt.Println(findMinIf(vet))
			}
		default:
			fmt.Println("fail: unknown command")
		}
	}
}

func str2vet(s string) []int {
	if s == "[]" {
		return make([]int, 0)
	}
	sub := s[1 : len(s)-1]
	tokens := strings.Split(sub, ",")
	vet := make([]int, 0, len(tokens))
	for _, token := range tokens {
		num, _ := strconv.Atoi(strings.TrimSpace(token))
		vet = append(vet, num)
	}
	return vet
}
