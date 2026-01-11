package main

import "fmt"

func main(){
	var n int 
	fmt.Scan(&n)
	for i:=1;i<=n;i++{
		fmt.Printf("%d",i)
		if i!=n{
			fmt.Printf(" + ")
		}else{
			fmt.Printf(" = %d",n*(n+1)/2)
		}
	}
}