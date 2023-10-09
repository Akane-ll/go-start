package main

import (
	"fmt"
	"reflect"
)

func main() {
	//声明一个Cat结构体
	type Cat struct {
		Name string
		Type int `json:"type" id:"100"`
	}
	//创建一个Cat实例
	ins := Cat{
		Name: "mimi",
		Type: 1,
	}
	fmt.Printf("ins type: %T\n", ins)// ins type: main.Cat

	//获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	fmt.Printf("typeOfCat type: %T\n", typeOfCat)// typeOfCat type: *reflect.rtype

	//遍历结构体成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		//获取每个结构体成员的字段类型
		fieldType := typeOfCat.Field(i)
		//输出成员名和tag
		fmt.Printf("name: %v, tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	//name: Name, tag: ''
	//name: Type, tag: 'json:"type" id:"100"'

	//通过字段名找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		//从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))// type 100
	}
}
