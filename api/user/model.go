package user

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Birthdate string `json:"birth_date"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
