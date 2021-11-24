package entity

type Lesson struct {
	Id                 uint
	IdModule           uint `db:"id_module"`
	Title              string
	ShortDescription   string `db:"short_description"`
	Content            string
	IsAvailableInTrial uint `db:"is_available_in_trial"` //0 or 1
}
