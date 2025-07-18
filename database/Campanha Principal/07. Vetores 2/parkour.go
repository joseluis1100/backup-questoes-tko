package main
import "fmt"
func main() {
    var n, p int
    fmt.Scan(&n)
    b := make([]int, n)
    for c := range b {
        fmt.Scan(&b[c])
        if c > 0 && (b[c] > b[c-1]+1 || b[c] < b[c-1]-1) {
            p++
        }
    }
    fmt.Println(p)
}
