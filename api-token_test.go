package ubereats

import (
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestClientWithGetToken(t *testing.T) {
	c := testClient()
	token, err := c.GetToken(os.Getenv("LOGIN_URL"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("SCOPE_1"))

	if err != nil {
		t.Error(err)
	}

	if token == nil || token.AccessToken == "" {
		t.Error("Token should no be nil.")
	}
}
