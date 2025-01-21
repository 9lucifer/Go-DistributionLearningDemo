package main

import "fmt"

func main ()  {
	// var myarr1 [10] int
	myarr1:= [10] int{1,2,3,4}
	
	// for i := 0; i < len(myarr1); i++ {
	// 	fmt.Println(myarr1[i])
	// }

	for index,value:=range myarr1{
		fmt.Println("indexa:",index,",value = ",value)
	}
	

	//查看数组类型
	fmt.Printf("%T\n",myarr1)
}