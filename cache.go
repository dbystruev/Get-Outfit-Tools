package main

import (
	"fmt"
	"net/http"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println("ERROR:", err)
		return true
	}
	return false
}

func get(url string) int64 {
	resp, err := http.Get(url)

	if isError(err) { return 0 }

	defer func() {
		err = resp.Body.Close()
		if isError(err) { return }
	}()

	return resp.ContentLength
}

func main() {
	expectedLengths := []int64{15185148, 15083397, 15179689}

	requests := []string{
		"http://server.getoutfit.ru/offers?vendor=a-cold-wall&vendor=acne%20studios&vendor=adidas&vendor=alexander_mcqueen&vendor=alexander_wang&vendor=balenciaga&vendor=balmain&vendor=brunello%20cucinelli&vendor=burberry&vendor=calvin%20klein&vendor=chanel&vendor=chlo&vendor=common_projects&vendor=diesel&vendor=dolce%20%26%20gabbana&vendor=dr_martens&vendor=fendi&vendor=giorgio%20armani&vendor=gucci&vendor=herme%CC%80s&vendor=heron_preston&vendor=jacquemus&vendor=jil%20sander&vendor=jimmy%20choo&vendor=loewe&vendor=louis%20vuitton&vendor=maison_margiela&vendor=marni&vendor=msgm&vendor=nike&vendor=off-white&vendor=palm_angels&vendor=philipp%20plein&vendor=prada&vendor=ralph%20lauren&vendor=rick_owens&vendor=saint_laurent&vendor=sandro&vendor=stone%20island&vendor=supreme&vendor=tom_ford&vendor=tommy%20hilfiger&vendor=valentino&vendor=versace&vendor=vetements&categoryId=135983&categoryId=135981&categoryId=135985&categoryId=136043&categoryId=136337&categoryId=136338&categoryId=136227&categoryId=136245&categoryId=136335&categoryId=136336&categoryId=135998&categoryId=135999&categoryId=136032&categoryId=136033&categoryId=136034&categoryId=136035&categoryId=136054&categoryId=136055&categoryId=136057&categoryId=136058&categoryId=136059&categoryId=136644&categoryId=139874&categoryId=139875&categoryId=136302&categoryId=136307&categoryId=136353&categoryId=136358&categoryId=136467&categoryId=136485&%D0%BF%D0%BE%D0%BB=%D0%B6%D0%B5%D0%BD%D1%81%D0%BA%D0%B8%D0%B9&limit=10000",
		"http://server.getoutfit.ru/offers?limit=10000&%D0%BF%D0%BE%D0%BB=%D0%BC%D1%83%D0%B6%D1%81%D0%BA%D0%BE%D0%B9&vendor=a-cold-wall&vendor=acne%20studios&vendor=adidas&vendor=alexander_mcqueen&vendor=alexander_wang&vendor=balenciaga&vendor=balmain&vendor=brunello%20cucinelli&vendor=burberry&vendor=calvin%20klein&vendor=chanel&vendor=chlo&vendor=common_projects&vendor=diesel&vendor=dolce%20%26%20gabbana&vendor=dr_martens&vendor=fendi&vendor=giorgio%20armani&vendor=gucci&vendor=herme%CC%80s&vendor=heron_preston&vendor=jacquemus&vendor=jil%20sander&vendor=jimmy%20choo&vendor=loewe&vendor=louis%20vuitton&vendor=maison_margiela&vendor=marni&vendor=msgm&vendor=nike&vendor=off-white&vendor=palm_angels&vendor=philipp%20plein&vendor=prada&vendor=ralph%20lauren&vendor=rick_owens&vendor=saint_laurent&vendor=sandro&vendor=stone%20island&vendor=supreme&vendor=tom_ford&vendor=tommy%20hilfiger&vendor=valentino&vendor=versace&vendor=vetements&categoryId=136332&categoryId=135981&categoryId=136043&categoryId=136337&categoryId=136338&categoryId=136226&categoryId=136227&categoryId=136331&categoryId=136334&categoryId=136335&categoryId=136336&categoryId=136006&categoryId=136323&categoryId=139874&categoryId=139875&categoryId=136302&categoryId=136303&categoryId=136306&categoryId=136353&categoryId=136354&categoryId=136359",
		"http://server.getoutfit.ru/offers?vendor=a-cold-wall&vendor=acne%20studios&vendor=adidas&vendor=alexander_mcqueen&vendor=alexander_wang&vendor=balenciaga&vendor=balmain&vendor=brunello%20cucinelli&vendor=burberry&vendor=calvin%20klein&vendor=chanel&vendor=chlo&vendor=common_projects&vendor=diesel&vendor=dolce%20%26%20gabbana&vendor=dr_martens&vendor=fendi&vendor=giorgio%20armani&vendor=gucci&vendor=herme%CC%80s&vendor=heron_preston&vendor=jacquemus&vendor=jil%20sander&vendor=jimmy%20choo&vendor=loewe&vendor=louis%20vuitton&vendor=maison_margiela&vendor=marni&vendor=msgm&vendor=nike&vendor=off-white&vendor=palm_angels&vendor=philipp%20plein&vendor=prada&vendor=ralph%20lauren&vendor=rick_owens&vendor=saint_laurent&vendor=sandro&vendor=stone%20island&vendor=supreme&vendor=tom_ford&vendor=tommy%20hilfiger&vendor=valentino&vendor=versace&vendor=vetements&limit=10000&categoryId=135983&categoryId=135981&categoryId=135985&categoryId=136043&categoryId=136337&categoryId=136338&categoryId=136227&categoryId=136245&categoryId=136335&categoryId=136336&categoryId=135998&categoryId=135999&categoryId=136032&categoryId=136033&categoryId=136034&categoryId=136035&categoryId=136054&categoryId=136055&categoryId=136057&categoryId=136058&categoryId=136059&categoryId=136644&categoryId=139874&categoryId=139875&categoryId=136302&categoryId=136307&categoryId=136353&categoryId=136358&categoryId=136467&categoryId=136485&categoryId=136332&categoryId=135981&categoryId=136043&categoryId=136337&categoryId=136338&categoryId=136226&categoryId=136227&categoryId=136331&categoryId=136334&categoryId=136335&categoryId=136336&categoryId=136006&categoryId=136323&categoryId=139874&categoryId=139875&categoryId=136302&categoryId=136303&categoryId=136306&categoryId=136353&categoryId=136354&categoryId=136359",
	}

	for index, request := range requests {
		expectedLength := expectedLengths[index]
		length := get(request)
		status := " OK"
		if expectedLength != length {
			status = fmt.Sprintf(", expected %d bytes", expectedLength)
		}
		fmt.Printf("Request %d: loaded %d bytes%s\n", index, length, status)
	}
}