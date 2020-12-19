	
func (rf *Archive) SaveRequest(req *request.MarshalledRequest, flushNow bool) (err error) {
    ...
    fbBytes := req.Bytes()
	fbLen := len(fbBytes)

	defer req.Release() // HL

	binary.LittleEndian.PutUint64(lbuf, uint64(fbLen))

	n, err := rf.Write(lbuf)
	if err != nil {
        ...
	}
	n, err = rf.Write(fbBytes) // HL
	if err != nil {
        ...
	}
    ...
    ...
}