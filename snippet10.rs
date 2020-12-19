let mut raw_fp1 = std::fs::File::create(filename.clone()).expect("Create failed");

//let mut raw_fp2 = tokio::fs::File::create(filename.clone())
//    .await
//    .expect("Create failed");

let mut encoder = GzEncoder::new(raw_fp1, Compression::Default);

while let Some(builder) = rx.recv().await {
    let finished_data = builder.finished_data();
    encoder.write(finished_data).expect("write failed");
    super::POOL.attach(builder) // return builder to the pool
}
