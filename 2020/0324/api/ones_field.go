package api

import (
	"fmt"
	"net/http"

	"github.com/bangwork/mygo/2020/0324/config"
)

var customerFieldUUID = "JrvswW8P"

type FieldRequest struct {
	Field *FieldPayload `json:"field"`
}

type FieldResponse struct {
	Field             *FieldPayload `json:"field"`
	ServerUpdateStamp int64         `json:"server_update_stamp"`
}

type ServerUpdateStampResponse struct {
	ServerUpdateStamp int64 `json:"server_update_stamp"`
}

type ListFieldsResponse struct {
	Fields            []*FieldPayload `json:"fields"`
	ServerUpdateStamp int64           `json:"server_update_stamp"`
}

type FieldProjectPayload struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type FieldPayload struct {
	UUID         string                `json:"uuid"`
	Name         *string               `json:"name"`
	NamePinyin   string                `json:"name_pinyin"`
	Desc         *string               `json:"desc"`
	Type         *int                  `json:"type"`
	Renderer     *int                  `json:"renderer"`
	FilterOption *int                  `json:"filter_option"`
	SearchOption *int                  `json:"search_option"`
	CreateTime   *int64                `json:"create_time"`
	BuiltIn      bool                  `json:"built_in"`
	Fixed        bool                  `json:"fixed"`
	Options      []*FieldOptionPayload `json:"options"`
}

type FieldOptionPayload struct {
	UUID            string `json:"uuid,omitempty"`
	Value           string `json:"value,omitempty"`
	Selected        bool   `json:"selected,omitempty"`
	BackgroundColor string `json:"background_color,"`
	Color           string `json:"color"`
	Desc            string `json:"desc,omitempty"`
}

func GetONESCustomerField() (*FieldPayload, error) {
	_path := fmt.Sprintf("/project/v1/team/%s/fields", teamUUID)
	client := NewClient(config.Config.ONESBaseURL)
	resp := new(ListFieldsResponse)
	_, err := client.Request(http.MethodGet, _path, nil, resp)
	if err != nil {
		return nil, err
	}
	for _, field := range resp.Fields {
		if field.UUID == customerFieldUUID {
			return field, nil
		}
	}
	return nil, nil
}

func AddCustomerField(fieldName string) error {
	filed, err := GetONESCustomerField()
	if err != nil {
		return err
	}
	if filed == nil {
		return nil
	}
	_path := fmt.Sprintf("/project/v1/team/%s/field/%s/update", teamUUID, filed.UUID)
	client := NewClient(config.Config.ONESBaseURL)
	option := &FieldOptionPayload{
		Value:           fieldName,
		BackgroundColor: "#e63422",
		Color:           "#fff",
	}
	filed.Options = append(filed.Options, option)
	resp := new(FieldResponse)
	req := &FieldRequest{
		Field: filed,
	}
	_, err = client.Request(http.MethodPost, _path, req, resp)
	if err != nil {
		return err
	}
	return nil
}
