package tools

type Climan struct {
}

func (Climan) Name() string {
	return "climan"
}

func (Climan) Type() string {
	return string(APIClient)
}

func (Climan) Latest() string {
	return "0.0.1"
}

func (Climan) Current() string {
	return "-"
}

func (Climan) Description() string {
	return "Command LIne tool version MANager "
}

func (Climan) List() []string {
	return []string{"0.0.0", "0.0.1"}
}

func (Climan) Install() error {
	return nil
}

func (Climan) Remove() error {
	return nil
}
