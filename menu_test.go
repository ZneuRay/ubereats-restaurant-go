package ubereats

import (
	"os"
	"testing"
)

func TestGetMenu(t *testing.T) {

	c := testClient()
	menu, err := c.GetMenu(os.Getenv("STORE_ID"))
	if err != nil {
		t.Error(err)
	}

	if menu == nil {
		t.Error("Api failed on getting menu.")
	}

}
