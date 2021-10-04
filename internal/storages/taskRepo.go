package storages

type TaskRepo interface {
	Save(u *User)
}

type taskRepo struct {
}

func NewTaskRepo() taskRepo {
	return taskRepo{}
}

func (tr taskRepo) Save(u *Task) *Task {
	db := Get()
	db.Create(u)
	return u
}
