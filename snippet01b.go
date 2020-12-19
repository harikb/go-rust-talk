
srv := &fasthttp.Server{
	Handler: fastHTTPHandler,
}
err := srv.Serve(_ln)
if err != nil {
	log.Fatalf("http server failed with error: %+v", err)
}
