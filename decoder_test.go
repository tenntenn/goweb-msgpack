package gowebmsgpack

import (
    "testing"
    "bytes"
    "net/http"
    "net/http/httptest"
    msgpack "github.com/ugorji/go-msgpack"
    "code.google.com/p/goweb/goweb"
)

// Test of MsgpackRequestDecoder.Decode().
func TestDecoderDecode(t *testing.T) {

    expect := []int{1, 2, 3}
    body, _ := msgpack.Marshal(expect, nil)

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

