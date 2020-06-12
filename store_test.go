package ubereats

import (
	"os"
	"testing"
)

func TestGetStoreInfo(t *testing.T) {

	c := testClient()
	store, err := c.GetStore(os.Getenv("STORE_ID"))
	if err != nil {
		t.Error(err)
	}

	if store == nil {
		t.Error("Api failed on get store info.")
	}

}
