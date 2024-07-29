package main

import (
	"fmt"
	"log"
	"math/rand/v2"
)

func main() {
	var ans, guesses int
	var flag bool
	ans = rand.IntN(100)
	fmt.Println("Введите число от 1 до 100:")
	_, err := fmt.Scan(&guesses)
	if err != nil {
		log.Fatal(err)
	}
	for i := 9; i >= 0; i-- {
		switch {
		case ans == guesses:
			flag = true
			fmt.Println("Ура! Вы угадали загаданное число!")
		case ans < guesses:
			fmt.Println("Неправильно. Загаданное число меньше предлагаемого.")
		case ans > guesses:
			fmt.Println("Неправильно. Загаданное число больше предлагаемого.")
		}
		if flag == true {
			switch {
			case i == 9:
				fmt.Println("Вы отзадали за 1 попытку.")
			case i <= 8 && i >= 6:
				fmt.Println("Вы отгадали за", 10-i, "попытки.")
			default:
				fmt.Println("Вы отзадали за", 10-i, "попыток.")
			}
			i = 0
		} else {
			switch {
			case i <= 9 && i >= 5:
				fmt.Println("У вас осталось", i, "попыток. Попробуйте еще.")
				_, err := fmt.Scan(&guesses)
				if err != nil {
					log.Fatal(err)
				}
			case i <= 4 && i >= 2:
				fmt.Println("У вас осталось", i, "попытки. Попробуйте еще.")
				_, err := fmt.Scan(&guesses)
				if err != nil {
					log.Fatal(err)
				}
			case i == 1:
				fmt.Println("У вас осталась", i, "попытка. Попробуйте еще раз.")
				_, err := fmt.Scan(&guesses)
				if err != nil {
					log.Fatal(err)
				}
			case i == 0:
				fmt.Println("К сожалению, у Вас не осталось попыток, попробуйте свои силы еще раз в новой игре!")
			}
		}
	}
}
