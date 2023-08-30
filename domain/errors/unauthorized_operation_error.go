package errors

type UnauthorizedOperationError struct{}

func NewUnauthorizedOperationError() *UnauthorizedOperationError {
	return &UnauthorizedOperationError{}
}

func (unauthorizedOperationError *UnauthorizedOperationError) Error() string {
	return ""
}
