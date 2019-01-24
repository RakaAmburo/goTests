package entities


type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname" binding:"exists,alphanum,min=4,max=255"`
	LastName  string `json:"lastname"`
}