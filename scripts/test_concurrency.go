package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// Dhyan de: Ye post ID aapke database me already exist karni chahiye!
	// Agar aapke paas post id 13 nahi hai, toh apni DB se koi id replace kar lena.
	postID := 2
	url := fmt.Sprintf("http://localhost:8080/v1/posts/%d", postID)

	var wg sync.WaitGroup
	wg.Add(2) // Hum 2 requests bhejenge ek sath

	// User A - Ye sirf Title update karega
	go func() {
		defer wg.Done()
		updatePost(url, `{"title": "title from User A"}`)
	}()

	// User B - Ye sirf Content update karega
	go func() {
		defer wg.Done()
		updatePost(url, `{"content": "content from User B"}`)
	}()

	// Dono requests ke complete hone ka wait karenge
	wg.Wait()
	fmt.Println("Both requests completed. Check your database to see the result!")
}

func updatePost(url, payload string) {
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Sent payload: %-35s | Received Status: %s\n", payload, resp.Status)
}
