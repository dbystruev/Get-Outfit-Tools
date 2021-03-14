# Get Outfit Tools

Utilities for working with Get Outfit server written in [Go](https://golang.org).

* [categories.go](categories.go) Get the list of categories from the [server](https://github.com/dbystruev/Get-Outfit-Server.git), then get the number of offers in each category.
* [cache.go](cache.go) Request items for female, male, and non-binary categories from the [server](https://github.com/dbystruev/Get-Outfit-Server.git) to pre-cache [client app](https://github.com/dbystruev/Outfit-Selection.git)'s requests.

## Run
```bash
go run categories.go
go run cache.go
```
