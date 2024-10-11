// protocol has been generated from message format json - DO NOT EDIT
package protocol

type RequestHeader struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// RequestApiKey contains the API key of this request.
	RequestApiKey int16
	// RequestApiVersion contains the API version of this request.
	RequestApiVersion int16
	// CorrelationID contains the correlation ID of this request.
	CorrelationID int32
	// ClientID contains the client ID string.
	ClientID *string
}

func (r *RequestHeader) encode(pe packetEncoder) (err error) {
	nonFlexiblePe := pe
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt16(r.RequestApiKey)

	pe.putInt16(r.RequestApiVersion)

	pe.putInt32(r.CorrelationID)

	if r.Version >= 1 {
		if err := nonFlexiblePe.putNullableString(r.ClientID); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *RequestHeader) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	nonFlexiblePd := pd
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.RequestApiKey, err = pd.getInt16(); err != nil {
		return err
	}

	if r.RequestApiVersion, err = pd.getInt16(); err != nil {
		return err
	}

	if r.CorrelationID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.ClientID, err = nonFlexiblePd.getNullableString(); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *RequestHeader) GetKey() int16 {
	return 0
}

func (r *RequestHeader) GetVersion() int16 {
	return r.Version
}

func (r *RequestHeader) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *RequestHeader) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *RequestHeader) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
