package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)


type Deque struct {
    items []int
}

func (d *Deque) IsEmpty() bool {
	flag := true
	if len(d.items) > 0 {
		flag = false
	}
	return flag
}

func (d *Deque) PushFront(item int) {
    d.items = append([]int{item}, d.items...)
}

func (d *Deque) PushBack(item int) {
    d.items = append(d.items, item)
}

func (d *Deque) PopFront() (int, bool) {
    if len(d.items) == 0 {
        return 0, false
    }
    frontElement := d.items[0]
    d.items = d.items[1:]
    return frontElement, true
}

func (d *Deque) PopBack() (int, bool) {
    if len(d.items) == 0 {
        return 0, false
    }
    rearElement := d.items[len(d.items)-1]
    d.items = d.items[:len(d.items)-1]
    return rearElement, true
}


func getPath() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите путь к файлу/каталогу: ")
	path, _ := reader.ReadString('\n')
	return strings.TrimSpace(path)
}


func main() {
	// Открываем исходный файл для чтения
	
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка открытия исходного файла:", err)
		return
	}
	defer inputFile.Close()

	// Создаем новый файл для записи
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла для записи:", err)
		return
	}
	defer outputFile.Close()

	// Создаем bufio.Reader для чтения из исходного файла
	inputReader := bufio.NewReader(inputFile)

	// Создаем bufio.Writer для записи в новый файл
	outputWriter := bufio.NewWriter(outputFile)

	// Читаем и обрабатываем строки из исходного файла
	for {
		line, err := inputReader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		// Переносим цифры в конец строки
		var digits []rune
		var nonDigits []rune
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			} else {
				nonDigits = append(nonDigits, char)
			}
		}

		// Записываем строку в новый файл с цифрами в конце
		
		_, err = outputWriter.WriteString(string(nonDigits) + string(digits) + "\n")
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}
	}

	// Сбрасываем буфер и сохраняем изменения
	err = outputWriter.Flush()
	if err != nil {
		fmt.Println("Ошибка сохранения изменений в файле:", err)
		return
	}

	fmt.Println("Преобразование завершено. Результат сохранен в output.txt")
}



/* 
	Проблемы выполнения:

	1. inputReader.ReadString('\n')
	2. При разделении строки цифры переносятся на новую строку - DONE
	3. Дек нужно будет реализовать через дженерики, т.e. чтобы можно было менять тип данных 
*/



/* 
Лабораторная работа №1. Реализация стека/дека.
Разработать программу обработки данных, содержащихся в заранее подготовленном файле,
в соответствии с индивидуальным заданием. Программа должна включать модуль, содержащий
набор всех необходимых средств (типов, подпрограмм и т.д.) для решения поставленной задачи.
Порядок выполнения работы:
1) Получить у преподавателя индивидуальное задание.
2) Построить схему алгоритма решения задачи.
3) Использовать подпрограммы, реализующие полный набор операций
для этой структуры:
- допустимые операции для стека: инициализация, проверка на
пустоту, добавление нового элемента в начало, извлечение элемента из
начала;
- допустимые операции для дека: инициализация, проверка на пустоту,
добавление нового элемента в начало, добавление нового элемента в конец,
извлечение элемента из начала, извлечение элемента из конца.
4) Составить спецификации используемых подпрограмм.
5) Составить
программу,
включающую
соответствующей динамической структуры.
6) Проверить и продемонстрировать преподавателю работу программы
на полном наборе тестов. Обеспечить одновременный показ в окнах на
экране содержимого входного и выходного файлов.
7) Оформить отчет о лабораторной работе. 
9.Дан текстовый файл. Используя дек, переписать содержимое его строк в новый
текстовый файл, перенося при этом в конец каждой строки все входящие в нее цифры,
сохраняя исходный порядок следования среди цифр и среди остальных символов строки.
*/