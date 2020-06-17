package main

import (
  "fmt"
  "bufio"
  "os"
)

type House struct {
  city string
  address string
  numberOfRooms int
  price int
}

func main() {

  first_house := 1
  last_house := 3

  // while loop
  for first_house < last_house {
    fmt.Println("List all your houses for sale")

    cityReader := bufio.NewReader(os.Stdin)
    fmt.Println("What is the city name?")
    city, _:= cityReader.ReadString('\n')

    addressReader := bufio.NewReader(os.Stdin)
    fmt.Println("What is the address?")
    address, _:= addressReader.ReadString('\n')

    roomsReader := bufio.NewReader(os.Stdin)
    fmt.Println("How many rooms does it have?")
    numberOfRooms, _:= roomsReader.ReadString('\n')

    priceReader := bufio.NewReader(os.Stdin)
    fmt.Println("How many price does it have?")
    price, _:= priceReader.ReadString('\n')

    fmt.Println("city: ", city, "address: ", address, "rooms: ", numberOfRooms, "price: ", price)

  }

}
