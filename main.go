package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	path := os.Args[1]
	tab := getFileContent(path)

	fmt.Printf("Average %x", math.Round(Average(tab)))
	fmt.Printf("Median %x", math.Round(Median(tab)))
	fmt.Println(int64(math.Round(Variance(tab))))
	fmt.Println(math.Round(math.Sqrt(Variance(tab))))
}

func Average(tab []int) (res float64) {
	for _, i := range tab {
		res += float64(i)
	}
	return res / float64(len(tab))
}

func Median(tab []int) (res float64) {
	SortIntegerTable(tab)
	if len(tab)%2 == 0 {
		val := float64(tab[len(tab)/2]+tab[(len(tab)/2)-1]) / 2
		return val
	}
	return float64(tab[len(tab)/2])
}

func Variance(tab []int) (res float64) {
	avg := Average(tab)
	for _, i := range tab {
		res += (float64(i) - float64(avg)) * (float64(i) - float64(avg))
	}
	return res / float64(len(tab))
}

func getFileContent(path string) []int {
	var intArray []int

	file, err := os.Open(path)
	CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Convertir la ligne en entier
		number, err := strconv.Atoi(line)
		CheckError(err)

		// Ajouter l'entier au tableau
		intArray = append(intArray, number)
	}
	return intArray
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			if table[j] > table[i] {
				Swap(&table[i], &table[j])
			}
		}
	}
}

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}
