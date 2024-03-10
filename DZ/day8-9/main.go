package main

import "fmt"

type chel struct{ // создаём человека даём имя, возрост, рост, вес, деньги
	Name string
	Age uint
	Height uint
	Weight uint
	Money bool

}

func Create(m map[string]chel, key string, nChel chel){
	m[key] = nChel
}

func Read(m map[string]chel, key string) chel{
	return  m[key]
}

func Update(m map[string]chel, key string, updateChel chel){
	m[key] = updateChel
}

func Delete(m map[string]chel, key string) {
	delete(m, key)
}


func main () {
	chels := make(map[string]chel)
	nChel := chel{
		Name: "Jek",
		Age: 12,
		Height: 178,
		Weight: 80,
		Money: true,
	}
	Create(chels, "Jon", nChel)
	fmt.Println(Read(chels,"Jon"))
	updateChel := chel{
		Name: "Jons",
		Age: 25,
		Height: 180,
		Weight: 80,
		Money: false,
	}
	
	Update(chels, "Jon", updateChel)
	fmt.Println(Read(chels, "Jon"))

	Delete(chels, "Jon")
	fmt.Println(Read(chels, "Jon",))
}