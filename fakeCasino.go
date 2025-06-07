package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	var a string
	fmt.Println("Готов проиграть всё, что у тебя есть? Да/Нет")
	fmt.Scan(&a)
	if a == "Да" {
		fmt.Println("Тогда начинаем")
		/**/
		for {
			fmt.Println("Лудка началась")
			time.Sleep(1 * time.Second)
			luck := rand.Intn(3)
			if luck == 0 {

			}
		}
	} else if a == "Нет" {
		fmt.Println("Возвращайся когда наберешься смелости")
	} else {
		fmt.Println("Писать научись. Баран")
	}
}