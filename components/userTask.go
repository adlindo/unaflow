package components

type UserTask struct {
}

func (o *UserTask) IsAutoExecute() bool {
	return false
}

func (o *UserTask) Execute() (string, error) {

	return "", nil
}
