pub fn service_fn<F, R, S>(f: F) -> ServiceFn<F, R>
where
    F: FnMut(Request<R>) -> S, // HL
    S: Future, // HL
{
    ServiceFn {
        f,
        _req: PhantomData,
    }
}
