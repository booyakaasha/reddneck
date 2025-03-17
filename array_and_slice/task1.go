/*
package main

import (
	"fmt"
	"math"
)

// Average принимает массив целых чисел и возвращает среднее арифметическое, округленное до ближайшего целого.
func Average(array []int) int {
	if len(array) == 0 {
		return 0 // Возвращаем 0, если массив пустой
	}

	sum := 0
	for _, value := range array {
		sum += value
	}

	average := float64(sum) / float64(len(array)) // Вычисляем среднее
	return int(math.Round(average))               // Округляем до ближайшего целого
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Average(numbers)) // Вывод: 3
}
*/
/*
1. На вход подаются два неупорядоченных слайса любой длины. Надо написать функцию, которая возвращает их пересечение


package main

import "fmt"

func main2() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(peresechenie(a, b))
}

func peresechenie(a, b []int) []int {
	count := make(map[int]int)
	result := []int{}

	// Подсчитываем количество элементов в первом срезе
	for _, elem := range a {
		count[elem]++
	}

	// Проверяем элементы второго среза
	for _, elem := range b {
		if c, ok := count[elem]; ok && c > 0 {
			count[elem]--                 // Уменьшаем счетчик
			result = append(result, elem) // Добавляем элемент в результат
		}
	}

	return result
}
*/

/* генератор случайных чисел

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randNumsGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(n)
		}
		close(out)
	}()
	return out
}

func main() {
	for num := range randNumsGenerator(10) {
		fmt.Println(num)
	}
}
*/

/*
4. Сделать конвейер чисел

Даны два канала. В первый пишутся числа. Нужно, чтобы числа читались из первого по мере поступления, что-то с ними происходило (допустим, возводились в квадрат) и результат записывался во второй канал.

	В одной пишем в первый канал.
	Во второй читаем из первого канала и пишем во второй.

package main

import (

	"fmt"

)

	func main() {
		naturals := make(chan int)
		squares := make(chan int)

		go func() {
			for x := 0; x <= 10; x++ {
				naturals <- x
			}

			close(naturals)
		}()

		go func() {
			for x := range naturals {
				squares <- x * x
			}

			close(squares)
		}()

		for x := range squares {
			fmt.Println(x)
		}
	}
*/
package main
