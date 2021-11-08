package entity

type Lesson struct {
	Id                 int
	IdModule           int `db:"id_module"`
	Title              string
	ShortDescription   string `db:"short_description"`
	Content            string
	IsAvailableInTrial int `db:"is_available_in_trial"` //0 or 1
}
