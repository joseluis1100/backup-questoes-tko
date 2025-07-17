package main
import "fmt"
func main() {
    var j1, j2, j3, j4, soma int
    fmt.Scan(&j1, &j2 ,&j3 ,&j4)
    soma = j1+ j2 + j3 + j4
    if soma == 0 {
        fmt.Println("nenhum")
    } else if soma % 4 == 1 {
        fmt.Println("jog1")
    } else if soma % 4 == 2 {
        fmt.Println("jog2")
    } else if soma % 4 == 3 {
        fmt.Println("jog3")
    } else if soma % 4 == 0 {
        fmt.Println("jog4")
    }
}
