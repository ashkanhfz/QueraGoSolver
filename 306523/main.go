package main

import "fmt"

func main(){
	var sum,s,c int64
	var n int
	fmt.Scan(&n)
	for ;n>0;n--{
		fmt.Scan(&s,&c)
		sum+=s*c
		
	}
	fmt.Println(sum)
}