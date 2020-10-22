package ubereats

import "fmt"

type PosStatus struct {
	client                *Client
	StoreID               string `json:"store_id"`
	PosIntegrationEnabled bool   `json:"pos_integration_enabled"`
	OnlineStatus          string `json:"online_status"`
	OrderReleaseEnabled   bool   `json:"order_release_enabled"`
	AutoAcceptEnabled     bool   `json:"auto_accept_enabled"`
}

func (c *Client) GetPosStatus(storeID string) (posStatus *PosStatus, err error) {
	if storeID == "" {
		return posStatus, fmt.Errorf("invalid storeID")
	}

	path := fmt.Sprintf("%s/%s", storeID, "pos_data")

	err = c.Get(path, &posStatus)
	if err != nil {
		return posStatus, err
	}

	return posStatus, err
}
