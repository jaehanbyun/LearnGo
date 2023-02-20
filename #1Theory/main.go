// package main

// import "fmt"

// func main() {
// 	const name1 string = "jae"
// 	var name2 string = "han"
// 	name2 = "han2"
// 	fmt.Println(name2)
// }

// 1.3 Functions
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // naked return, defer
// func lenAndUpper(name string) (lengtht int, uppercase string) {
// 	defer fmt.Println("I'm done")
// 	lengtht = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

// func main() {
// 	//totalLengtht, _ := lenAndUpper("jaehan")
// 	totalLengtht, upperName := lenAndUpper("jaehan")
// 	fmt.Println(totalLengtht, upperName)

// 	repeatMe("nico", "jaehan", "dal")
// }

// 1.5 for, range, ...args
// package main

// import (
// 	"fmt"
// )

// func supperAdd(numbers ...int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// func main() {
// 	result := supperAdd(1, 2, 3, 4, 5, 6)
// 	fmt.Println(result)
// }

// 1.6 if with a Twist
// package main

// import "fmt"

// // variable expressions
// func canIDrink(age int) bool {
// 	if koreanAge := age + 2; koreanAge < 18 {
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	fmt.Println(canIDrink(16))

// }

// 1.7 Switch

// package main

// import "fmt"

// func canIDrink(age int) bool {
// 	switch koreanAge := age + 2; koreanAge {
// 	case 18:
// 		return false
// 	case 20:
// 		return true
// 	case 50:
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	fmt.Println(canIDrink(16))
// }

// 1.8 Pointers
// package main

// import "fmt"

// func main() {
// 	a := 2
// 	b := &a
// 	*b = 20
// 	fmt.Println(a)
// }

// 1.9 Arrays and Slices
// package main

// import "fmt"

// func main() {
// 	names := [5]string{"jaehan", "jaejae", "jae"}
// 	names[3] = "alala"
// 	names[4] = "alala"
// 	//Slice
// 	names2 := []string{"jaehan", "jaejae", "jae"}
// 	// java, python의 append와는 다르게 go의 append는 값을 리턴함
// 	names2 = append(names2, "fly")
// 	fmt.Println(names2)
// }

// 1.10 Maps
// package main

// import "fmt"

// func main() {
// 	nico := map[string]string{"name": "nico", "age": "12"}
// 	for _, value := range nico {
// 		fmt.Println(value)
// 	}
// }

// 1.11 Structs
// package main

// import "fmt"

// type person struct {
// 	name    string
// 	age     int
// 	favFood []string
// }

// func main() {
// 	favFood := []string{"kimchi", "ramen"}
// 	nico := person{name: "nico", age: 18, favFood: favFood}
// 	fmt.Println(nico)
// }

// package main

// import (
// 	"fmt"

// 	"github.com/jaehan/learngo/accounts"
// )

// func main() {
// 	account := accounts.NewAccount("jaehan")
// 	account.Deposit(10)
// 	err := account.Withdraw(20)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(account)
// }
