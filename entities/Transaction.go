package entities

type Transaction struct {
	Id         string
	SenderId   string
	ReceiverId string
	Status     string
	Balance    int
	Time       string
}

type Transfer struct {
	Pengirim      string
	Penerima      string
	PhonePengirim string
	PhonePenerima string
	Balance       int
	Time          string
}
