package main

import "fmt"

type Human struct {
  name string
  age int
  phone string
}

type Employee struct{
  Human
  speciality string
  phone string
}

func main()  {
   Bob := Employee{Human{"heriipurnama",21,"444-xxxx"},"Designer","666-666"}

   fmt.Println("bobs work phone is : ",Bob.phone)
   fmt.Println("Bobs personal phone is : ",Bob.Human.phone)
}
