Loop:
for {
    select {
    case req, more := <-recordReqChan: // HL
        if !more {
            break Loop
        }
		numRequests++
		err = rf.SaveRequest(req, false) // HL
		if err != nil {
			...
		}
		...
		...
