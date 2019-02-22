package convvls

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/loivis/convolvulus-update-triggers/c9r"
	"google.golang.org/genproto/googleapis/pubsub/v1"
)

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	CreateTime time.Time     `json:"createTime"`
	Fields     FireStoreData `json:"fields"`
	Name       string        `json:"name"`
	UpdateTime time.Time     `json:"updateTime"`
}

type FireStoreData struct {
	Author FireStoreStringValue
	Site   FireStoreStringValue
	Title  FireStoreStringValue
}

type FireStoreStringValue struct {
	StringValue string `json:"stringValue"`
}

func FirestoreTrigger(ctx context.Context, e FirestoreEvent) error {
	fav := parseEvent(&e)

	b, _ := json.Marshal([]*c9r.Favorite{fav})
	req := &pubsub.PublishRequest{
		Topic:    fmt.Sprintf("projects/%s/topics/%s", projectID, topic),
		Messages: []*pubsub.PubsubMessage{{Data: b}},
	}

	_, err := pubClient.Publish(context.Background(), req)
	if err != nil {
		return err
	}

	log.Printf("%+v published to %s", fav, topic)

	return nil
}

func parseEvent(e *FirestoreEvent) *c9r.Favorite {
	var fav *c9r.Favorite

	switch {
	case e.Value.Fields.Author.StringValue != "":
		fav = &c9r.Favorite{
			Author: e.Value.Fields.Author.StringValue,
			Site:   e.Value.Fields.Site.StringValue,
			Title:  e.Value.Fields.Title.StringValue,
		}
	// this never happens in case of firestore create
	case e.OldValue.Fields.Author.StringValue != "":
		fav = &c9r.Favorite{
			Author: e.OldValue.Fields.Author.StringValue,
			Site:   e.OldValue.Fields.Site.StringValue,
			Title:  e.OldValue.Fields.Title.StringValue,
		}
	}

	return fav
}
