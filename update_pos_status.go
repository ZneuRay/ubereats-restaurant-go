package ubereats

import "fmt"

func (c *Client) UpdatePosStatus(storeID string, posStatusParam interface{}) (response interface{}, err error) {

	if storeID == "" {
		return response, fmt.Errorf("invalid storeID")
	}

	path := fmt.Sprintf("%s/%s", storeID, "pos_data")

	err = c.Patch(path, posStatusParam, response)
	if err != nil {
		return response, err
	}

	return response, err
}
