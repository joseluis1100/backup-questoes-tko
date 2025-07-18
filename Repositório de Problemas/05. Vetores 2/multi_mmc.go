package main
import "fmt"
func main() {
    var n int
    x := 2
    fmt.Scan(&n)
    vetor := make([]int, n)
    for i := range vetor {
        fmt.Scan(&vetor[i])
    }
    for {
        for i := range vetor {
            if x % vetor[i] != 0 {
                break
            } 
            if i == len(vetor)-1 {
                fmt.Println(x)
                return
            }
        }
        x++
    }
}
