package Hydradblayer

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora"
)

type oraclestruct struct {
	db *sql.DB
}

func NewOracleDatabase(cstring string) (*oraclestruct, error) {
	conn, err := sql.Open("oracle", cstring)
	if err != nil {
		fmt.Println("Error connecting to database :", err)
		return nil, err
	}
	return &oraclestruct{
		db: conn,
	}, nil
}

func (oracle *oraclestruct) AddMember(cm CrewMember) error {
	res, err := oracle.db.Exec("Insert into Personnel(id,name,securityclearance,position) values(:1,:2,:3,:4)",
		cm.Id, cm.Name, cm.Securityclearance, cm.Position)
	if err != nil {
		return err
	}
	fmt.Println("Inserted Successfullu", res)
	return nil
}

func (oracle *oraclestruct) GetMemberById(id int) (CrewMember, error) {
	row := oracle.db.QueryRow("Select * from personnel where id=:1", id)
	cm := CrewMember{}
	if err := row.Scan(&cm.Id, &cm.Name, &cm.Securityclearance, &cm.Position); err != nil {
		fmt.Println("No rows available")
		return cm, err
	}
	return cm, nil
}

func (oracle *oraclestruct) GetAllMembers() (crew, error) {
	row, err := oracle.db.Query("select * from personnel")
	defer row.Close()
	if err != nil {
		fmt.Println("Error fetching rows")
		return nil, err
	}
	cr := crew{}
	for row.Next() {
		c := CrewMember{}
		row.Scan(&c.Id, &c.Name, &c.Securityclearance, &c.Position)
		cr = append(cr, c)
	}
	return cr, nil
}
