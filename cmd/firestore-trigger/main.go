package main

import (
	"context"
	"log"

	convvls "github.com/loivis/convolvulus-update-triggers"
)

func main() {
	e := convvls.FirestoreEvent{
		Value: convvls.FirestoreValue{
			Fields: convvls.FireStoreData{
				Author: convvls.FireStoreStringValue{StringValue: "foo"},
				Site:   convvls.FireStoreStringValue{StringValue: "bar"},
				Title:  convvls.FireStoreStringValue{StringValue: "baz"},
			},
		},
	}

	if err := convvls.FirestoreTrigger(context.Background(), e); err != nil {
		log.Fatal(err)
	}
}
