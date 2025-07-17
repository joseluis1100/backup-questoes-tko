package main
import "fmt"
func main() {
    var a, x rune
    fmt.Scanf("%c\n", &a)
    r := 0
    for {
        fmt.Scanf("%c", &x)
        if x == '\n' {
            break
        }
        if a == x {
            r++
        }
    }
    fmt.Println(r)
}
