// Request returns a pointer to the fbr.Request represented by the UnmarshalledRequest
func (umr *UnmarshalledRequest) Request() *fbr.Request {
	return fbr.GetRootAsRequest(umr.data, 0) // HL
}
