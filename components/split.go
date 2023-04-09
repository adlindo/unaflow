package components

type Split struct {
}

func (o *Split) IsAutoExecute() bool {
	return true
}

func (o *Split) Execute() (string, error) {

	return "", nil
}
