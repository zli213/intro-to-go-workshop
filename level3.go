package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Level3Handler is a HTTP handler that returns a link to a GIF in the JSON
// format.
// The returned JSON should have the following format:
//
//	{
//		 "gif_url": "URL"
//	}
//
// Step 1: return a static JSON containing a link to your favorite GIF. Use `w`
// to send data back to the client and the `json` package to format your JSON.
// See https://pkg.go.dev/encoding/json (hint: look at NewEncoder or Marshal).\
//
// Step 2: fetch a GIF from Giphy and return it. See the gifURL function below.
//
// Step 3: Get the "query" query parameter from the HTTP request and use it to
// call the gifURL function.
// The request uses the following format:
//
//	/level3?query=search
//
// This means you can get the query and use it in your search!
// The http.Request parameter contains information about the current HTTP
// request, look into r.URL to find the parameter!
type Response struct {
	GifUrl string `json:"gif_url"`
}

func Level3Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	fmt.Println("Query:", query) // 调试输出
	if query == "" {
		// Step 1: Return static JSON
		response := Response{
			GifUrl: "https://user-images.githubusercontent.com/14011726/94132137-7d4fc100-fe7c-11ea-8512-69f90cb65e48.gif",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// Step 2: Fetch a GIF from Giphy
	url, err := gifURL(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		GifUrl: strings.Replace(url, "&ct=g", "", -1),
	}

	json.NewEncoder(w).Encode(response)
}

// Step 2/3 only
// gifURL returns the first GIF returned by the given Giphy search
func gifURL(search string) (string, error) {

	req, err := http.NewRequest(http.MethodGet, "https://api.giphy.com/v1/gifs/search", nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	urlValues := req.URL.Query()
	urlValues.Add("q", search)
	urlValues.Add("limit", "10") // Fetching only one GIF
	urlValues.Add("api_key", "hwyIEPiedVZCoVWbqgOXUgLjlFw1jcqE")
	req.URL.RawQuery = urlValues.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	var result struct {
		Data []struct {
			Images struct {
				FixedWidth struct {
					Webp string `json:"url"`
				} `json:"fixed_width"`
			} `json:"images"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return "", fmt.Errorf("decode error: %v", err)
	}

	if len(result.Data) == 0 {
		return "", fmt.Errorf("no results found")
	}

	return result.Data[0].Images.FixedWidth.Webp, nil
}
