package helper

type successRespond struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessRespond(data, paging, filter interface{}) *successRespond {
	return &successRespond{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessRespond(data interface{}) *successRespond {
	return &successRespond{data, nil, nil}
}
