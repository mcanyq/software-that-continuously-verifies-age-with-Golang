package main
import "fmt"

func main(){

	for{
		fmt.Println("What is your name?: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println("How old are you?: ")
	var age int
	fmt.Scanln(&age)

	if age<18{
		difference:= 18-age

		fmt.Println("Sorry!",name,"We cannot let you in because you are under 18..",difference,"We wait after all these years")
	
		}else {
			fmt.Println("Welcome",name,"Have a good time!")
		}
	}
}