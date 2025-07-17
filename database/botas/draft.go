package main
import "fmt"
func main() {
    var quantidade, tamanho, pares int
    var pe rune
    esquerdo := make(map[int]int)
    direito := make(map[int]int)
    fmt.Scan(&quantidade)
    for range quantidade {
        fmt.Scanf("%v %c ", &tamanho, &pe)
        if pe == 'E' {
            esquerdo[tamanho]++
        } else {
            direito[tamanho]++
        }
    }
    for valor := range esquerdo {
        if _, existe := direito[valor]; existe {
            pares += min(esquerdo[valor], direito[valor])
        }
    }
    fmt.Println(pares)
}
