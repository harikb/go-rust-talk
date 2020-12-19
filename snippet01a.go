// fastHTTPHandler is the request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {

    if recordReqChan != nil {
        ar := request.CreateRequestFromFastHTTPCtx(ctx) // HL
        recordReqChan <- ar // HL
    }
}
