package main

type fetchTask struct {
	taskID  int
	method  string
	path    string
	headers map[string][]string
	body    string
}

type userResponse struct {
	httpStatus int
	headers    map[string]string
	bodyLen    int
}
