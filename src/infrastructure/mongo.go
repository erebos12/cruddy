package infrastructure

import (
	"fmt"

	"configuration"

	. "utils"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var initialSession = initMongo()

func initMongo() *mgo.Session {
	method := "initMongo"
	session, err := mgo.Dial(configuration.GetMongoUrl())
	if err != nil {
		message := fmt.Sprintf("Error connecting to mongo. Error '%v'", err)
		LogIt(FATAL, method, message)
		panic(err)
	}
	message := fmt.Sprintf("Jippi  (°͜ °)...Mongo successfully connected. (%+v)", session)
	LogIt(INFO, method, message)
	return session
}

func getGollection(session *mgo.Session) mgo.Collection {
	method := "getGollection"
	db := *session.DB(configuration.GetMongoDB())
	collection := db.C(configuration.GetMongoCollection())
	LogIt(INFO, method, fmt.Sprintf("Successfully retrieved db name : '%s' and collection: '%s'", db.Name, collection.Name))
	return *collection
}

// Insert - insert a user to mongodb
func Insert(object interface{}) {
	method := "Insert"
	sess := initialSession.Copy()
	collection := getGollection(sess)
	defer sess.Close()
	err := collection.Insert(&object)
	if err != nil {
		message := fmt.Sprintf("Error while inserting to mongo: '%+v'\n", err)
		LogIt(ERROR, method, message)
		panic(err)
	}
	LogIt(INFO, method, fmt.Sprintf("Inserted object '%+v' successfully\n", object))
}

// FetchByGenericAttr - selects a document from mongo by attribute and value
func FetchByGenericAttr(attribute string, value string) ([]interface{}, error) {
	method := "FetchByGenericAttr"
	sess := initialSession.Copy()
	collection := getGollection(sess)
	defer sess.Close()
	var results []interface{}
	err := collection.Find(bson.M{attribute: value}).All(&results)
	if err != nil {
		LogIt(ERROR, method, fmt.Sprintf("Entity with attribute '%s' and value '%s' not found\n", attribute, value))
	}
	return results, err
}

// DeleteByGenericAttr - deletes a document from mongo by attribute and value
func DeleteByGenericAttr(attribute string, value string) error {
	method := "DeleteByGenericAttr"
	sess := initialSession.Copy()
	collection := getGollection(sess)
	defer sess.Close()
	err := collection.Remove(bson.M{attribute: value})
	if err != nil {
		LogIt(ERROR, method, fmt.Sprintf("( ͡° ʖ̯ ͡°) Cannot delete entity with attribute '%s' and value '%s' not found\n", attribute, value))
		return err
	}
	return nil
}

func DeleteAll() (string, error) {
	method := "DeleteAll"
	sess := initialSession.Copy()
	collection := getGollection(sess)
	defer sess.Close()
	info, err := collection.RemoveAll(nil)
	if err != nil {
		message := fmt.Sprintf("Could not delete all entries: '%+v'\n", err)
		LogIt(ERROR, method, message)
		return "", err
	}
	message := fmt.Sprintf("%d elements removed", info.Removed)
	LogIt(INFO, method, message)
	return message, nil
}

func GetAll() []interface{} {
	method := "GetAll"
	sess := initialSession.Copy()
	collection := getGollection(sess)
	defer sess.Close()
	var results []interface{}
	err := collection.Find(nil).All(&results)
	if err != nil {
		message := fmt.Sprintf("Could not get all entries: '%+v'", err)
		LogIt(ERROR, method, message)
	}
	LogIt(INFO, method, fmt.Sprintf("Results: %s", results))
	return results
}
