# Get Outfit Tools

Utilities for working with Get Outfit server written in [Go](https://golang.org).

* [acoola_categories.go](acoola_categories.go) Get the list of Acoola categories from the [server](https://github.com/dbystruev/Get-Outfit-Server.git), then get the number of offers in each category.
* [farfetch_categories.go](farfetch_categories.go) Get the list of Farfetch categories from the [server](https://github.com/dbystruev/Get-Outfit-Server.git), then get the number of offers in each category.
* [cache.go](cache.go) Request items for female, male, and non-binary categories from the [server](https://github.com/dbystruev/Get-Outfit-Server.git) to pre-cache [client app](https://github.com/dbystruev/Outfit-Selection.git)'s requests.

## Run
```bash
go run acoola_categories.go
go run farfetch_categories.go
go run cache.go
```
