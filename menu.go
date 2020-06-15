package ubereats

import "fmt"

func (c *Client) GetMenu(storeID string) (menu interface{}, err error) {
	if storeID == "" {
		return menu, fmt.Errorf("Invalid storeID.")
	}

	path := fmt.Sprintf("%s/%s", storeID, "menus")

	err = c.Get(path, &menu)
	if err != nil {
		return menu, err
	}

	return
}
