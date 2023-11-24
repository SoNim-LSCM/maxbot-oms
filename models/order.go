package models

type Order struct {
	OrderID       string `json:"orderId"`
	StartLocation int    `json:"startLocation"`
	OrderTime     int    `json:"orderTime"`
	EndLocation   int    `json:"endLocation"`
	CreatedBy     int    `json:"createdBy"`
}
