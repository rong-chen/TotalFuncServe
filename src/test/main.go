package main

type MyInterface interface {
	Print()
	Log()
}

func TestFunc(x MyInterface) {
	x.Print()
}

type MyStruct struct {
}

func (me MyStruct) Print() {
	//fmt.Println(123)
}
func (me MyStruct) Log() {
	//fmt.Println(123)
}
func main() {
	var me MyStruct
	TestFunc(me)
}
