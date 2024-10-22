package protocol

import "fmt"

func (r *RequestHeader) String() string {
	return fmt.Sprintf("API Key: %d\nAPI Version: %d\nCorrelation ID: %d\nClient Name: %s",
		r.RequestApiKey,
		r.RequestApiVersion,
		r.CorrelationID,
		*r.ClientID)
}

func (r *RequestHeader) Len() int {
	// 10 = static fields count
	length := 10

	if r.Version >= 1 {
		// 1 = nullable string length field
		length += 1 + len(*r.ClientID)
	}

	// TODO: calculate the tagged fields size correctly
	if r.Version >= 2 {
		length += 1
	}

	return length
}
