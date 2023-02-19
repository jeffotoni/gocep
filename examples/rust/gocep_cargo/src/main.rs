pub(crate) use hyper::{body::HttpBody as _, Client, Uri};
let client = Client::new();
let res = client.get(Uri::from_static("http://localhost:8080/api/v1/08226021")).await?;
println!("status: {}", res.status());
let buf = hyper::body::to_bytes(res).await?;
println!("body: {:?}", buf);