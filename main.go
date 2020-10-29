package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Número de strings: ")
	fmt.Scan(&n)
	strs := make([]string, 0, n)
	scanner.Scan()
	for i := 0; i < n; i++ {
		fmt.Print("String #", i+1, ": ")
		scanner.Scan()
		s := scanner.Text()
		strs = append(strs, s)
	}
	sort.Strings(strs)

	asc := strings.Join(strs, "\n")

	sort.Sort(sort.Reverse(sort.StringSlice(strs)))

	desc := strings.Join(strs, "\n")

	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString(asc)

	file, err = os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString(desc)

}
