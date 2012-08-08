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
	"bytes"
	"code.google.com/p/goweb/goweb"
	msgpack "github.com/ugorji/go-msgpack"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test of MsgpackRequestDecoder.Decode().
func TestDecoderDecode(t *testing.T) {

	expect := []int{1, 2, 3}
	body, _ := msgpack.Marshal(expect)

	// Create context 
	r, _ := http.NewRequest("GET", "http://localhsot:8080", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	pathPrams := goweb.ParameterValueMap(make(map[string]string))
	cx := &goweb.Context{r, w, pathPrams, MSGPACK_FORMAT}

	// Create Decoder
	decoder := new(MsgpackRequestDecoder)

	var actual []int
	if err := decoder.Unmarshal(cx, &actual); err != nil {
		t.Error(err)
	}

	if len(expect) != len(actual) {
		t.Errorf("expect length of array is %d but actual is %d.", len(expect), len(actual))
	}

	for i := range expect {
		if expect[i] != actual[i] {
			t.Errorf("expect %dth element is %d but actual is %d", expect[i], actual[i])
			break
		}
	}
}
