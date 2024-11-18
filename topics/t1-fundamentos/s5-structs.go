package t1fundamentos

type person struct {
	Name string
	Age  int
	sayHello func()
}

func TestStruct() {

	var p person
	p.Name = "John"
	p.Age = 25
	p.sayHello = func() {
		println("Hello, my name is", p.Name)
	}
	p.sayHello()
}