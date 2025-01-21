package main
import "fmt"

//数组的长度是不固定的

//动态数组传参时候是引用传递的，方便在内部进行修改
func printArray(myarr []int)  {
	//不需要index的时候可以用小下划线代替
	for _ , value := range myarr{
		value = value - 1
		fmt.Println("value = ",value)
	} 
}

func main ()  {
	myarr2 := [] int {1,2,3,4}
	fmt.Printf("%T\n",myarr2)
	printArray(myarr2)
}