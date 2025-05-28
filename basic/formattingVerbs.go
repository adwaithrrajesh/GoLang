package main 
import ("fmt");

func main (){
	var i int = 80;
	var txt string = "text";

	  fmt.Printf("%v\n", i)
	  fmt.Printf("%#v\n", i)
	  fmt.Printf("%v%%\n", i) //Prints the % sign 
	  fmt.Printf("%T\n", i) // Prints Datatype = int
	  
	  fmt.Printf("%v\n", txt) //Prints the value in the default format
	  fmt.Printf("%#v\n", txt) //Prints the value in Go-syntax format
	  fmt.Printf("%T\n", txt) // Prints Datatype  = string
}