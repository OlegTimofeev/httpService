package main

type fetchTask struct {
	ID      int                 `json:"task_id"`
	Method  string              `json:"method"`
	Path    string              `json:"path"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

type userResponse struct {
	HttpStatus int                 `json:"http_status"`
	Headers    map[string][]string `json:"headers"`
	Body       string              `json:"body"`
	BodyLen    int                 `json:"body_len"`
}
