package service

import (
	"encoding/json"
)

type JsonService interface {
	Format(data string, indent bool) (string, error)
	Validate(data string) bool
	Escape(data string) (string, error)
	Unescape(data string) (string, error)
}

type jsonService struct{}

func NewJsonService() JsonService {
	return &jsonService{}
}

func (s *jsonService) Format(data string, indent bool) (string, error) {
	var obj interface{}
	err := json.Unmarshal([]byte(data), &obj)
	if err != nil {
		return "", err
	}

	var result []byte
	if indent {
		result, err = json.MarshalIndent(obj, "", "  ")
	} else {
		result, err = json.Marshal(obj)
	}

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (s *jsonService) Validate(data string) bool {
	var obj interface{}
	return json.Unmarshal([]byte(data), &obj) == nil
}

func (s *jsonService) Escape(data string) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	// json.Marshal 产生的字符串带有首尾双引号，我们需要去掉它们
	res := string(b)
	if len(res) >= 2 {
		return res[1 : len(res)-1], nil
	}
	return res, nil
}

func (s *jsonService) Unescape(data string) (string, error) {
	var res string
	// 给输入内容包裹双引号，模拟一个 JSON 字符串字面量进行解析
	err := json.Unmarshal([]byte(`"`+data+`"`), &res)
	if err != nil {
		return "", err
	}
	return res, nil
}
