// Create Slice 

package main
import ("fmt")

func createSlice(){
	myslice1 := []int{}
	fmt.Println(len(myslice1))// returns the length of the slice (the number of elements in the slice)
	fmt.Println(cap(myslice1))// returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	fmt.Println(myslice1)

	myslice2 := [2]int{2,4};
	fmt.Println(len(myslice2));
	fmt.Println(cap(myslice2));
	fmt.Println(myslice2)
}

func createSliceFromAnArray(){
	arr1 := [6] int{1,2,3,4,5,6};
	myslice := arr1[1:3];

	fmt.Printf("Create Slice From an array 👇🏼 \n")
	fmt.Printf("myslice = %v\n", myslice)
  	fmt.Printf("length = %d\n", len(myslice))
  	fmt.Printf("capacity = %d\n", cap(myslice))
}

func createSliceWithMake(){

	fmt.Printf("Create Slice with make 👇🏼 \n")
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
}


func modifySlice(){
	fmt.Printf("Modified Slice 👇🏼")
	prices := []int{10,20,30}
	fmt.Println(prices[0])
	fmt.Println(prices[2])
}

func changeElement(){
	fmt.Println("Change element 👇🏼");
	prices := []int{1,2,3};
	prices[2] = 10
	fmt.Println(prices[2]);
}

func main(){
	createSlice();
	createSliceFromAnArray();
	createSliceWithMake();
	changeElement();
}