package sample

import(
	"fmt"
)

// 大文字だと定数になる？？
var Message string = "hello world"

// 関数名の先頭大文字でないとダメ
func Sum (i,j int ){
	fmt.Println("関数")
	fmt.Println(i+j)
}

func ReturnSample(i,j int)(int,int){
	return i+1,i+j
}

func ReturnSample2(i,j int)(x int, y int){
	x=i+1 // すでに変数定義済み
	y=j+1
	return // x,yがreturnされる
}

// ポインタ
func CallbyRef(s *string){
	*s="ポインタ渡しで直接書き換え"
}