package main
import "fmt"
func main() {
    var n, m int
    fmt.Scan(&n, &m)
    for c := 1; ; c++ {
        if c % m == 0 && c % n == 0 {
            fmt.Println(c)
            break
        }
    }
}
