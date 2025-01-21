## 理论知识记录

pkg的包的最底层的init函数是最先执行的

### 关于导包的三种方式
1. 匿名
2. 起别名
3. 全部导入
<hr>
- 如果包导入之后不使用是会报错的
> 补救:在不直接用的包前面加上_表示是匿名导入
```go
import (
	_ "goLearningDemo/init/lib1"
	"goLearningDemo/init/lib2"
)
```

- 可以给包起别名
```go
import (
	_ "goLearningDemo/init/lib1"
	myLib2 "goLearningDemo/init/lib2"
)
```

- 如果是用.导包的话，视作是将包内所有的文件全都导入到目前的文件下面，那么当前文件内可以不用连着报名调用函数，而是直接调用
<span style="color: red;">尽量不用轻易使用</span>
```go
import (
	_ "goLearningDemo/init/lib1"
	. "goLearningDemo/init/lib2"
)
```


### 指针
```go
package main
import "fmt"

func swap(a * int,b *int){
	var tmp int
	tmp = *a
	*a = *b 
	*b = tmp
}

func main(){
	a:=100
	b:=200
	swap(&a,&b)

	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

}
```

### defer关键字
- 谁最先提出谁最后执行
- 如果defer和return都存在的情况，那么也是return先执行
```
package main
import "fmt"
func main ()  {
	defer fmt.Println("close1.....")
	defer fmt.Println("close2.....")
	fmt.Println("hello1")
	fmt.Println("hello2")
}
```
###### 结果
> hello1<br>
> hello2<br>
> close2.....<br>
> close1.....

### 切片 => 数组    (不初始化默认是0)

#### 静态数组
- 创建空数组，且大小固定
```go
	var myarr1 [10] int
```
遍历方式：
```go
    //第一种
	for i := 0; i < len(myarr1); i++ {
		fmt.Println(myarr1[i])
	}
    //第二种    
    for index,value:=range myarr1{
		fmt.Println("indexa:",index,",value = ",value)
	}
```

- 创建大小固定的数组，初始化了前几个
```go
	myarr1:= [10] int{1,2,3,4}
```


- 查看数组的类型
```go
	fmt.Printf("%T\n",myarr1)
    //res:[10]int
```

#### 动态数组
```go
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
```
#### 数组相关的api操作