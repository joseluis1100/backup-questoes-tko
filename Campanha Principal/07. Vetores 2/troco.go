package main
import "fmt"
func main() {
    var valor float32 //cria a vaariavel do valor inicial
    fmt.Scan(&valor) //recebe o valor float
    resto := int(valor) //inicialmente o resto será o valor original
    var notas_int = []int{100, 50, 20, 10, 5, 2, 1} //vetor com todos os valores de notas e 1 real
    var notas_float = []int{50, 25, 10, 5} //vetor com todos os valores de moedas
    for i := 0; i < len(notas_int); i++ { //for que precorre as notas inteiras
        if resto / notas_int[i] != 0 { //filtro para não imprimir valores nulos
            fmt.Printf("%d de %.2f\n", resto / notas_int[i], float32(notas_int[i]))
        }
        resto = resto % notas_int[i] //atualização do resto
    }
    resto = int(valor * 100) % 100 //atualização para se adequar as moedas
    for i := 0; i < len(notas_float); i++ { //for que precorre as moedas
        if resto / notas_float[i] != 0 { //filtro para não imprimir valores nulos
            fmt.Printf("%d de %.2f\n", resto / notas_float[i], float32(notas_float[i]) / 100)
        }
        resto = resto % notas_float[i] //atualização do resto
    }
    if resto != 0 { //filtro para não imprimir valores nulos
        fmt.Printf("Falta %.2f\n", float32(resto) / 100)
    }
}
