package main
import "fmt"
func main() {
    var atual, anterior rune
    for {
        fmt.Scanf("%c", &atual)
        if atual == '\n' {
            break
        }
        if atual == ' ' && anterior == ' ' {
            continue
        }
        fmt.Printf("%c", atual)
        anterior = atual
    }
    fmt.Println()
}
