package main

import (
	"encoding/json"
	. "fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

type result struct {
	Path         []int
	Result       bool
	Attemptsleft int
	Time         string
}

func save(path []int, result_ans bool, attemptsleft int) {
	var result_struct result = result{path, result_ans, attemptsleft, time.Now().Format("2006-01-02 15:04:05")}
	var results []result
	file, err := os.OpenFile("results.json", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	decoder.Decode(&results)
	results = append(results, result_struct)
	file.Close()
	file, err = os.OpenFile("results.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(results)
	if err != nil {
		log.Fatal("Ошибка записи JSON:", err)
	}

}

func main() {
	var run, result bool = true, false
	var input, max, ans, attempts int = 0, 100, 0, 5
	var err error
	var discard string

	for run {
		Println("1 - Настройка")
		Println("2 - Старт")
		Println("0 - Выход")
		_, err = Scanln(&input)
		if err != nil {
			Println("Неверное значение")
			Scanln(&discard)
			continue
		}

		switch input {
		case 1:

			Println("Выбери максимальное значение")
			_, err = Scanln(&input)
			if err != nil {
				Println("Неверное значение")
				Scanln(&discard)
				continue
			}
			max = input
			Println("Выбери количество попыток")
			_, err = Scanln(&input)
			if err != nil {
				Println("Неверное значение")
				Scanln(&discard)
				continue
			}
			attempts = input

		case 2:
			var path []int
			ans = rand.Intn(max)
			Println("Число загадано")
			Println("Введи предположение")
			for i := 0; i < attempts; i++ {
				_, err = Scanln(&input)
				if err != nil {
					Println("Неверное значение")
					Scanln(&discard)
					continue
				}
				path = append(path, input)
				diff := math.Abs(float64(input - ans))
				if input > ans {
					if diff > 15 {
						Println("Холодно, загаданное число меньше")
					} else if diff > 5 {
						Println("Тепло, загаданное число меньше")
					} else {
						Println("Горячо, загаданное число меньше")
					}
				} else if ans > input {
					if diff > 15 {
						Println("Холодно, загаданное число больше")
					} else if diff > 5 {
						Println("Тепло, загаданное число больше")
					} else {
						Println("Горячо, загаданное число больше")
					}
				} else {
					color.Green("Поздравляю ты победил!")
					result = true
					save(path, result, attempts-i-1)
					break
				}
				if !result && (attempts-i) != 1 {
					color.Yellow("Осталось %d\n", (attempts - i - 1))
				}
			}
			if !result {
				color.Red("Попытки кончились игра окончена")
				save(path, result, 0)

			}

			result = false
			Println("Сыграть ещё?")
			Println("1 - Да")
			Println("2 - Нет")

			_, err = Scanln(&input)
			if err != nil {
				Println("Неверное значение")
				Scanln(&discard)
				continue
			}
			if input != 1 {
				run = false
			}
		case 0:
			run = false
		default:
			Println("Неверное значение")
		}
	}
}
