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
		bal := 10000
		fmt.Println("Тогда начинаем. Твой баланс:", bal)
		for {
			fmt.Println("Лудка началась")
			time.Sleep(1 * time.Second)
			luck := rand.Intn(3)
			if luck == 0 {
				bal = 0
				fmt.Println("Ты проиграл всё! Твой баланс равен:", bal)
				time.Sleep(1 * time.Second)
				if bal == 0 {
					break
				}
			} else if luck == 1 {
				fmt.Println("Удача только собирается повернуться к тебе лицом!")
				fmt.Println("Твой баланс:", bal)
				time.Sleep(1 * time.Second)
				if bal == 0 {
					break
				}
			} else {
				bal = bal * 2
				fmt.Println("Удача на твоей стороне! Продолжай!")
				fmt.Println("Твой баланс равен:", bal)
				time.Sleep(1 * time.Second)
				if bal == 0 {
					break
				}
			}
		}
	} else if a == "Нет" {
		fmt.Println("Возвращайся когда наберешься смелости")
	} else {
		fmt.Println("Писать научись. Баран")
	}
}