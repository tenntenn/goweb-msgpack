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
    "testing"
    "bytes"
    "compress/gzip"
    "net/http"
    "net/http/httptest"
)

func TestWriterFormatterFormat(t *testing.T) {

    // compress writer
    buff := bytes.NewBuffer(make([]byte, 0, 1000))
    flateWriter := gzip.NewWriter(buff)

    // Create and regist formatter
	msgpackFormatter := new(MsgpackFormatter)
    const MsgPackGzip string = "MsgPackGzip"
    formatter := &WriterFormatter{flateWriter, buff, MsgPackGzip, msgpackFormatter}
	goweb.AddFormatter(formatter)

	// Create context 
	r, _ := http.NewRequest("GET", "http://localhsot:8080", nil)
	w := httptest.NewRecorder()
	pathPrams := goweb.ParameterValueMap(make(map[string]string))
	cx := &goweb.Context{r, w, pathPrams, MsgPackGzip}

	data := []int{1, 2, 3}
	result := struct {
		C string
		S int
		D interface{}
		E []string
	}{
		"",
		http.StatusOK,
		data,
		nil,
	}
	expectPacked, _ := msgpack.Marshal(result, nil)
    expectBuff := bytes.NewBuffer(make([]byte, 0, 1000))
    expectFlateWriter := gzip.NewWriter(expectBuff)
    expectFlateWriter.Write(expectPacked)
    expect := expectBuff.Bytes()
	cx.RespondWithData(data)
	actual := w.Body.Bytes()
	if !bytes.Equal(expect, actual) {
		t.Errorf("expect %v but actual %v.", expect, actual)
	}
}

// Test of gowebmsgpack.WriterFormatter.Match().
func TestWriterFormatterMatch(t *testing.T) {

    // compress writer
    buff := bytes.NewBuffer(make([]byte, 0, 1000))
    flateWriter := gzip.NewWriter(buff)

    // Create and regist formatter
	msgpackFormatter := new(MsgpackFormatter)
    const MsgPackGzip string = "MsgPackGzip"
    formatter := &WriterFormatter{flateWriter, buff, MsgPackGzip, msgpackFormatter}
	goweb.AddFormatter(formatter)

	// Create context 
	r, _ := http.NewRequest("GET", "http://localhsot:8080", nil)
	w := httptest.NewRecorder()
	pathPrams := goweb.ParameterValueMap(make(map[string]string))
	cx := &goweb.Context{r, w, pathPrams, MsgPackGzip}

	if !formatter.Match(cx) {
		t.Error("expect formatter.Match returns true but actual is false")
	}
}
