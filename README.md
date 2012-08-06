goweb-msgpack
=============

A plugin of goweb, it provides an implementation of msgpack formatter.
It depend on following libraries:
    + [goweb](https://code.google.com/p/goweb/)
    + [msgpack](https://github.com/ugorji/go-msgpack)

How to install
-------------

    go get github.com/tenntenn/goweb-msgpack

How to use
-------------

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
