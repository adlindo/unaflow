package components

type Condition struct {
}

func (o *Condition) IsAutoExecute() bool {
	return true
}

func (o *Condition) Execute() (string, error) {

	return "", nil
}
