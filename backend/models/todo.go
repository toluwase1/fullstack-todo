package models

const (
	StatusPending    = "pending"
	InProgress       = "in-progress"
	StatusDone       = "completed"
	CategoryPersonal = "personal"
	CategoryWork     = "work"
	CategoryHome     = "home"
)

type Todo struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Status      string `json:"status"`
}

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Status      string `json:"status"`
}
