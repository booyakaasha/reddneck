package main

import "fmt"

// обычное объявление
func singleIn(in int) int {
	return in
}

// много параметров
func multIn(a, b int, c int) int {
	return a + b + c
}

// именованный результат
func namedReturn() (out int) {
	out = 2
	return
}

// несколько результатов
func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}

// несколько именованных результатов
func multipleNamedReturn(ok bool) (rez int, err error) {
	rez = 1
	if ok {
		err = fmt.Errorf("some error happend")
		// аналогично return rez, err
		return 3, fmt.Errorf("some error happend")
		return
	}
	rez = 2
	return
}

// не фиксированное количество параметров
func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	// fmt.Println(multipleNamedReturn(false))
	// return

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums, sum(nums...))
	return

	/*

	   func (p Person) String() string {
	   return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
	   }
	   (!!!) Объявление метода выглядит так же, как объявление функции, но с одним от-
	   личием: здесь дополнительно указывается приемник метода. Приемник метода
	   указывается после ключевого слова func перед именем метода. Как и при объ-
	   явлении любой другой переменной, сначала указывается имя приемника, а за-
	   тем — его тип. Согласно общепринятому соглашению имя приемника должно
	   представлять собой сокращение от имени типа: обычно используется только
	   первая буква имени типа. Использование в качестве имени приемника слова
	   this или self не соответствует идиоматическому подходу.

	   приемники указателей (когда используется указательный тип) или приемники
	   значений (когда используется значимый тип).
	   Определиться с тем, когда следует
	   использовать тот или иной вид приемника, вам помогут следующие правила.
	   Если метод вносит изменения в приемник, необходимо использовать при-
	   емник указателей.
	   Если метод должен учитывать вероятность того, что экземпляр будет равен
	   nil, необходимо использовать приемник указателей.
	   Если метод не вносит изменения в приемник, можно использовать приемник
	   значений.



	*/

}
