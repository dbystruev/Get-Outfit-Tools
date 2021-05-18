package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Category is the model for items' categories
type Category struct {
	ID int
	Name string
	//ParentID int
}

// Count is the model for JSON counters
type Count struct {
	Count int
}

// Offer is the model for the item
//type Offer struct {
//	Name string
//}

func isError(err error) bool {
	if err != nil {
		fmt.Println("ERROR:", err)
		return true
	}
	return false
}

// https://stackoverflow.com/a/31129967
func get(url string, target interface{}) error {
	resp, err := http.Get(url)

	if isError(err) { return err }

	defer func() {
		err = resp.Body.Close()
		if isError(err) { return }
	}()

	err = json.NewDecoder(resp.Body).Decode(target)

	if isError(err) { return err }

	return err
}

func main() {
	categories := make([]Category, 0)

	err := get("http://server.getoutfit.ru/categories?limit=-1", &categories)

	if err != nil { return }

	printCategories(categories)
}

func printCategories(categories []Category) {
	fmt.Println("Number of categories:", len(categories))

	offers := 0

	for index, category := range categories {
		offers += printCategory(index, category)
	}

	fmt.Println("Number of offers in categories:", offers)

	counter := Count{}

	if get("http://server.getoutfit.ru/offers?count=true&limit=-1", &counter) != nil { return }

	fmt.Println("Total number of offers:", counter.Count)
}

func printCategory(index int, category Category) int {
	fmt.Print(index, " Category id: ", category.ID, ", name: ", category.Name, ", offers: ")

	counter := Count{}
	url := fmt.Sprintf("http://server.getoutfit.ru/offers?categoryId=%d&count=true&limit=-1", category.ID)

	if get(url, &counter) != nil {
		fmt.Println("none")
		return 0
	}

	fmt.Println(counter.Count)

	return counter.Count
}