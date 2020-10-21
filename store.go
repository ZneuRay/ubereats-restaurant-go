package ubereats

import "fmt"

type Store struct {
	client      *Client
	Name        string `json:"name"`
	StoreID     string `json:"store_id"`
	AvgPrepTime uint32 `json:"avg_prep_time"`
	Status      string `json:"status"`
}

func (c *Client) GetStore(storeID string) (store *Store, err error) {
	if storeID == "" {
		return store, fmt.Errorf("Invalid storeID.")
	}

	err = c.Get(storeID, &store)
	if err != nil {
		return store, err
	}

	if store != nil {
		store.client = c
	}

	return
}
