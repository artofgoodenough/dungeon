package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(values []int, i, j int) {
	temp := values[i]
	values[i] = values[j]
	values[j] = temp
	return
}

func sort(values []int) {
	if len(values) < 2 {
		return
	}
	if len(values) == 2 {
		if values[0] > values[1] {
			swap(values, 0, 1)
		}
		return
	}
	pivot := values[0]
	i := 1
	j := len(values) - 1
	for i < j {
		for i < j && values[i] <= pivot {
			i++
		}
		for j > i && values[j] >= pivot {
			j--
		}
		swap(values, i, j)
	}
	k := i
	if values[k] > pivot {
		k = i - 1
	}
	swap(values, 0, k)
	sort(values[:k])
	sort(values[(k + 1):])
}

func main() {
	values := make([]int, 0, 10)
	var fileName string
	fmt.Print("Filename: ")
	reader := bufio.NewReader(os.Stdin)
	fileName, _ = reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	f, _ := os.Open(fileName)
	reader = bufio.NewReader(f)
	for {
		valStr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		valStr = strings.TrimSpace(valStr)
		val, _ := strconv.Atoi(valStr)
		values = append(values, val)
	}
	f.Close()
	fmt.Println(values)
	sort(values)
	fmt.Println(values)
}
