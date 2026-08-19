package models

type Execs struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Subject   string `json:"subject,omitempty"`
	Class     string `json:"class,omitempty"`
}
