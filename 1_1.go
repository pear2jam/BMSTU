package main

import(
	"fmt"
	"strconv"
)

func main(){
	var ind, hunderd int64 = 0, 1 //ind - искомый индекс в блоках не ниже данного, hunderd - количество элементов в блоке
	fmt.Scanln(&ind)
	ind += 1
	for i := 1 ; ; i++ {
		if ind <= 9*int64(i)*hunderd {
			ind -= 1
			fmt.Println(string(strconv.FormatInt(hunderd + int64(ind/int64(i)), 10)[ind % int64(i)]))
			break
		}
		ind -= 9*int64(i)*hunderd
		hunderd *= 10
	}
}

