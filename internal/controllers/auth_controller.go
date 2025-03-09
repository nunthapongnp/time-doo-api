package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authClient *auth.Client
}

func NewAuthController(authClient *auth.Client) *AuthController {
	return &AuthController{authClient: authClient}
}

// Request Payload
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// API: For get Firebase ID Token with Email & Password
func (tc *AuthController) GetFirebaseIDToken(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Sign in with Email & Password
	token, err := signInWithFirebase(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"idToken": token})
}

// Sign In via Firebase REST API
func signInWithFirebase(email, password string) (string, error) {
	apiKey := os.Getenv("FIREBASE_API_KEY")
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", apiKey)

	payload := map[string]string{
		"email":             email,
		"password":          password,
		"returnSecureToken": "true",
	}
	jsonData, _ := json.Marshal(payload)

	// HTTP Request to Firebase REST API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read Response
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid login credentials")
	}

	// JSON to Struct
	var result struct {
		IDToken string `json:"idToken"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result.IDToken, nil
}
