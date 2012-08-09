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

/*

A plugin of goweb, it provides an implementation of msgpack formatter.
It depend on following libraries:

    * goweb (https://code.google.com/p/goweb/)
    * msgpack (https://github.com/ugorji/go-msgpack)



Installation

    go get github.com/tenntenn/goweb-msgpack


Usage

This plugin provide implementations of goweb.Formatter and goweb.RequestDecoder.
You can use MsgpackFormatter as following:

    import gowebmsgpack "github.com/tenntenn/goweb-msgpack"

    //...

    // Add formatter
    msgpackFormatter := new(gowebmsgpack.MsgpackFormatter)
    goweb.AddFormatter(msgpackFormatter)

    // regist handler
    goweb.Map("/sample", handler)

    //...

    // In handler
    func handler(cx *goweb.Context) {
        data := []int{1, 2, 3}

        cx.Format = gowebmsgpack.MSGPACK_FORMAT
        cx.RespondWithData(data)
    }

MsgpackRequestDecoder is used as following:

    // In handler
    func handler(cx *goweb.Context) {
        decoder := new(gowebmsgpack.MsgpackRequestDecoder)
        var v []int
        // decode from cx.Request.Body
        decoder.Decode(cx, v)
    }
*/
package gowebmsgpack
