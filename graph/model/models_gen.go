// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Response struct {
	Status *string `json:"status"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AdditionResponse struct {
	Number3 *int `json:"number3"`
}

type CalculateResponse struct {
	Number3 *int `json:"number3"`
}

type Calculation struct {
	Number1 *int    `json:"number1"`
	Number2 *int    `json:"number2"`
	Opertor *string `json:"opertor"`
}

type ChartResponse struct {
	Status *string `json:"status"`
}

type Chartenter struct {
	Cht *string `json:"cht"`
}

type DataFetch struct {
	Inst *string `json:"inst"`
}

type DbInput struct {
	Inp *string `json:"inp"`
}

type Div struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

type DivResponse struct {
	Number3 *int `json:"number3"`
}

type Exeloutput struct {
	Status *string `json:"Status"`
}

type MulResponse struct {
	Number3 *int `json:"number3"`
}

type PieResponse struct {
	Status *string `json:"status"`
}

type Product struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

type StudentResponse struct {
	ID     *string   `json:"id"`
	Status []*string `json:"Status"`
}

type Studententer struct {
	ID         *string `json:"id"`
	Firstname  string  `json:"firstname"`
	Lastname   *string `json:"lastname"`
	Dob        string  `json:"DOB"`
	Phno       *int    `json:"Phno"`
	BloodGroup string  `json:"BloodGroup"`
	Address    string  `json:"Address"`
	Gender     *string `json:"Gender"`
}

type Studentoutput struct {
	ID         *string `json:"id"`
	Firstname  *string `json:"firstname"`
	Lastname   *string `json:"lastname"`
	Dob        *string `json:"dob"`
	Phno       *int    `json:"Phno"`
	BloodGroup *string `json:"BloodGroup"`
	Address    *string `json:"Address"`
	Gender     *string `json:"Gender"`
}

type Sub struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

type SubResponse struct {
	Number3 *int `json:"number3"`
}

type Sum struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}