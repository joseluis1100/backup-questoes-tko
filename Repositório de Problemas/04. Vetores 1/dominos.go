package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    dominos := make([]int, n)
    for c := range dominos {
        fmt.Scan(&dominos[c])
        if c > 0 && dominos[c] < dominos[c-1] {
            fmt.Println("precisa de ajuste")
            return
        }
    }
    fmt.Println("ok")
}
