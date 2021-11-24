package entity

type Course struct {
	Id               uint
	Title            string
	ShortDescription string `db:"short_description"`
	LongDescription  string `db:"long_description"`
	Price            int
}
