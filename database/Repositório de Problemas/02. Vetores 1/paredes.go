package main
import "fmt"
func main() {
    var q, p, m int
    fmt.Scan(&q)
    n := make([]int, q)
    for c := range n {
        fmt.Scan(&n[c])
    }
    for c := range n {
        if n[c] > m {
            m = n[c]
            p++
        }
    }
    fmt.Println(p)
}
