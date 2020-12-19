func CreateRequest(
    id, method, uri, headers, body []byte) (mr *MarshalledRequest) {
    mr = arPool.Get().(*MarshalledRequest) // HL

    mr.fb.Reset() // HL
    idFB := mr.fb.CreateByteString(id)
    methodFB := mr.fb.CreateByteString(method)
    uriFB := mr.fb.CreateByteString(uri)
    headersFB := mr.fb.CreateByteString(headers)
    bodyFB := mr.fb.CreateByteVector(body)
    fbr.RequestStart(mr.fb)
    fbr.RequestAddId(mr.fb, idFB)
    fbr.RequestAddMethod(mr.fb, methodFB)
    fbr.RequestAddUri(mr.fb, uriFB)
    fbr.RequestAddHeaders(mr.fb, headersFB)
    fbr.RequestAddBody(mr.fb, bodyFB)
    req := fbr.RequestEnd(mr.fb) // HL
    mr.fb.Finish(req) // HL

    return mr
}
