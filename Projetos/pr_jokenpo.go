package main

import (
	"fmt"
	"math/rand"
)
func sj(j int) string {
    var sj string
    if j == 1 {
        sj = "PEDRA"
    } else if j == 2 {
        sj = "PAPEL"
    } else if j == 3 {
        sj = "TESOURA"
    } else if j == 4 {
        sj = "SPOCK"
    } else {
        sj = "LAGARTO"
    }
    return sj
}
func v1() {
    var j1, j2, pontos1, pontos2 int
    var sj1, sj2 string
    for round := 1 ; round <= 5; round++ {
        fmt.Println("# JOKENPÔ V1 #")
        fmt.Printf("Você: %v | PC: %v\n", pontos1, pontos2)
        fmt.Printf("Round: %v / 5\n\n", round)
        fmt.Println("1 - Pedra")
        fmt.Println("2 - Papel")
        fmt.Println("3 - Tesoura")
        fmt.Print(">> ")
        fmt.Scan(&j1)
        sj1 = sj(j1)
        j2 = rand.Intn(3)+1
        sj2 = sj(j2)
        fmt.Printf("Você jogou %v e o PC %v\n", sj1, sj2)
        if j1 == j2{
            fmt.Println("Ninguém ganhou!")
        } else if j1 == 1 && j2 == 3 || j1 == 2 && j2 == 1 || j1 == 3 && j2 == 2 {
            fmt.Println("Você ganhou!")
            pontos1++
        } else {
            fmt.Println("O PC ganhou!")
            pontos2++
        }
        fmt.Println()
    }
    fmt.Println("PLACAR FINAL:")
    fmt.Printf("Você: %v | PC : %v\n\n", pontos1, pontos2)
}

func v2() {
    var j1, j2, pontos1, pontos2 int
    var sj1, sj2 string
    for round := 1 ; round <= 5; round++ {
        fmt.Println("# JOKENPÔ V2 #")
        fmt.Printf("Você: %v | PC: %v\n", pontos1, pontos2)
        fmt.Printf("Round: %v / 5\n\n", round)
        fmt.Println("1 - Pedra")
        fmt.Println("2 - Papel")
        fmt.Println("3 - Tesoura")
        fmt.Println("4 - Spock")
        fmt.Println("5 - Lagarto")
        fmt.Print(">> ")
        fmt.Scan(&j1)
        sj1 = sj(j1)
        j2 = rand.Intn(5)+1
        sj2 = sj(j2)
        fmt.Printf("Você jogou %v e o PC %v\n", sj1, sj2)
        if j1 == j2{
            fmt.Println("Ninguém ganhou!")
        } else if j1 == (j2+1)%5 || j1 == (j2+3)%5 {
            fmt.Println("Você ganhou!")
            pontos1++
        } else {
            fmt.Println("O PC ganhou!")
            pontos2++
        }
        fmt.Println()
    }
    fmt.Println("PLACAR FINAL:")
    fmt.Printf("Você: %v | PC : %v\n\n", pontos1, pontos2)
}

func main() {
    var v int
    fmt.Println("VOCÊ QUER JOGAR A VERSÂO 1 OU 2?")
    fmt.Print(">> ")
    fmt.Scan(&v)
    for r := 1; r != 0; {
        if v == 1 {
            v1()
        } else {
            v2()
        }
        fmt.Println("JOGAR NOVAMENTE?")
        fmt.Println("1 - Sim")
        fmt.Println("0 - Sair")
        fmt.Print(">> ")
        fmt.Scan(&r)
        fmt.Println()
    }
}
