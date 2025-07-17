package main
import "fmt"
func main() {
    var x int
    fmt.Scan(&x)
    for c := range x + 1 {
        if x == 0 {
            fmt.Println("nao")
            break
        }else if c * c == x {
            fmt.Println("sim")
            break
        } else if c == x {
            fmt.Println("nao")
        }
    }
}