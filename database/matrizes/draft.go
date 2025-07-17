package main
import "fmt"
func main() {
    var linhas, colunas int
    fmt.Scan(&linhas, &colunas)
    mat1 := make([][]int, linhas)
    mat2 := make([][]int, linhas)
    mat3 := make([][]int, linhas)
    for i := range mat1 {
        mat1[i] = make([]int, colunas)
        mat2[i] = make([]int, colunas)
        mat3[i] = make([]int, colunas)
    }
    for i := range mat1 {
        for j := range mat1[i] {
            fmt.Scan(&mat1[i][j])
        }
    }
    for i := range mat2 {
        for j := range mat2[i] {
            fmt.Scan(&mat2[i][j])
            mat3[i][j] = mat1[i][j] + mat2[i][j]
        }
    }
    for i := range mat3 {
        fmt.Print("[ ")
        for j := range mat3[i] {
            fmt.Print(mat3[i][j])
            if j == len(mat3[i])-1 {
                fmt.Println(" ]")
            } else {
                fmt.Print(" ")
            }
        }
    }
}
