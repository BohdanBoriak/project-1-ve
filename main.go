package main

import (
	"fmt"
	"math/rand"
	"project-1-ve/domain"
	"strconv"
	"time"
)

const (
	points            int = 10
	pointsPerQuestion int = 5
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у найкращій грі! Йде завантаження...")
	time.Sleep(1 * time.Second)

	var users []domain.User

	for {
		menu()
		punct := ""
		fmt.Scan(&punct)

		switch punct {
		case "1":
			u := play()
			users = append(users, u)
		case "2":
			fmt.Println("Поки рейтингу не завезли... =(")
		case "3":
			return
		default:
			fmt.Println("Не коректний вибір")
		}
	}
}

func menu() {
	fmt.Println("1. Почати гру")
	fmt.Println("2. Переглянути рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
	fmt.Println("Підготуватись!")

	for i := 3; i > 0; i-- {
		fmt.Printf("До початку: %v\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("ВПЕРЕД!")
	myPoints := 0
	pointsPerGame := points
	now := time.Now()
	for pointsPerGame > 0 {
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
				pointsPerGame -= pointsPerQuestion
				myPoints += pointsPerQuestion
				fmt.Printf("Правильно! У тебе: %v очок!\n", myPoints)
				fmt.Printf("Залишилось зібрати: %v очок!\n", pointsPerGame)
			} else {
				fmt.Println("Не правильно! Спробуй ще!")
			}
		}
	}
	then := time.Now()
	spent := then.Sub(now)

	fmt.Printf("Ти топ! Впорався за %s\n", spent)

	fmt.Println("Введіть своє ім'я: ")
	playerName := ""
	fmt.Scan(&playerName)

	// var user domain.User
	// user.Id = id
	// user.Name = playerName
	// user.Time = spent

	user := domain.User{
		Id:   id,
		Name: playerName,
		Time: spent,
	}

	id++
	return user
}
