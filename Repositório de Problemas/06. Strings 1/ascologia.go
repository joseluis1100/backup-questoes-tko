package main
import "fmt"
func main() {
    var x rune
    var soma rune
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        soma += x
    }
    fmt.Println(soma % 50)
}
