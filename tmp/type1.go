package tmp

import "fmt"

type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

func Type1() {
	var t interface{}
	t = Person{Name: "run.wu"}
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T 打印任何类型的 t
	case bool:
		fmt.Printf("boolean %t\n", t) // t 是 bool 类型
	case int:
		fmt.Printf("integer %d\n", t) // t 是 int 类型
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t 是 *bool 类型
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t 是 *int 类型
	}

}

func type2() {
	var t interface{}
	t = Person{Name: "run.wu"}
	fmt.Println(t)
}
