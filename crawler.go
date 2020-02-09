package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func countPat(s string, pat string) int {
	s = strings.ToLower(s)
	pat = strings.ToLower(pat)
	c := strings.Count(s, pat)
	return c
}
func crawl(depth int, link string, word string, dataCollection map[string]map[string]int) {

	if depth == 0 {
		return
	}
	// get the page body return if this link is not usable
	resp, err := http.Get(link)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(body)

	// store the word count in the page
	wordCount := countPat(bodyString, word)
	if wordCount > 0 {
		dataCollection[link] = map[string]int{
			word:        wordCount,
			"crosslink": 0,
		}
	}
	// find new links in the page
	var re = regexp.MustCompile(`\b(([\w-]+://?|www[.])[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|/)))`)
	for _, newLink := range re.FindAllString(bodyString, -1) {
		if val, ok := dataCollection[newLink]; !ok { // if page is already crawl then don't visit it again and increase the crosslinks count
			crawl(depth-1, newLink, word, dataCollection)
		} else {
			if d, ok := val["crosslink"]; ok {
				val["crosslink"] = d + 1
			}
		}

	}
}
func main() {

	args := os.Args

	link := args[1]
	if !strings.HasPrefix(link, "http://") {
		link = "http://" + link
	}
	depth, _ := strconv.Atoi(args[2])
	word := args[3]

	// start crawling
	dataCollection := make(map[string]map[string]int)
	crawl(depth, link, word, dataCollection)

	// print the collected data
	j, err := json.MarshalIndent(dataCollection, "", "  ")
	if err == nil {
		fmt.Println(string(j))
	}
}
