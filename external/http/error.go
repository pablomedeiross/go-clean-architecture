package http

type Error struct {
	RequestPath  string `json:"request_path"`
	RequestParms string `json:"request_params"`
	RequestBody  string `json:"request_body"`
	ErrorMsg     string `json:"error"`
}
