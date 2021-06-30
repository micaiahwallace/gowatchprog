package gowatchprog

type InvalidContextError string

func (e InvalidContextError) Error() string {
	return string(e)
}

const ErrInvalidContext InvalidContextError = "Unknown application context"
