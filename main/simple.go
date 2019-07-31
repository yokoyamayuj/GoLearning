package main

import(
	"fmt"
	"gosample/sample"
)

func main(){

  //変数定義
  var message string = "普通に定義したメッセージ"
  messe := "型を省略した定義方法"

  //インポート
  fmt.Println("インポートした定数："+ sample.Message)
  fmt.Println(messe)
  fmt.Println(message)

  //IF文
  if true {
	  fmt.Println("if文の条件式に()は不要")
  } else if 1 < 2 {
	  fmt.Println("else if文の条件式に()は不要")
  }

  //For文
  fmt.Println("ループにも()はいらない")
  for i:=0; i<3; i++{
	fmt.Println(i)
  }

  //関数
  sample.Sum(1,2)

  // 複数返却
  fmt.Println("retrun sample1")
  x ,y := sample.ReturnSample(1,2)
  fmt.Println(x,y)


  // 戻り値を定義し、returnでは省略
  fmt.Println("retrun sample2")
  x ,y = sample.ReturnSample2(1,2)
  fmt.Println(x,y)


  // TODO:スライス、配列

  // TODO:MAP

  //ポインタ
  fmt.Println("ポインタ渡し")
  var s string
  sample.CallbyRef(&s)
  fmt.Println(s)

  // TODO:panic,recover,defer

}

