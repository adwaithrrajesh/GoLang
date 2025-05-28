// There are 2 ways to declare variable in go 


package main
import ("fmt")


// here is the example of how you can declare variable in go 👇🏼
func singleVariables()  {

	// Declaring variables
	var car string = "Lamborgini";
	newCar := "Buggatti"

	var car1 string 
	car1 = "toyota"

	var  car1Rate int 
	car1Rate = 200000

	fmt.Printf("Single Variables =>",car, car1, newCar,car1Rate)
}


// here is the example of how you can the variable 👇🏼
func multiVariables(){
	var a,b,c,d int=1,2,3,4
	fmt.Println("\n","Multi Variables1  =>",a,b,c,d)

	var e, f = 6,"10"
	g, h :=  1 ,"test2"

	fmt.Println("\n","Multi Varaibles2 => ", e,f,g,h)

	var(
	i int 
	j int =  1 
	k string = "k"
	)
	
	fmt.Println("\n","Variable block=>",i,j,k)

}

func constants(){
	const a = 123
	const b string = "test"
	fmt.Println("\n","Constant without Data type specification => ", a)
	fmt.Println("\n","Constant with Data type specification => ", b)

	const (
		c = "test"
		d int = 1
		e bool = true
	)

	fmt.Println("\n","Multi Constants =>",c,d,e)
}



func main(){
	singleVariables()
	multiVariables();
	constants();
}