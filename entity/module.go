package entity

type Module struct {
	Id                 int
	IdCourse           int    `db:"id_course"`
	Title              string `db:"title"`
	ShortDescription   string `db:"short_description"`
	IsAvailableInTrial int    `db:"is_available_in_trial"` //0 or 1
	Price              int
}
