package main
import "fmt"
func main() {
    var letra rune
    var pos int
    fmt.Scanf("%c\n%d", &letra, &pos)
    fmt.Printf("%c\n", (int(letra) + pos - 97 + 26)%26+97)
}
