package components

type End struct {
}

func (o *End) IsAutoExecute() bool {
	return true
}

func (o *End) Execute() (string, error) {

	return "", nil
}
