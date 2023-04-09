package components

type Join struct {
}

func (o *Join) IsAutoExecute() bool {
	return true
}

func (o *Join) Execute() (string, error) {

	return "", nil
}
