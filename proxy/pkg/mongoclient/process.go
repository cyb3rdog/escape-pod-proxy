// Cyb3rVector EscapePod Proxy
//  by cyb3rdog
//
// (c) 2021 Vaclav Macha
// All rights reserved.
//
// mongoclient - handles the communication with mongo db
//
package mongoclient

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//
func (client *MongoClient) InsertIntent(intent_json string) error {

	var newIntent interface{}
	err := bson.UnmarshalExtJSON([]byte(intent_json), true, &newIntent)
	if err != nil {
		return err
	}
	_, err = client.intents.InsertOne(context.TODO(), newIntent)
	if err != nil {
		return err
	}

	return nil
}

//
func (client *MongoClient) DeleteIntent(intent_id string) error {

	filter := bson.D{primitive.E{Key: "_id", Value: bson.D{primitive.E{Key: "$oid", Value: intent_id}}}}
	_, err := client.intents.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

//
func (client *MongoClient) SelectIntents(filter_json string) (map[string]string, error) {

	filter := bson.D{primitive.E{Key: "extended_options.external_parser", Value: true}}
	if filter_json != "" {
		err := bson.UnmarshalExtJSON([]byte(filter_json), true, &filter)
		if err != nil {
			return nil, err
		}
	}

	var results []interface{}
	cur, err := client.intents.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}

	json_data := make(map[string]string)
	var intent_id []byte
	var json_bytes []byte

	for _, record := range results {

		var json_record map[string]interface{}
		json_bytes, _ = bson.MarshalExtJSON(record, true, true)
		json.Unmarshal(json_bytes, &json_record)
		intent_id, _ = bson.MarshalExtJSON(json_record["_id"], true, true)
		json_data[string(intent_id)] = string(json_bytes)
	}

	return json_data, nil
}
