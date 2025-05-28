package main
import ("fmt")

func forLoop() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      continue
    }
   fmt.Println(i)
  }
}

func whileLoop(){
	i:=0
	for  i<10{
		fmt.Println(i)
		i++;
	}
}

func main(){
	forLoop();
	whileLoop();
}