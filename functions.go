package Pass

import (
	"fmt"
	"math/rand"
	"os"
)

func HowLongThePassword() int {
	var num int
	fmt.Println("How long the password you want?")
	fmt.Scan(&num)
	return num
}

func CheckIfFileExist(nb int) string {
	name := "PasswordFile" + string(nb+'0')
	_, err := os.Stat(name)
	for err == nil {
		nb++
		name = "PasswordFile" + string(nb+'0')
		_, err = os.Stat(name)
	}
	if err != nil {
		os.Create(name)
	}
	return name
}

func CreatePassword(num int) []byte {
	var Password string
	var char rune
	for i := 0; i < num; i++ {
		char = rune(rand.Intn(93) + 33)
		Password += string(char)
	}
	Password += "\n"
	pass := []byte(Password)
	return pass
}

func HowManyPasswordDoYouWant() {
	var nbr int
	var choice int8
	fmt.Println("Hello how many Password do you want ?")
	fmt.Scan(&nbr)
	fmt.Println("In 1- one file or 2- a file for each password (1 or 2 )?")
	fmt.Scan(&choice)
	if choice == 1 {
		Choice1(nbr)
	} else if choice == 2 {
		Choice2(nbr)
	} else {
		fmt.Println("Invalid Input!")
	}
}

func Choice1(nbr int) {
	var pass string
	var file string
	nb := 0
	name := CheckIfFileExist(nb)
	fmt.Println("Loading....")
	for i := 0; i < nbr; i++ {
		num := HowLongThePassword()
		pass = string(CreatePassword(num))
		file += pass
	}
	os.WriteFile(name, []byte(file), 0644)
	fmt.Println("Done")
}

func Choice2(nbr int) {
	for i := 1; i <= nbr; i++ {
		num := HowLongThePassword()
		nb := 0
		name := CheckIfFileExist(nb)
		pass := CreatePassword(num)
		os.WriteFile(name, pass, 0644)
	}
}
