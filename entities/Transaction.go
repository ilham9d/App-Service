package entities

type Transaction struct {
	Id         string
	SenderId   string
	ReceiverId string
	Status     string
	Balance    int
	Time       string
}
