/*
GVK
Copyright (C) 2023-2024 The GVK Devs

GVK is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published
by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

GVK is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

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
