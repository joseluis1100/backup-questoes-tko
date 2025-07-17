package main
import "fmt"
func main() {
    frase := make([]rune, 0)
    var x rune
    for {
        fmt.Scanf("%c",&x)
        if x == '\n' {
            break
        }
        if x == ' ' {
            continue
        }
        frase = append(frase, x)
    }
    frase_formatada := []rune(string(frase))
    if len(frase_formatada) % 2 != 0 {
        frase_formatada = append(frase_formatada, 'X')
    }
    for i := range frase_formatada {
        if i < len(frase_formatada)-1 && frase_formatada[i] == frase_formatada[i+1] {
            frase_formatada = append(frase_formatada[:i+1], append([]rune{'X'}, frase_formatada[i+1:]...)...)
        }
    }
    frase_dividida := make([]rune, 0)
    for i := range frase_formatada {
        frase_dividida = append(frase_dividida, frase_formatada[i])
        if i % 2 == 1 {
            frase_dividida = append(frase_dividida, ' ')
        }
    }
    fmt.Println(string(frase_dividida))
}
