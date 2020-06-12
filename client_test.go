package ubereats

import (
	"os"
)

func testClient() *Client {
	c := NewClient(os.Getenv("BASE_URL"), os.Getenv("TOKEN"))
	return c
}
