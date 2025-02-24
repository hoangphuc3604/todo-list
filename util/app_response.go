package util

type creatTodoRes struct {
	ID int `json:"id"`
}

type errorRes struct {
	Error error `json:"error"`
}

func CreateTodoResponse(id int) *creatTodoRes {
	return &creatTodoRes{
		ID: id,
	}
}

func ErrorResponse(err error) *errorRes {
	return &errorRes{
		Error: err,
	}
}