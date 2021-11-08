package entity

type Task struct {
	Id       int
	IdLesson int `db:"id_lesson"`
	Content  string
	NeedFile int `db:"need_file"` //0 or 1
}
