// protocol has been generated from message format json - DO NOT EDIT
package protocol

type ResponseHeader struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// CorrelationID contains the correlation ID of this response.
	CorrelationID int32
}

func (r *ResponseHeader) encode(pe packetEncoder) (err error) {
	if r.Version >= 1 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.CorrelationID)

	if r.Version >= 1 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ResponseHeader) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 1 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.CorrelationID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ResponseHeader) GetKey() int16 {
	return 0
}

func (r *ResponseHeader) GetVersion() int16 {
	return r.Version
}

func (r *ResponseHeader) GetHeaderVersion() int16 {
	if r.Version >= 1 {
		return 1
	}
	return 0
}

func (r *ResponseHeader) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *ResponseHeader) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
