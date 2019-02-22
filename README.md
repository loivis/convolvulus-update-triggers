# convvls-update-triggers

+ cloud scheduler as trigger

publish existing documents in books collection

+ cloud firestore onCreate as trigger

publish a single new document from favorites collection

+ deploy cloud functions

```
gcloud functions deploy firestoreTrigger --entry-point FirestoreTrigger --memory 128m \
    --runtime go111 \
    --trigger-event providers/cloud.firestore/eventTypes/document.create \
    --trigger-resource projects/convolvulus/databases/\(default\)/documents/favorites/\{id\} \
    --region asia-northeast1
```
