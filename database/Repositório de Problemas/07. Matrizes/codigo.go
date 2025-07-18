package main
import "fmt"
func main() {
    var quantidade, vezes int
    fmt.Scan(&quantidade)
    vetor := make([]int, quantidade)
    for i := range vetor {
        fmt.Scan(&vetor[i])
    }
    for i := range vetor {
        if i < len(vetor) - 2 && vetor[i] == 1 && vetor[i+1] == 0 && vetor[i+2] == 0 {
            vezes++
        }
    }
    fmt.Println(vezes)
}
