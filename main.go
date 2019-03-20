package main

import "fmt"

const helloWorld = "Gopher says: Hello %s %s\n"

func main() {

	name := getName()
	lastName := "Lopez"
	var number = sum(99, 1)
	a, b, c := getVariables()

	fmt.Printf(helloWorld, name, lastName)
	println(a, b, c, number)
}

func getName() string {
	var name string
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getVariables() (int, int, int) {
	return 1, 2, 3
}

func sum(a int, b int) int {
	return a + b
}
