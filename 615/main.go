package main

import "fmt"


func main(){
	var x string
	fmt.Scan(&x)
	fmt.Printf("saal:%v\nmaah:%v",x[:2],x[2:])
}

