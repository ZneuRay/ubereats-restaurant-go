package ubereats

import "fmt"

func (c *Client) GetOrderDetail(orderID string) (orderDetail interface{}, err error) {
	if orderID == "" {
		return orderDetail, fmt.Errorf("Invalid orderID.")
	}

	path := orderID

	err = c.Get(path, &orderDetail)
	if err != nil {
		return orderDetail, err
	}

	return
}
