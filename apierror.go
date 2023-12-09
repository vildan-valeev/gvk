package gvk

import (
	"fmt"
)

type APIError struct {
	Code       int64  `json:"error_code"`    // todo const type
	Subcode    int64  `json:"error_subcode"` // todo const type
	Message    string `json:"error_msg"`
	Text       string `json:"error_text"`
	CaptchaSID string `json:"captcha_sid"`
	CaptchaImg string `json:"captcha_img"`
}

// Error returns the error string.
func (a APIError) Error() string {
	return fmt.Sprintf("API error: %d %s. %s", a.Code, a.Message, a.Text)
}

func (a APIError) Base() APIError {
	return a
}
