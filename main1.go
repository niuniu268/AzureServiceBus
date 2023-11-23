package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/clientcredentials"
	`io`
	"io/ioutil"
	"net/http"
)

const (
	tokenURL = "https://login.microsoftonline.com/8008c85d-baa3-47ff-a30f-63d8142742a9/oauth2/v2.0/token"
	clientID = "acb53893-0c84-46ce-81ca-0011d388ba08"
	// Update with the newly generated client secret
	clientSecret = "your-new-client-secret-value"
	// Replace with your actual Application ID URI
	applicationIDURI = "api://acb53893-0c84-46ce-81ca-0011d388ba08"
)

func main() {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		Scopes:       []string{applicationIDURI + "/.default"},
	}

	token, err := config.Token(context.Background())
	if err != nil {
		fmt.Printf("Error getting token: %v\n", err)
		return
	}

	fmt.Printf("Access Token: %s\n", token.AccessToken)

	// Now you can use the token to make authenticated requests to your API.
	// For example, using the token in a request to a secured resource:
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://your-api/resource", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("Response: %s\n", body)
}
