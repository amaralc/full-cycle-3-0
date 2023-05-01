// Declare the main package
package main

// Import required packages
import (
	"context"
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
		RedirectURL:  "http://localhost:8081/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	// Set a static state value (it's recommended to generate a random state for production)
	state := "foobar" // TODO: generate random state

	// Handle the root route by redirecting the user to the authorization URL
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, config.AuthCodeURL(state), http.StatusFound)
	})

	// Start the HTTP server on port 8081
	log.Fatal(http.ListenAndServe(":8081", nil))

}
