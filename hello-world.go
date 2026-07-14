package main
import ("fmt")

func main() {
	//fmt.Printf("Hello, world! \n")
	var age int = 53
	fmt.Println(age)

	var pointer *int 
	//fmt.Println(&pointer)
	pointer = &age
	
	fmt.Println(&pointer)
	fmt.Println(*pointer)
}