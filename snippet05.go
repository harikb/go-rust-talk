func fastHTTPHandler(ctx *fasthttp.RequestCtx) {

    if recordReqChan != nil {
        ar := request.CreateRequestFromFastHTTPCtx(ctx) // HL
        recordReqChan <- ar
    }
}
