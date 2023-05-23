package main

import "fmt"

func main() {
	var readers, book, magazine, overall int
	var k map[string]map[string]string = map[string]map[string]string{
		"Петя":   {"Книга": "Онегин"},
		"Олег":   {"Книга": "Крестный отец"},
		"Игорь":  {"Журнал": "Игромания"},
		"Андрей": {"Журнал": "Ф1"},
		"Роман":  {"Журнал": "Playboy"},
	}
	for key, value := range k {
		if key != "" {
			readers += 1
		}
		for key1, _ := range value {
			if key1 == "Книга" {
				book += 1
			}
			if key1 == "Журнал" {
				magazine += 1
			}
		}
		overall = book + magazine
	}
	fmt.Println("Количество читателей", readers, "\n КОличество книг", book, "\n КОличество журналов", magazine)
	fmt.Println("\nОбщее количество изданий", overall)
}
