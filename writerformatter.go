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
    "io"
    "io/ioutil"
    "strings"
	"code.google.com/p/goweb/goweb"
)

// It format by given formatter and write by given writer after formatting.
// For example, given formatter is gowebmsgpack.MsgpackFormatter and writer is gzip.Writer as followings:
// 
// buff := bytes.NewBuffer(make([]byte, 0, 1000))
// flateWriter := gzip.NewWriter(buff)
// msgpackFormatter := new(MsgpackFormatter)
// const MsgPackGzip string = "MsgPackGzip"
// formatter := &WriterFormatter{flateWriter, buff, MsgPackGzip, msgpackFormatter}
// goweb.AddFormatter(formatter)
//
type WriterFormatter struct {
    // It is used for writing data after formatting.
    Writer io.Writer
    // It shoud be connected with WriterFormatter.Writer.
    // It shoud be able to get data which is written by WriterFormatter.Writer.
    // In general, it is bytes.Buffer and it is given to WriterFormatter.Writer initilizer as io.Writer.
    Reader io.Reader
    // It shoud be given to initilizer of *goweb.Context.
    FormatString string
    // It is used for formating data before writting by WriterFormatter.Writer.
    Formatter goweb.Formatter
}

// It is an implementation of goweb.Formatter.Format().
func (wf *WriterFormatter) Format(cx *goweb.Context, input interface{}) ([]uint8, error) {

    // Convert internal format
    data, err := wf.Formatter.Format(cx, input)
    if err != nil {
        return nil, err
    }

    _, err = wf.Writer.Write(data)
    if err != nil {
        return nil, err
    }

    return ioutil.ReadAll(wf.Reader)
}

// It is an implementation of goweb.Formatter.Match().
func (wf *WriterFormatter) Match(cx *goweb.Context) bool {
	return strings.ToUpper(cx.Format) == strings.ToUpper(wf.FormatString)
}
