package main

import (
	"fmt"
	"math/rand"
	"time"
)

type orc struct {
	nome             string
	forca            int
	vida             int
	vivo             bool
	revidar          int
	revidar_original int
	raiva            int
}

func main() {
	vivos := 8
	a_vivos := 4
	b_vivos := 4
	vogais := []rune{'a', 'e', 'i', 'o', 'u'}
	consoantes := []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z'}
	time_a := make([]orc, 4)
	time_b := make([]orc, 4)
	fmt.Println("Time A:")
	time.Sleep(1 * time.Second)
	for i := range time_a {
		time_a[i].nome = string(consoantes[rand.Intn(len(consoantes))]) + string(vogais[rand.Intn(len(vogais))]) + string(consoantes[rand.Intn(len(consoantes))]) + string(vogais[rand.Intn(len(vogais))])
		time_a[i].forca = rand.Intn(100) + 1
		time_a[i].vida = rand.Intn(100) + 1
		time_a[i].vivo = true
		time_a[i].revidar = rand.Intn(3) + 1
		time_a[i].revidar_original = time_a[i].revidar
		time_a[i].raiva = rand.Intn(100) + 1
		fmt.Println("Orc", i+1)
		fmt.Println("Nome do Orc:", time_a[i].nome)
		fmt.Println("Força:", time_a[i].forca)
		fmt.Println("Vida:", time_a[i].vida)
		fmt.Println("Vivo:", time_a[i].vivo)
		fmt.Println("Revidar:", time_a[i].revidar)
		fmt.Println("Raiva:", time_a[i].raiva)
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Time B:")
	time.Sleep(1 * time.Second)
	for i := range time_b {
		time_b[i].nome = string(consoantes[rand.Intn(len(consoantes))]) + string(vogais[rand.Intn(len(vogais))]) + string(consoantes[rand.Intn(len(consoantes))]) + string(vogais[rand.Intn(len(vogais))])
		time_b[i].forca = rand.Intn(100) + 1
		time_b[i].vida = rand.Intn(100) + 1
		time_b[i].vivo = true
		time_b[i].revidar = rand.Intn(2)
		time_b[i].revidar_original = time_b[i].revidar
		time_b[i].raiva = rand.Intn(100) + 1
		fmt.Println("Orc", i+1)
		fmt.Println("Nome do Orc:", time_b[i].nome)
		fmt.Println("Força:", time_b[i].forca)
		fmt.Println("Vida:", time_b[i].vida)
		fmt.Println("Vivo:", time_b[i].vivo)
		fmt.Println("Revidar:", time_b[i].revidar)
		fmt.Println("Raiva:", time_b[i].raiva)
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
	time.Sleep(2 * time.Second)
	for rodada := 1; a_vivos >= 0 && b_vivos >= 0; rodada++ {
		for i := range time_a {
			time_a[i].revidar = time_a[i].revidar_original
			time_b[i].revidar = time_b[i].revidar_original
		}
        fmt.Println("-----------------------------")
		fmt.Println("Rodada", rodada)
        fmt.Println("-----------------------------")
		if a_vivos == 0 {
			fmt.Println("Time B venceu!")
			return
		}
		if b_vivos == 0 {
			fmt.Println("Time A venceu!")
			return
		}
		for i := range 4 {
			if b_vivos != 0 {
				if !time_a[i].vivo {
					continue
				}
				atacado := rand.Intn(4)
				for {
					if time_b[atacado].vivo {
						break
					} else {
						atacado = rand.Intn(4)
					}
				}
				fmt.Println("Orc", time_a[i].nome, "[A] ataca Orc", time_b[atacado].nome, "[B]")
				time_b[atacado].vida -= time_a[i].forca
				if time_b[atacado].vida <= 0 {
					time_b[atacado].vivo = false
					vivos--
					b_vivos--
					fmt.Println("Orc", time_b[atacado].nome, "[B] morreu")
					time.Sleep(1 * time.Second)
				} else {
					fmt.Println("Orc", time_b[atacado].nome, "[B] agora tem", time_b[atacado].vida, "de vida")
					time.Sleep(1 * time.Second)
				}
				if time_b[atacado].revidar > 0 {
					fmt.Println("Orc", time_b[atacado].nome, "[B] revida e ataca Orc", time_a[i].nome, "[A]")
					time_a[i].vida -= time_b[atacado].forca / 2
					if time_a[i].vida <= 0 {
						time_a[i].vivo = false
						vivos--
						a_vivos--
						fmt.Println("Orc", time_a[i].nome, "[A] morreu")
						time.Sleep(1 * time.Second)
					} else {
						fmt.Println("Orc", time_a[i].nome, "[A] agora tem", time_a[i].vida, "de vida")
						time.Sleep(1 * time.Second)
					}
				} else {
					fmt.Println("Orc", time_b[atacado].nome, "[B] não pode revidar")
					time_b[atacado].forca += time_b[atacado].raiva
					fmt.Println("Orc", time_b[atacado].nome, "[B] agora tem", time_b[atacado].forca, "de força")
					time.Sleep(1 * time.Second)
				}
			}
			if a_vivos != 0 {
				if !time_b[i].vivo {
					continue
				}
				atacado := rand.Intn(4)
				for {
					if time_a[atacado].vivo {
						break
					} else {
						atacado = rand.Intn(4)
					}
				}
				fmt.Println("Orc", time_b[i].nome, "[B] ataca Orc", time_a[atacado].nome, "[A]")
				time_a[atacado].vida -= time_b[i].forca
				if time_a[atacado].vida <= 0 {
					time_a[atacado].vivo = false
					vivos--
					a_vivos--
					fmt.Println("Orc", time_a[atacado].nome, "[A] morreu")
					time.Sleep(1 * time.Second)
				} else {
					fmt.Println("Orc", time_a[atacado].nome, "[A] agora tem", time_a[atacado].vida, "de vida")
					time.Sleep(1 * time.Second)
				}
				if time_a[atacado].revidar > 0 {
					fmt.Println("Orc", time_a[atacado].nome, "[A] revida e ataca Orc", time_b[i].nome, "[B]")
					time_b[i].vida -= time_a[atacado].forca / 2
					if time_b[i].vida <= 0 {
						time_b[i].vivo = false
						vivos--
						b_vivos--
						fmt.Println("Orc", time_b[i].nome, "[B] morreu")
						time.Sleep(1 * time.Second)
					} else {
						fmt.Println("Orc", time_b[i].nome, "[B] agora tem", time_b[i].vida, "de vida")
						time.Sleep(1 * time.Second)
					}
				} else {
					fmt.Println("Orc", time_a[atacado].nome, "[A] não pode revidar")
					time_a[atacado].forca += time_a[atacado].raiva
					fmt.Println("Orc", time_a[atacado].nome, "[A] agora tem", time_a[atacado].forca, "de força")
					time.Sleep(1 * time.Second)
				}
			}
		}
		fmt.Println("Orcs restantes:")
		fmt.Println("Time A:")
		if a_vivos == 0 {
			fmt.Println("todos os Orcs do Time A morreram")
		} else {
			for i := range time_a {
				if time_a[i].vivo {
					fmt.Println("Orc", time_a[i].nome, "está vivo com", time_a[i].vida, "de vida")
				}
			}
		}
		time.Sleep(2 * time.Second)
		fmt.Println("Time B:")
		if b_vivos == 0 {
			fmt.Println("todos os Orcs do Time B morreram")
		} else {
			for i := range time_b {
				if time_b[i].vivo {
					fmt.Println("Orc", time_b[i].nome, "está vivo com", time_b[i].vida, "de vida")
				}
			}
		}
		time.Sleep(2 * time.Second)
	}
}
