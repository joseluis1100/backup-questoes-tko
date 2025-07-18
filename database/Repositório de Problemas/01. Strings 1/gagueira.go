package main
import "fmt"
func main() {
    frase := make([][]rune, 1)
    var x rune
    var linha int
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        if x == ' ' {
            frase = append(frase, []rune{})
            linha++
        } else {
            frase[linha] = append(frase[linha], x)
        }
    }
    for i := range frase {
        fmt.Printf("%v %v", string(frase[i]), string(frase[i]))
        if i == len(frase) - 1 {
            fmt.Println()
        } else {
            fmt.Print(" ")
        }
    }
}
