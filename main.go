/*
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Post represents the structure of a post from the JSONPlaceholder API
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func fetchPost(postID int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Make an HTTP GET request to fetch a post by ID
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching post %d: %v\n", postID, err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body for post %d: %v\n", postID, err)
		return
	}

	// Unmarshal the JSON response into a Post struct
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Printf("Error decoding JSON for post %d: %v\n", postID, err)
		return
	}

	fmt.Printf("Post %d: %s\n", postID, post.Title)
}

func main() {
	// List of post IDs to fetch from JSONPlaceholder API
	postIDs := []int{1, 2, 3, 4, 5}

	// WaitGroup to wait for all Goroutines to finish
	var wg sync.WaitGroup

	// Loop through each post ID and spawn a Goroutine to fetch the post concurrently
	for _, postID := range postIDs {
		wg.Add(1)
		go fetchPost(postID, &wg)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	fmt.Println("API requests completed.")
}
*/

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetchData(url string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %s", url, err)
		return
	}
	defer resp.Body.Close()

	ch <- fmt.Sprintf("URL: %s\nStatus Code: %d", url, resp.StatusCode)
}

func main() {
	urls := []string{"https://www.youtube.com", "https://www.google.com", "https://www.github.com"}

	var wg sync.WaitGroup
	ch := make(chan string)

	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
}
