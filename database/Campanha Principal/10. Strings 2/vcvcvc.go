package main
import "fmt"
func main() {
    frase := make([]rune, 0)
    for {
        var x rune
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        if x == ' '{
            frase = append(frase, ' ')
        } else if x == 65 || x == 69 || x == 73 || x == 79 || x == 85 || x == 97 || x == 101 || x == 105 || x == 111 || x == 117 {
            frase = append(frase, 'v')
        } else {
            frase = append(frase, 'c')
        }
    }
    for c := range frase {
        fmt.Printf("%c", frase[c])
    }
    fmt.Println()
}
