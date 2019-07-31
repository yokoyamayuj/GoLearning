package main

import(
	"fmt"
)

type Task struct{
  iD int
  name string
  isMan bool
}

func main(){
	simpleStruct()
	pointerSample()
	newSample()
}

// simpleな構造体
func simpleStruct(){
	var task Task = Task{
		iD : 1,
		name : "yuj",
		isMan : true,
	}

	fmt.Println(task.iD)
	fmt.Println(task.name)
	fmt.Println(task.isMan)
}
func pointerSample(){
	task := Task{1,"not changed",true}
	callByVal(task)
	fmt.Println(task.name)
	callByRef(&task)
	fmt.Println(task.name)
}
func callByVal(task Task){
	task.name="changed by val"
}
func callByRef(task *Task){
	task.name="changed by pointer"
}

func newSample(){
	var task *Task = new(Task) 
	fmt.Println("newの初期値: "+task.iD)
}