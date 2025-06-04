package api

type Error struct {
	Message string `json:"message"`
}

type Number struct {
	ID        int    `json:"id"`
	Timestamp string `json:"timestamp"`
}

type NewNumberRequest struct {
	ID int `json:"id"`
}
