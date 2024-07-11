package types

type Author struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Book struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PublisherID int64  `json:"publisher_id"`
	AuthorID    int64  `json:"author_id"`
	IsBorrowed  bool   `json:"is_borrowed"`
}

type Publisher struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Borrower struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user"`
	EndDate   string `json:"end_date"`
	StartDate string `json:"start_date"`
	BookID    int64  `json:"book"`
}

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
