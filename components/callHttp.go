package components

type CallHTTP struct {
}

func (o *CallHTTP) IsAutoExecute() bool {
	return true
}

func (o *CallHTTP) Execute() (string, error) {

	return "", nil
}
