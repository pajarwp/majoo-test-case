package omzet

type MerchantOmzetRequest struct {
	MerchantID  int `json:"merchant_id" validate:"required"`
	Page        int `json:"page" validate:"required"`
	DataPerPage int `json:"data_per_page" validate:"required"`
}

type OutletOmzetRequest struct {
	OutletID    int `json:"outlet_id" validate:"required"`
	Page        int `json:"page" validate:"required"`
	DataPerPage int `json:"data_per_page" validate:"required"`
}

type MerchantOmzet struct {
	Date  string `json:"date"`
	Omzet int    `json:"omzet"`
}

type OutletOmzet struct {
	Date  string `json:"date"`
	Omzet int    `json:"omzet"`
}

type MerchantOmzetResponse struct {
	MerchantName string          `json:"merchant_name"`
	Data         []MerchantOmzet `json:"data"`
	Pagination   Pagination      `json:"pagination"`
}

type OutletOmzetResponse struct {
	MerchantName string        `json:"merchant_name"`
	OutletName   string        `json:"outlet_name"`
	Data         []OutletOmzet `json:"data"`
	Pagination   Pagination    `json:"pagination"`
}

type Pagination struct {
	TotalData   int `json:"total_data"`
	DataPerPage int `json:"data_per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}
