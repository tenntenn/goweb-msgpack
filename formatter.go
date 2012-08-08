/*
A plugin of goweb, it provides an implementation of msgpack formatter.
https://github.com/tenntenn/goweb-msgpack

Copyright (c) 2012, Takuya Ueda.
All rights reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice,
  this list of conditions and the following disclaimer.
* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.
* Neither the name of the author nor the names of its contributors may be used
  to endorse or promote products derived from this software
  without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package gowebmsgpack

import (
	"code.google.com/p/goweb/goweb"
	msgpack "github.com/ugorji/go-msgpack"
	"strings"
)

// Header key of Content-Type
const CONTENT_TYPE string = "Content-Type"

// Content-Type of msgcpak.
const MSGPACK_CONTENT_TYPE string = "application/x-msgpack"

// Constant string for Msgpack format.
const MSGPACK_FORMAT string = "MSGPACK"

// An implementation goweb.Formatter for msgpack.
type MsgpackFormatter struct{}

// Readies response and converts input data into Msgpack.
func (f *MsgpackFormatter) Format(cx *goweb.Context, input interface{}) ([]uint8, error) {
	// marshal msgpack
	output, err := msgpack.Marshal(input)
	if err != nil {
		return nil, err
	}

	cx.ResponseWriter.Header().Set(CONTENT_TYPE, MSGPACK_FORMAT)

	return output, nil
}

// Gets the "application/x-msgpack" content type
func (f *MsgpackFormatter) Match(cx *goweb.Context) bool {
	return strings.ToUpper(cx.Format) == MSGPACK_FORMAT
}
