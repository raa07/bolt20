package entity

type Module struct {
	Id                 uint
	IdCourse           uint    `db:"id_course"`
	Title              string `db:"title"`
	ShortDescription   string `db:"short_description"`
	IsAvailableInTrial uint    `db:"is_available_in_trial"` //0 or 1
	Price              uint
}
