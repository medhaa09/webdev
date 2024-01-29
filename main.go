package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Cipher interface {
	encrypt() []int
}

type MyString string
type IntArray []int
type MyMap map[string]int

func (s MyString) Encrypt() []int {
	result := make([]int, len(s))
	for i, char := range s {
		result[i] = int(char) - 64
	}

	return result
}
func (a IntArray) Encrypt() []int {
	result := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		for a[i] != 1 {
			if a[i]%2 == 0 {
				a[i] = a[i] / 2
			} else {
				a[i] = 3*a[i] + 1
			}

		}
		result[i] = a[i]
	}
	return result
}
func (m MyMap) Encrypt() []int {
	result := make([]int, len(m))
	var values []int
	var keys []rune
	for k, v := range m {
		values = append(values, v)
		keys = append(keys, []rune(k)...)
	}
	for i, k := range keys {
		result[i] = int(k) + values[(i+2)%len(keys)]

	}
	return result

}
func main() {

	var Input interface{}
	fmt.Println("Enter input (enter array elements with comma and map key value pairs with comma as well and colon between each key and value")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		Input = scanner.Text()
		var intArr IntArray
		flag := 0
		inputBytes := []byte(Input.(string))
		if err := json.Unmarshal(inputBytes, &intArr); err == nil {
			result := intArr.Encrypt()
			fmt.Println(result)
			flag = 1
		}
		var mapp MyMap
		inputBytes2 := []byte(Input.(string))
		if err2 := json.Unmarshal(inputBytes2, &mapp); err2 == nil {
			result := mapp.Encrypt()
			fmt.Println(result)
			flag = 1
		}
		if flag == 0 {
			result := Input.(MyString).Encrypt()
			fmt.Println(result)
		}

	}
}
