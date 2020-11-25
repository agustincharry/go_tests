package main

import (
	"fmt"
	"strconv"
)

type User struct {
	age  int
	name string
}

func main() {
	loopTest()
	arrayTest()
	sliceTest()
	makeAndAppendTest()
	structureTest()
}

func loopTest() {
	fmt.Println("---------------")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(strconv.Itoa(i) + " es par")
		} else {
			fmt.Println(strconv.Itoa(i) + " es impar")
		}
	}
}

func arrayTest() {
	fmt.Println("---------------")
	theArray := [3]int{1, 2, 3}
	fmt.Println(theArray)

	for i := 0; i < len(theArray); i++ {
		fmt.Println(theArray[i])
	}
}

func sliceTest() {
	fmt.Println("---------------")
	theArray := [5]int{9, 8, 7, 6, 5}
	slice := theArray[2:5]
	fmt.Println(slice)
}

func makeAndAppendTest() {
	fmt.Println("---------------")
	theArray := make([]int, 3, 5)
	fmt.Println(theArray)
	fmt.Println(len(theArray))
	fmt.Println(cap(theArray))
	theArray = append(theArray, 2)
	fmt.Println(theArray)
}

func structureTest() {
	fmt.Println("---------------")
	myUser := new(User)
	myUser.age = 25
	myUser.name = "Agustín"
	info := myUser.getUserInfo()
	fmt.Println(info)

	user2 := User{30, "Charry"}
	info2 := user2.getUserInfo()
	fmt.Println(info2)
}

func (this *User) getUserInfo() string {
	return "Age: " + strconv.Itoa(this.age) + ", Name: " + this.name
}
