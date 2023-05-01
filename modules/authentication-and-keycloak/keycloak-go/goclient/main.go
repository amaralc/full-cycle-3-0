// Declare the main package
package main

// Import required packages
import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

// Declare global variables for clientID and clientSecret
var (
	clientID     = "myclient"
	clientSecret = "elDB0T9FEJLZPWi5emIsPvPt5r4lnpY7"
)

// Define the main function
func main() {
	// Create a context with a background
	ctx := context.Background()

	// Initialize a new OIDC provider with the specified realm URL
	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/realms/myrealm")
	if err != nil {
		log.Fatal(err)
	}

	// Configure the OAuth2 client with clientID, clientSecret, endpoint, and other settings
	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	// Set a static state value (it's recommended to generate a random state for production)
	state := "foobar" // TODO: generate random state

	// Handle the root route by redirecting the user to the authorization URL
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, config.AuthCodeURL(state), http.StatusFound)
	})

	// Handle the '/auth/callback' route, which is called after the user is redirected from the authorization URL
	http.HandleFunc("/auth/callback", func(writer http.ResponseWriter, request *http.Request) {
		// Check if the state query parameter matches the expected state value
		if request.URL.Query().Get("state") != state {
			// If the state value does not match, respond with an error and a 400 Bad Request status
			http.Error(writer, "Invalid state", http.StatusBadRequest)
			return
		}

		// Attempt to exchange the authorization code received in the query parameter for an access token
		token, err := config.Exchange(ctx, request.URL.Query().Get("code"))
		if err != nil {
			// If there is an error during the token exchange, respond with an error and a 500 Internal Server Error status
			http.Error(writer, "Failed to exchange token", http.StatusInternalServerError)
			return
		}

		// Create a struct containing the access token
		resp := struct {
			AccessToken *oauth2.Token
		}{
			AccessToken: token,
		}

		// Marshal the struct into JSON format
		data, err := json.Marshal(resp)
		if err != nil {
			// If there is an error during JSON marshaling, respond with an error and a 500 Internal Server Error status
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the JSON data to the HTTP response
		writer.Write(data)
	})

	// Log the server start-up message
	log.Default().Println("Server listening on http://localhost:8081")

	// Start the HTTP server on port 8081
	log.Fatal(http.ListenAndServe(":8081", nil))
}
