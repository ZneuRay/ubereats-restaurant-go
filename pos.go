package ubereats

import "fmt"

type Pos struct {
	client                *Client
	StoreID               string `json:"store_id"`
	PosIntegrationEnabled bool   `json:"pos_integration_enabled"`
	OnlineStatus          string `json:"online_status"`
	OrderReleaseEnabled   bool   `json:"order_release_enabled"`
	AutoAcceptEnabled     bool   `json:"auto_accept_enabled"`
}

func (c *Client) GetPosStatus(storeID string) (pos *Pos, err error) {
	if storeID == "" {
		return pos, fmt.Errorf("Invalid storeID.")
	}

	path := fmt.Sprintf("%s/%s", storeID, "pos_data")

	err = c.Get(path, &pos)
	if err != nil {
		return pos, err
	}

	if pos != nil {
		pos.client = c
	}

	return
}
