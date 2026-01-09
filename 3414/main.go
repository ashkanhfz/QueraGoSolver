package main

import "fmt"


func main(){
	var x1,y1,x2,y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)
	if x1 == x2{
		fmt.Println("Vertical")
	}else if y1==y2{
		fmt.Println("Horizontal")
	}else{
		fmt.Println("Try again")
	}
}

