package server

type CreateServerRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Environment string `json:"environment"`
}

type ServerResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Environment string `json:"environment"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
