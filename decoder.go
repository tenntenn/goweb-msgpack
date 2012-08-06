package gowebmsgpack

import (
    "io/ioutil"
    "code.google.com/p/goweb/goweb"
    msgpack "github.com/ugorji/go-msgpack"
)

// a Msgpack decoder for request body (just a wrapper to msgpack.Unmarshal)
type MsgpackRequestDecoder struct{}

// Unmarshal msgpack data from request body.
func (d *MsgpackRequestDecoder) Unmarshal(cx *goweb.Context, v interface{}) error {
    // read body
    data, err := ioutil.ReadAll(cx.Request.Body)
    if err != nil {
        return err
    }

    return msgpack.Unmarshal(data, v)
}
