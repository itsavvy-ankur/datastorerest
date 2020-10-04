# Golang DataStore example

* Connects to a [Datastore emulator](https://cloud.google.com/datastore/docs/tools/datastore-emulator)
* Adds a new task and then retrieves it

## TODO 
* Learn rest api's on Golang and implement !

## Execution Details

* Install gcloud beta emulators datastore
* ```shell
  gcloud beta emulators datastore start
  ```
* Change project id in `main.go`
```golang
    // Set your Google Cloud Platform project ID.
    projectID := "aqueous-tube-252313"
```

* Execute using
  ```golang
  go run main.go
  ```