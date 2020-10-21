package ubereats

import (
	"os"
	"testing"
)

func TestGetPosStatus(t *testing.T) {

	c := testClient()
	pos, err := c.GetPosStatus(os.Getenv("STORE_ID"))
	if err != nil {
		t.Error(err)
	}

	if pos == nil {
		t.Error("Api failed on get pos status.")
	}

}
