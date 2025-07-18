package main
import "fmt"
func main() {
    var n, min, max, i int
    fmt.Scan(&n, &min, &max)
    numeros := make([]int, n)
    for c := range numeros {
        fmt.Scan(&numeros[c])
        if numeros[c] >= min && numeros[c] <= max {
            i++
        }
    }
    fmt.Println(i)
}
