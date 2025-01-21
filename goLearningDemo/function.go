package main
import "fmt"

// 单个返回值
func fool(a string , b int)int{
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	c:=100
	return c
}
// 多个返回值
func fool2(a string,b int)(int,int)  {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	return 666,777
}
// 多个有名字的返回值
func fool3(a string,b int)(ret1 int,ret2 int)  {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	ret1 = 111
	ret2 = 222
	// ret1 和 ret2 也是函数的形参，默认的初始化值是0
	//作用域是整个fool3函数
	return ret1,ret2
}
//也可以简写为
// func fool3(a string,b int)(ret1 ,ret2 int)  {
// 	fmt.Println("a = ",a)
// 	fmt.Println("b = ",b)
// 	ret1 = 111
// 	ret2 = 222
// 	return ret1,ret2
// }

func main()  {
	c := fool("abc",12)
	ret1,ret2 := fool2("aaa",12)
	fmt.Println("c = ",c)
	ret3,ret4 := fool3("aaa",12)
	fmt.Println("ret1 = ",ret1)
	fmt.Println("ret2 = ",ret2)
	fmt.Println("ret3 = ",ret3)
	fmt.Println("ret4 = ",ret4)
}