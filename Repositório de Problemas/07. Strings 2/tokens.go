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
        if x == '#' || x == ';' {
            x = '\n'
        }
        frase = append(frase, x)
    }
    fmt.Print(string(frase))
}
