package main

import (
  "fmt"
  "math"
  "strconv"
)


type Person struct {
  name string
  age int
}

func (person Person) hello() {
  fmt.Println(person.name+ "my age is:"+ strconv.Itoa(person.age))
}

func (person *Person) addAge() {
  print("hengine..")
  person.age ++
}

func main() {
  fmt.Println("my app")
  fmt.Print(math.Floor(2.7))
  // arryaz()
  mymaps()
  // myrange()
  person1 := Person{name: "Layo", age: 45}
  person1.addAge()
  person1.hello()
}

func max(num1, num2 int) int {
  if (num1>num1) {
    return num1
  } else {
    return num2
  }
}


func myrange() {
  ids := []int{10,389,292,93}
  for i, id := range(ids) {
    fmt.Printf("%d - ID %d\n", i, id)
  }

  // printing the su o fhte array
  var sum = 0
  for _, id := range(ids) {
    sum += id
  }
  fmt.Println("Sum", sum)
}
func mymaps() {
  // email := make(map[string]string)

  // email["ben"] = "Bennkarson"
  // email["Terry"] = "Terry@gmail.com"

  email := map[string]string{"ben":"ben@gmail.om","sayo":"sayo@gmail.com"}

  for i, e := range(email) {
    fmt.Printf("%s : %s\n", i, e)
  }
  print(email["ben"])
}

func arryaz() {
  // a slice is an array without a fixed type
  // var fruitArray [2]string 
  // fruitArray[0] = "Apple"
  // fruitArray[1] = "Mango"

  fruitArray := [2]string{"apple","carrot"}

  fruitSlice := []string{"Yam","gexxui","noyyer"}
  fmt.Print(fruitArray, len(fruitSlice))
}