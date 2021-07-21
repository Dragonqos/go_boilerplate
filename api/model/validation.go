package model

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func ValidatePostChannel(r *http.Request) (*ChannelPostType, map[string]interface{}) {
	var channel ChannelPostType

	rules := govalidator.MapData{
		"name":        []string{"required", "between:3,255"},
	}
	opts := govalidator.Options{
		Request:         r,
		Data:            &channel,
		Rules:           rules,
		RequiredDefault: true, // all the field to be pass the rules
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()

	if len(e) > 0 {
		return nil, map[string]interface{}{"validationError": e}
	}

	return &channel, nil
}