package errors

type UnauthorizedOperationError struct{}

func (unauthorizedOperationError *UnauthorizedOperationError) Error() string {
	return ""
}
