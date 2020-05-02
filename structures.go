package main

type FetchTask struct {
	ID      int                 `json:"task_id"`
	Method  string              `json:"method"`
	Path    string              `json:"path"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

type UserResponse struct {
	FetchTaskId int                 `json:"fetch_task_id"`
	HttpStatus  int                 `json:"http_status"`
	Headers     map[string][]string `json:"headers"`
	BodyLen     int                 `json:"body_len"`
}
