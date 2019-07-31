package main

import(
	"fmt"
)

type Task struct{
   id int
   name string
}

func main(){
	var task Task = Task{
		id:1,
		name:"yuj",
	}
	fmt.Println(task.toString())
	task.callByVal()
	fmt.Println(task.name)
	task.callByRef()
	fmt.Println(task.name)

}

func (task Task) toString() string{
	return "recieverのname: " + task.name; 
}

func (task Task)callByVal(){
     task.name="changed by Val"
}

func (task *Task)callByRef(){
     task.name="changed by Ref"
}

