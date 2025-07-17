package main
import "fmt"
func main() {
    var x int
    var x2 rune
    v := make([]int, 0)
    for {
        fmt.Scanf("%d%c", &x, &x2)
        v = append(v, x)
        if x2 == '\n' {
            break
        }
    }
    fmt.Print("[ ")
    for c := len(v)-1; c > 0; c-- {
        fmt.Printf("%d ", v[c])
    }
    fmt.Println(v[0], "]")
}
