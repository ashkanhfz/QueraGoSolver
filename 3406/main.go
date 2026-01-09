package main

import "fmt"

func reverse(x int) (y int){
	for x>0{
		y*=10
		y+=x%10
		x/=10
	}
	return
}


func main(){
	var x1,x2 int
	fmt.Scan(&x1,&x2)
	if x1 == x2{
		fmt.Printf("%d = %d",x1,x2)
	}else if reverse(x1)<reverse(x2){
		fmt.Printf("%d < %d",x1,x2)
	}else{
		fmt.Printf("%d < %d",x2,x1)
	}
}

