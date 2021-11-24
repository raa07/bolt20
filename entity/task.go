package entity

type Task struct {
	Id       int
	IdLesson uint `db:"id_lesson"`
	Content  string
	NeedFile uint `db:"need_file"` //0 or 1
}
