package main
import "fmt"


//没有空间是不可以进行赋值的！
//如果一个数组木有可用的空间，那么就是非法的


func main()  {
	//1.声明slice是一个切片，并且默认值是1，2，3
	slice := []int {1,2,3}
	//%v表示打出详细信息
	fmt.Printf("len = %d ,slice = %v\n",len(slice),slice) 
	//len = 3 ,slice = [1 2 3]

	//2.声明slice是一个切片，但是没有去分配空间
	var slice1[]int
	fmt.Printf("len = %d ,slice1 = %v\n",len(slice1),slice1) 
	//len = 0 ,slice1 = []
	//对其开辟空间
	slice1 = make([]int, 3)
	fmt.Printf("len = %d ,slice1 = %v\n",len(slice1),slice1) 
	//len = 3 ,slice1 = [0 0 0]

	//3.声明slice并且给他分配空间
	var slice2[]int = make([]int, 3)
	fmt.Printf("len = %d ,slice2 = %v\n",len(slice2),slice2) 
	//len = 3 ,slice2 = [0 0 0]

	//4.声明slice并且给他分配空间 ==> 目前最常用
	slice3 := make([]int, 3)
	fmt.Printf("len = %d ,slice3 = %v\n",len(slice3),slice3) 
	//len = 3 ,slice3 = [0 0 0]

	if slice == nil {
		fmt.Printf("是空切片\n")
	}else{ //else的两侧的括号是必须在一行的
		fmt.Printf("不是空切片\n")
	}
}