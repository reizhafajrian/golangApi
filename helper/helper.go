package helper

type Response struct{
	Meta Meta `json:"meta"`
	Data interface{}  `json:"data"`
}
type Meta struct{
	Message string  `json:"message"`
	Code int  `json:"code"`
	Status string  `json:"status"`
}

func APIResponse(Message string, Code int, Status string,Data interface{}) Response{
 meta:=Meta{
 	Message: Message,
 	Code:    Code,
 	Status:  Status,
}
jsonResponse:=Response{
	Meta: meta,
	Data: Data,
}

return jsonResponse

}