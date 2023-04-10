package components

type UserTask struct {
	BaseComponent
}

func (o *UserTask) IsAutoNext() bool {
	return false
}
