package decode

import "encoding/json"

type VerifyMailTopic[T interface{}] struct {
	Type  string
	Value T
}

type VerifyMailTopicData struct {
	VerifyUrl string   `json:"verifyUrl"`
	To        []string `json:"to"`
	Lang      string   `json:"lang"`
}

func SendMailTopicConverter[T interface{}](input []byte, output *T) (*T, error) {
	err := json.Unmarshal(input, output)
	if err != nil {
		return output, err
	}
	return output, nil
}
