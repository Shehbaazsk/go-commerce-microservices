package models

type CreateUserAddressRequest struct {
	UserID    int32   `json:"user_id" binding:"required"`
	Label     *string `json:"label"`
	Address1  string  `json:"address1" binding:"required"`
	Address2  *string `json:"address2"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	PinCode   string  `json:"pin_code"`
	IsDefault bool    `json:"is_default"`
}

type UpdateUserAddressRequest struct {
	Label     *string `json:"label"`
	Address1  *string `json:"address1"`
	Address2  *string `json:"address2"`
	City      *string `json:"city"`
	State     *string `json:"state"`
	Country   *string `json:"country"`
	PinCode   *string `json:"pin_code"`
	IsDefault *bool   `json:"is_default"`
}
