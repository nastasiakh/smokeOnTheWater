package models

type OrderWithProducts struct {
	Order         Order           `json:"order"`
	OrderProducts []*OrderProduct `json:"orderProducts"`
}
