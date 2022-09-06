# zipsa-alarm

### Clone후 초기화 작업
```shell
go get -u <모듈>
go mod vendor
```

### HTTP Trigger
```shell
gcloud functions deploy http-function --entry-point Main --region=asia-northeast3 --runtime go116 --trigger-http
```

### Firestore Trigger
```shell
gcloud functions deploy test-functions --entry-point HelloWorld --region=asia-northeast3 --runtime go116 --trigger-event "providers/cloud.firestore/eventTypes/document.write" --trigger-resource "projects/zipsa-c7baf/databases/(default)/documents/messages/{pushId}"
gcloud functions deploy firestore-function --entry-point Main --region=asia-northeast3 --runtime go116 --trigger-event "providers/cloud.firestore/eventTypes/document.write" --trigger-resource "projects/zipsa-c7baf/databases/(default)/documents/zipsa-alarm/{pushId}"
```
