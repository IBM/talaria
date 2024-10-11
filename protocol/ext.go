package protocol

import "fmt"

func (r *RequestHeader) String() string {
	return fmt.Sprintf("API Key: %d\nAPI Version: %d\nCorrelation ID: %d\nClient Name: %s", r.RequestApiKey, r.RequestApiVersion, r.CorrelationID, *r.ClientID)
}

func (r *RequestHeader) Len() int {
	// 10 = static fields count
	// 1 = nullable string length field
	return 10 + 1 + len(*r.ClientID)
}
