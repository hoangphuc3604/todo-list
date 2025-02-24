package util

type creatTodoRes struct {
	ID int `json:"id"`
}

type updateTodoRes struct {
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

func UpdateTodoResponse(id int) *updateTodoRes {
	return &updateTodoRes{
		ID: id,
	}
}

func ErrorResponse(err error) *errorRes {
	return &errorRes{
		Error: err,
	}
}