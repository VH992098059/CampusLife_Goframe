// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gerror

import (
	gcode2 "demo3/utility/middleware/gcodeNew"
)

// Code returns the error code.
// It returns CodeNil if it has no error code.
func (err *Error) Code() gcode2.Code {
	if err == nil {
		return gcode2.CodeNil
	}
	if err.code == gcode2.CodeNil {
		return Code(err.Unwrap())
	}
	return err.code
}

// SetCode updates the internal code with given code.
func (err *Error) SetCode(code gcode2.Code) {
	if err == nil {
		return
	}
	err.code = code
}
