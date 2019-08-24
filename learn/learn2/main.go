package main

import (
	"fmt"
	"log"
	"reflect"
)

type Gift struct {
	Sender    string
	Recipient string
	Number    uint
	Contents  string
}

func (g Gift) Info() {
	log.Printf("%s-to-%d", g.Sender, g.Number)
}

func (g *Gift) SetContents(content string) string {
	g.Contents = content
	return content
}

func nice(i interface{}) {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice {
		return
	}
	if e := v.Type().Elem(); e.Kind() != reflect.Struct {
		return
	}
	st := v.Type().Elem()
	if nameField, found := st.FieldByName("Sender"); found == false || nameField.Type.Kind() != reflect.String {
		return
	}
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		sender := e.FieldByName("Sender")
		num := e.FieldByName("Number")
		if sender.String() == "ming" {
			num.SetUint(100)
		}
	}

}

func main() {
	g := Gift{
		Sender:    "Hank",
		Recipient: "Minx",
		Number:    1,
		Contents:  "apple",
	}
	t := reflect.TypeOf(g)
	if kind := t.Kind(); kind != reflect.Struct {
		log.Fatalf("This program expects to work on a struct; we Got a %v instead.", kind)
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		log.Printf("Field %03d: %-10.10s %v", i, f.Name, f.Type.Kind())
	}
	for i := 0; i < t.NumMethod(); i++ {
		log.Printf("%s", t.Method(i).Name) //Info
	}
	//动态调用方法
	gv := reflect.ValueOf(&g)
	info := gv.MethodByName("Info")
	info.Call([]reflect.Value{})

	setM := gv.MethodByName("SetContents")
	res := setM.Call([]reflect.Value{reflect.ValueOf("andy")})
	log.Print(res)
	log.Print(g)

	gs := []Gift{
		{Sender: "ming"},
		{Sender: "xi"},
	}
	nice(gs)
	log.Printf("%v", gs) //[{ming  100 } {xi  0 }]

	x := 2
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(100)
	fmt.Println("x=", x)
}
