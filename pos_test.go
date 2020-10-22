package ubereats

import (
	"os"
	"testing"
)

func TestGetPosStatus(t *testing.T) {

	c := testClient()
	posStatus, err := c.GetPosStatus(os.Getenv("STORE_ID"))
	if err != nil {
		t.Error(err)
	}

	if posStatus == nil {
		t.Error("Api failed on get pos status.")
	}

}
