package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct { //container embeds a base. an embedding looks like a field without a name.
	base
	str string
}

func main() {

	co := container{
		base: base{ //embedded type serves as the field name
			num: 1,
		},
		str: "some name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe()) //methods of base also become methods of container

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}


