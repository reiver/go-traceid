package traceid

import (
	"time"

	"github.com/reiver/go-erorr"
)

func Bytes(traceID []byte) error {
	unixTime := uint64(time.Now().Unix())

	return bytes(traceID, unixTime)
}


func bytes(traceID []byte, unixTime uint64) error {
	{
		const min = 32

		length := len(traceID)

		if length < min {
			return erorr.Errorf("traceid: cannot generate trade-id into slice with length less than %d — slice has length %d", min, length)
		}
	}

	// 1111_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000
	traceID[0]  = defaultTraceIDCharacters[ (0xf000000000000000 & unixTime) >> 58]

	// 0000_1111_1100_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000
	traceID[1]  = defaultTraceIDCharacters[ (0x0fc0000000000000 & unixTime) >> 52]

	// 0000_0000_0011_1111_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000
	traceID[2]  = defaultTraceIDCharacters[ (0x003f000000000000 & unixTime) >> 46]

	// 0000_0000_0000_0000_1111_1100_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000
	traceID[3]  = defaultTraceIDCharacters[ (0x0000fc0000000000 & unixTime) >> 40]

	// 0000_0000_0000_0000_0000_0011_1111_0000_0000_0000_0000_0000_0000_0000_0000_0000
	traceID[4]  = defaultTraceIDCharacters[ (0x000003f000000000 & unixTime) >> 36]

	// 0000_0000_0000_0000_0000_0000_0000_1111_1100_0000_0000_0000_0000_0000_0000_0000
	traceID[5]  = defaultTraceIDCharacters[ (0x0000000fc0000000 & unixTime) >> 30]

	// 0000_0000_0000_0000_0000_0000_0000_0000_0011_1111_0000_0000_0000_0000_0000_0000
	traceID[6]  = defaultTraceIDCharacters[ (0x000000003f000000 & unixTime) >> 24]

	// 0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_1111_1100_0000_0000_0000_0000
	traceID[7]  = defaultTraceIDCharacters[ (0x0000000000fc0000 & unixTime) >> 18]

	// 0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0011_1111_0000_0000_0000
	traceID[8]  = defaultTraceIDCharacters[ (0x000000000003f000 & unixTime) >> 12]

	// 0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_1111_1100_0000
	traceID[9]  = defaultTraceIDCharacters[ (0x0000000000000fc0 & unixTime) >> 6]

	// 0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0011_1111
	traceID[10] = defaultTraceIDCharacters[ (0x000000000000003f & unixTime)]

	const offset = 10 // <---- This needs to match the last index manually assigned to in the code before it.

	length := len(defaultTraceIDCharacters)
	for i,_ := range traceID[offset:] {
		traceID[offset+i] = defaultTraceIDCharacters[randomness.Intn(length)]
	}

	return nil
}
