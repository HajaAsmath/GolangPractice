package main

import (
	"Hydradblayer"
	"fmt"
	"log"
)

func main() {
	db, err := Hydradblayer.GetDBInstance("oracle", "oracle://SYSTEM:sachin123@localhost/orcl?TRACE FILE=trace.log")
	checkError(err)
	// cm := Hydradblayer.CrewMember{
	// 	Id:                2,
	// 	Name:              "Kylo Ren",
	// 	Securityclearance: 4,
	// 	Position:          "Chief",
	// }
	// err = db.AddMember(cm)
	cm, err := db.GetMemberById(2)
	checkError(err)
	fmt.Println(cm)

	db2, err := Hydradblayer.NewMongoDatabase("mongodb://127.0.0.1")

	cm1, err := db2.GetAllMembers()
	fmt.Println(cm1)

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
