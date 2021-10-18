package entity

type Course struct {
	Id int
	Title string
	ShortDescription string `db:"short_description"`
	LongDescription string `db:"long_description"`
	Price int
}
