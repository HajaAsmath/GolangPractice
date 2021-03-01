package Hydradblayer

import "errors"

const (
	mongo  = "mongo"
	oracle = "oracle"
)

type DBlayer interface {
	AddMember(cm CrewMember) error
	GetMemberById(id int) (CrewMember, error)
	GetAllMembers() (crew, error)
}

type CrewMember struct {
	Id                int    `json:"id",bson:"id"`
	Name              string `json:"name",bson:"name"`
	Securityclearance int    `json:"securityclearance",bson:"securityclearance"`
	Position          string `json:"position",bson:"position"`
}

type crew []CrewMember

func GetDBInstance(db string, cstring string) (DBlayer, error) {
	switch db {
	case mongo:
		db, err := NewMongoDatabase(cstring)
		if err != nil {
			return nil, err
		}
		return db, nil
	case oracle:
		db, err := NewOracleDatabase(cstring)
		if err != nil {
			return nil, err
		}
		return db, nil
	}
	return nil, errors.New("Invalid db")
}
