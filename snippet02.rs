let make_svc = make_service_fn(|_conn| {
    let (tx, rx) = tokio::sync::mpsc::channel(1000);

    let output_directory = output_directory.to_string();

    task::spawn(async move { attempt4::recorder(output_directory, rx).await });

    // tx is now a separate clone for each instance of http-connection
    async move {
        Ok::<_, Infallible>(service_fn(move |req: Request<Body>| {
            attempt4::handle(req, record, tx.clone())
        }))
    }
});

let server = Server::bind(&addr).serve(make_svc);
