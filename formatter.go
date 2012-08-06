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
type MsgpackFormatter struct {}

// Readies response and converts input data into Msgpack.
func (f *MsgpackFormatter) Format(cx *goweb.Context, input interface{}) ([]uint8, error) {
    // marshal msgpack
    output, err := msgpack.Marshal(input)
    if err != nil {
        return nil, err
    }

    cx.ResposeWriter.Header().Set(CONTENT_TYPE, MSGPACK_FORMAT)
}

// Gets the "application/x-msgpack" content type
func (f *MsgpackFormatter) Match(cx *goweb.Context) bool {
   return strings.ToUpper(cx.Format) == MSGPACK_FORMAT
}
