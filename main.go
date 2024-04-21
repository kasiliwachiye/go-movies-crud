package main

import "fmt"

type Movie struct {
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Director *Director `json:"director"`
}

type Director struct {
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
}

var movies []Movie

func main() {
  fmt.Println("Hello, World!")
}