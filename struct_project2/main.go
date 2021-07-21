package main

import "fmt"

type contactInfo struct {
	email string
	phone int
}

type person struct {
	name string
	age  int
	
	//embedded struct
	contact contactInfo
}

// type persons []person
// func main() {
// 	allPersons:=persons{}
// 	newPerson:=person{name:"Sudipta",age:12}
// 	allPersons = append(allPersons,newPerson )
// 	newPerson=person{name:"Shorobana",age:11}
// 	allPersons = append(allPersons,newPerson )
// 	fmt.Println(allPersons)
// }

func main() {
	var sudipta person

	sudipta.name="Sudipta"
	sudipta.age=23
	sudipta.contact.email="adak07151@gmail.com"
	sudipta.contact.phone=8910345282
	fmt.Println(sudipta)
	

	shrobana:=person{
		name:"Shrobana",
		age:22,
		contact:contactInfo{
			email:"shrobana1996@gmail.com",
			phone:8910345678,
		},
	}
	shrobana.print()

	kylie:= &person{name:"kyli",age:23}
	kylie.print()
	kylie.personDetails(22,"Kylie")
	kylie.print()
	a:=99
	b:=87
	fmt.Println()
	fmt.Println(a,b)
	a,b=swap(&a,&b)
	fmt.Println(a,b)

	//pointer not required
	//for slice,maps,channels

	myslice:=[]string{"Hey,","You","Doing","Well!"}
	fmt.Println(myslice)
	updateSlice(myslice)
	fmt.Println(myslice)

}


func (p person) print(){
	fmt.Printf("%v",p)
}

func (obj *person)personDetails(age int,name string){
	(*obj).age=age
	(*obj).name=name
}


func swap(x *int,y *int) (a int, b int){
	temp:=*x
	*x=*y
	*y=temp
	return *x,*y
}


func updateSlice(s []string){
	s[0]="Hello,"
}