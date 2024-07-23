package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
)

const (
	clientID     = "client-id"
	clientSecret = "client-secret"
	redirectURL  = "http://localhost:8080/callback"
	serverURL    = "http://localhost:8080"
)

var (
	oauthConfig *oauth2.Config
	privateKey  *rsa.PrivateKey
	publicKey   *rsa.PublicKey
	authCodes   = make(map[string]string)
)

func main() {
	// Generate RSA key pair for JWT signing and verification
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey = &privateKey.PublicKey

	// Set up OAuth2 configuration
	oauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  serverURL + "/auth",
			TokenURL: serverURL + "/token",
		},
	}

	// Set up HTTP server
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/auth", handleAuth)
	http.HandleFunc("/token", handleToken)

	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="/login">Log in with Custom OAuth2</a></body></html>`
	fmt.Fprint(w, html)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := oauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Verify the JWT
	parsedToken, err := jwt.Parse(token.AccessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil || !parsedToken.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Extract claims from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid claims", http.StatusUnauthorized)
		return
	}

	// Display user information
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(claims)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	state := r.URL.Query().Get("state")
	responseType := r.URL.Query().Get("response_type")

	// Validate parameters
	if clientID != oauthConfig.ClientID {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}
	if redirectURI != oauthConfig.RedirectURL {
		http.Error(w, "Invalid redirect URI", http.StatusBadRequest)
		return
	}
	if responseType != "code" {
		http.Error(w, "Invalid response type", http.StatusBadRequest)
		return
	}

	// Generate a random authorization code
	authCode, err := generateRandomString(32)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Store the authorization code (in a real scenario, associate it with the user)
	authCodes[authCode] = clientID

	// Redirect back to the client with the auth code
	redirectURL := fmt.Sprintf("%s?code=%s&state=%s", redirectURI, authCode, state)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func handleToken(w http.ResponseWriter, r *http.Request) {
	// This should be a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Extract form parameters
	grantType := r.Form.Get("grant_type")
	authCode := r.Form.Get("code")
	clientID := r.Form.Get("client_id")
	clientSecret := r.Form.Get("client_secret")

	// Validate parameters
	if grantType != "authorization_code" {
		http.Error(w, "Invalid grant type", http.StatusBadRequest)
		return
	}
	if clientID != oauthConfig.ClientID || clientSecret != oauthConfig.ClientSecret {
		http.Error(w, "Invalid client credentials", http.StatusUnauthorized)
		return
	}

	// Verify the authorization code
	storedClientID, valid := authCodes[authCode]
	if !valid || storedClientID != clientID {
		http.Error(w, "Invalid authorization code", http.StatusBadRequest)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": "user123", // In a real scenario, this would be the authenticated user's ID
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	response := struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   86400, // 24 hours
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// Remove the used authorization code
	delete(authCodes, authCode)
}

// Helper function to generate a random string
func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}
