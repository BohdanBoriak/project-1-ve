package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	points            int = 50
	pointsPerQuestion int = 5
)

func main() {
	fmt.Println("Вітаємо у найкращій грі! Йде завантаження...")
	time.Sleep(5 * time.Second)
	fmt.Println("Підготуватись!")

	for i := 5; i > 0; i-- {
		fmt.Printf("До початку: %v\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("ВПЕРЕД!")
	myPoints := 0
	now := time.Now()
	for points > 0 {
		x, y := rand.Intn(100), rand.Intn(100)
		res := x + y

		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Printf("Error: %s", err)
		} else {
			if ansInt == res {
				points -= pointsPerQuestion
				myPoints += pointsPerQuestion
				fmt.Printf("Правильно! У тебе: %v очок!\n", myPoints)
				fmt.Printf("Залишилось зібрати: %v очок!\n", points)
			} else {
				fmt.Println("Не правильно! Спробуй ще!")
			}
		}
	}
	then := time.Now()
	spent := then.Sub(now)

	fmt.Printf("Ти топ! Впорався за %s", spent)
	time.Sleep(10 * time.Second)
}
