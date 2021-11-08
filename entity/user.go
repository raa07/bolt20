package entity

type User struct {
	Id int
	//Courses []UserCourse
	//CompletedLessons []UserCompletedLesson
}

type UserCompletedLesson struct {
	Id int
	Lesson
}

type UserCourse struct {
	Id int
	Course
}

type UserTask struct {
	Id int
	Task
	Answer string
	Status int
}

type UserTaskConversation struct {
	Task
	UserTask
	UserAuthor User
	Content    string
}
