package gowebmsgpack

import (
    "testing"
    "bytes"
    "net/http"
    "net/http/httptest"
    msgpack "github.com/ugorji/go-msgpack"
    "code.google.com/p/goweb/goweb"
)

// Test of gowebmsgpack.MsgpackFormatter.Format().
func TestFormatterFormat(t *testing.T) {

    // Create and regist formatter
    formatter := new(MsgpackFormatter)
    goweb.AddFormatter(formatter)

    // Create context 
    r, _ := http.NewRequest("GET", "http://localhsot:8080", nil)
    w := httptest.NewRecorder()
    pathPrams := goweb.ParameterValueMap(make(map[string]string))
    cx := &goweb.Context{r, w, pathPrams, MSGPACK_FORMAT}

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
    expect, _ := msgpack.Marshal(result, nil)
    cx.RespondWithData(data)
    actual := w.Body.Bytes()
    if !bytes.Equal(expect, actual) {
        t.Errorf("expect %v but actual %v.", expect, actual)
    }
}

// Test of gowebmsgpack.MsgpackFormatter.Match().
func TestFormatterMatch(t *testing.T) {

    // Create and regist formatter
    formatter := new(MsgpackFormatter)
    goweb.AddFormatter(formatter)

    // Create context
    r, _ := http.NewRequest("GET", "http://localhsot:8080", nil)
    w := httptest.NewRecorder()
    pathPrams := goweb.ParameterValueMap(make(map[string]string))
    cx := &goweb.Context{r, w, pathPrams, MSGPACK_FORMAT}

    if !formatter.Match(cx) {
        t.Error("expect formatter.Match returns true but actual is false")
    }
}
