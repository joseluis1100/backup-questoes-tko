package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func count(vet []int, x int) int {
	var contador int
	for c := range vet {
		if vet[c] == x {
			contador++
		}
	}
	return contador
}

func sum(vet []int) int {
	var s int
	for c := range vet {
		s += int(math.Abs(float64(vet[c])))
	}
	return s
}

func average(vet []int) float64 {
	var s float64
	for c := range vet {
		s += math.Abs(float64(vet[c]))
	}
	return s / float64(len(vet))
}

func moreMen(vet []int) string {
	var qh, qm int
	for c := range vet {
		if vet[c] > 0 {
			qh++
		} else {
			qm++
		}
	}
	if qh > qm {
		return "men"
	} else if qm > qh {
		return "women"
	}
	return "draw"
}

func halfCompare(vet []int) string {
	var f, s int
	for c := range vet {
		if c < len(vet)/2 {
			f += int(math.Abs(float64(vet[c])))
		} else {
			s += int(math.Abs(float64(vet[c])))
		}
	}
	if f > s {
		return "first"
	} else if s > f {
		return "second"
	}
	return "draw"
}

func sexBattle(vet []int) string {
	var qh, qm int
	var sh, sm int
	for c := range vet {
		if vet[c] > 0 {
			qh++
			sh += int(math.Abs(float64(vet[c])))
		} else {
			qm++
			sm += int(math.Abs(float64(vet[c])))
		}
	}
	if sh / qh > sm / qm {
		return "men"
	} else if sm / qm > sh / qh {
		return "women"
	}
	return "draw"
}

// Funções auxiliares

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("$%s\n", line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "end":
			return
		case "count":
			fmt.Println(count(str2vet(args[1]), toInt(args[2])))
		case "sum":
			fmt.Println(sum(str2vet(args[1])))
		case "average":
			fmt.Printf("%.2f\n", average(str2vet(args[1])))
		case "more_men":
			fmt.Println(moreMen(str2vet(args[1])))
		case "half_compare":
			fmt.Println(halfCompare(str2vet(args[1])))
		case "sex_battle":
			fmt.Println(sexBattle(str2vet(args[1])))
		default:
			fmt.Println("fail: unknown command")
		}
	}
}

func str2vet(s string) []int {
	sub := s[1 : len(s)-1]
	parts := strings.Split(sub, ",")
	var vet []int
	for _, part := range parts {
		if part == "" {
			continue
		}
		val, _ := strconv.Atoi(part)
		vet = append(vet, val)
	}
	return vet
}

func toInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
