package traceid

func String() string {
	var traceID [32]byte
	Bytes(traceID[:])
	return string(traceID[:])
}
