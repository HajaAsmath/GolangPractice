package Hydradblayer

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongostruct struct {
	db *mgo.Session
}

func NewMongoDatabase(cstring string) (*mongostruct, error) {
	session, err := mgo.Dial(cstring)
	if err != nil {
		return nil, err
	}
	return &mongostruct{
		db: session,
	}, nil
}

func (mongo *mongostruct) AddMember(cm CrewMember) error {
	session := mongo.db.Copy()
	defer session.Close()
	personnel := session.DB("Hydra").C("Personnel")
	return personnel.Insert(cm)
}

func (mongo *mongostruct) GetMemberById(id int) (CrewMember, error) {
	session := mongo.db.Copy()
	defer session.Close()
	personnel := session.DB("Hydra").C("Personnel")
	cm := CrewMember{}
	err := personnel.Find(bson.M{"id": id}).One(&cm)
	if err != nil {
		return cm, err
	}
	return cm, nil
}

func (mongo *mongostruct) GetAllMembers() (crew, error) {
	session := mongo.db.Copy()
	defer session.Close()
	personnel := session.DB("Hydra").C("Personnel")
	cr := crew{}
	err := personnel.Find(bson.M{}).Iter().All(&cr)
	if err != nil {
		fmt.Println("Error fetching rows", err)
		return nil, err
	}
	return cr, nil
}
