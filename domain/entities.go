package domain

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Posts struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
