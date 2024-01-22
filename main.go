package main

import (
	"fmt"
	"net/http"

	"github.com/cheynewallace/tabby"
	"github.com/gin-gonic/gin"
)

// type Team struct {
// 	name        string
// 	yearFounded int
// 	players     []Player
// }

type Player struct {
	Name         string
	Age          int
	JerseyNumber int // 21
	Team         string
}

// KOMPILATORN BESTÄMMER HEAP ELLER STACK
// new -> malloc = HEAP
// annars stack

// func test() *Player {
// 	// Ingen new
// 	foppa := Player{
// 		name:         "Stefan",
// 		age:          51,
// 		jerseyNumber: 21,
// 		team:         "Colorado",
// 	}
// 	zata := Player{
// 		name:         "Henrik",
// 		age:          48,
// 		jerseyNumber: 40,
// 		team:         "Timrå",
// 	}
// 	return &zata
// }

// func test(age int) {
// 	age = age + 1
// 	fmt.Println(age)
// }

// func testP(p *Player) {
// 	p.age = p.age + 1
// 	fmt.Println(p.age)
// }

// func testArray(age [4]int) {
// 	age[0] = age[0] + 1
// 	fmt.Println(age)
// }

// func testSlice(age []int) {
// 	age[0] = age[0] + 1
// 	age = append(age, 123)
// 	fmt.Println(age)
// }

func main() {

	r := gin.Default()
	r.GET("/contact", func(c *gin.Context) {
		c.String(http.StatusOK, "<h1>HEJ</h1>")
		//   c.JSON(http.StatusOK, gin.H{
		// 	"message": "pong",
		//   })
	})

	foppa := Player{
		Name:         "Stefan",
		Age:          51,
		JerseyNumber: 21,
		Team:         "Colorado",
	}
	zata := Player{
		Name:         "Test",
		Age:          2341,
		JerseyNumber: 2234,
		Team:         "Colorado",
	}
	zata2 := Player{
		Name:         "2",
		Age:          2341,
		JerseyNumber: 2234,
		Team:         "Colorado",
	}

	var players = [...]Player{zata, zata2, foppa}

	r.GET("/players", func(c *gin.Context) {
		c.JSON(http.StatusOK, players)
	})

	r.Run()

	//TABBY ÄR BARA ETT EXEMPEL - vi avänder den ALDRIG!!!
	// bra/enkel första  tredjepartsmodul PLUS New() medlemsfunktioner

	//fmt.Println(12,"Stefan")

	t := tabby.New()
	t.AddHeader("Name", "Age", "Jersey", "Team")
	for _, player := range players {
		t.AddLine(player.Name, player.Age, player.JerseyNumber, player.Team)
	}
	t.Print()

	// foppa := Player{
	// 	name:         "Stefan",
	// 	age:          51,
	// 	jerseyNumber: 21,
	// 	team:         "Colorado",
	// }
	// testP(&foppa) // COPY BY VALUE
	// fmt.Println(foppa.age)

	age := 51 // COPY BY VALUE
	// test(age)
	fmt.Println(age)

	// var ages = [...]int{50, 12, 134, 4} // COPY BY VALUE
	// testArray(ages)
	// fmt.Println(ages)

	// type Slice struct{
	// 	length int
	// 	Data *[] // Riktig array - backing array
	// }

	// var ages2 = []int{50, 12, 134, 4} // COPY BY VALUE
	// testSlice(ages2)
	// fmt.Println(ages2)

	// //USEL OOP?  - what about ensuring valid state etc?? Public private getter setters etc
	// foppa := Player{
	// 	name:         "Stefan",
	// 	age:          51,
	// 	jerseyNumber: 21,
	// 	team:         "Colorado",
	// }
	// fmt.Println(foppa.name)

	// // SINGLE SOURCE OF TRUTH
	// // ARRAY - vet att vi har FIXED antal saker
	// var cars = [...]string{"Volvo", "BMW", "Ford", "Mazda", "Renault", "Saab"}
	// langden := len(cars) // "OOP"   employee.GetSalary()
	// fmt.Println(langden)
	// // var months = []string{"Jan", "Feb"}
	// // // DUMT med siffran?
	// // var maxAntalvisitors = [5]string{}
	// // var antal = 0
	// // Men är det ingen skillnad mellan []string och […]string?

	// // antal++
	// // maxAntalvisitors[antal] = "Stefan"
	// // antal++
	// // maxAntalvisitors[antal] = ""

	// // NGTING SOM KAN MINSKA OCH ÖKA I ANTAL
	// var carsSlice = []string{"Volvo", "BMW", "Ford", "Mazda", "Renault", "Saab"}

	// // 1. Skapa ny car
	// //carsSlice.append("Lincoln")
	// // c - malloc PAXA 40b MINNE och får illbaka en minnesadress - realloc 120b

	// // Slices = dynamic size
	// carsSlice = append(carsSlice, "Lincoln")

	// //carSLIce "typ" som en pekare i C - slice datastruktur  slice.length slice.data = []

	// // char test[100];

	// // test[0] // index * sizeof(char) + addre(test)
	// // test[50]  // index * sizeof(char) + addre(test)

	// // 2. Ta bort en car

	// for _, value := range carsSlice {
	// 	fmt.Printf("Bil: %v\n", value)
	// }

	// fmt.Println(cars)

}
