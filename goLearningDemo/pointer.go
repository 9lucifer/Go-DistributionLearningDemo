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