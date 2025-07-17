package main
import "fmt"
func main() {
    frase := make([]rune, 0)
    var x rune
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        frase = append(frase, x)
    }
    var letra rune
    var contador int
    fmt.Scanf("%c", &letra)
    for i := range frase {
        if frase[i] == letra {
            contador++
        }
    }
    fmt.Println(contador)
}
