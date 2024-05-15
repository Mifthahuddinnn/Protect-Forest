package response

import "forest/entities"

type Approve struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Title  string `json:"title"`
}

func FromUseCaseApprove(approve *entities.Report) *Approve {
	return &Approve{
		ID:     approve.ID,
		Status: approve.Status,
		Title:  approve.Title,
	}
}
