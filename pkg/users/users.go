package users

// User represents an application user.
type User struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Manager is the interface for an application user.
type Manager interface {
	GetUser(string) (*User, error)
}
