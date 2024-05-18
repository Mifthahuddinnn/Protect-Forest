package response

import "forest/entities"

type AdminResponse struct {
	ID        int    `json:"id" `
	Username  string `json:"username" `
	CreatedAt string `json:"created_at"`
}

func FromAdmin(admin *entities.Admin) *AdminResponse {
	return &AdminResponse{
		ID:        admin.ID,
		Username:  admin.Username,
		CreatedAt: admin.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
