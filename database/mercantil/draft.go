package main
import "fmt"
func main() {
    var r, p1, p2 int
    fmt.Scan(&r)
    v := make([]float32, r)
    c1 := make([]float32, r)
    c2 := make([]string, r)
    for c := range v {
        fmt.Scan(&v[c])
    }
    for c := range c1 {
        fmt.Scan(&c1[c])
    }
    for c := range c2 {
        fmt.Scan(&c2[c])
    }
    for c := range c2 {
        if (v[c] > c1[c] && c2[c] == "M") || (v[c] < c1[c] && c2[c] == "m") {
            p2++
        } else {
            p1++
        }
    }
    if p1 == p2 {
        fmt.Println("empate")
    } else if p1 > p2 {
        fmt.Println("primeiro")
    } else {
        fmt.Println("segundo")
    }
}
