package tasks_domain

type Task struct {
	Name   string `validate:"required"`
	Status string `validate:"required"`
}

type TaskMongo struct {
	Name   string `bson:"name,omitempty"`
	Status string `bson:"status,omitempty"`
}

func (t *Task) ToMongo() TaskMongo {
	return TaskMongo{
		Name:   t.Name,
		Status: t.Status,
	}
}
