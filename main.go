package main

import "fmt"

func main(){
	f()
	fmt.Println("b")
	fmt.Println("f")
}

func f(){
	defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
		if err:=recover();err!=nil{
			fmt.Println(err)
			// 这里的err其实就是panic传入的内容，55
			// recover here, stop the function but cut the panic
		}
	}()
	fmt.Println("a")

	//cause panic
	panic(55)

}