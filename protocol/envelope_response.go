// protocol has been generated from message format json - DO NOT EDIT
package protocol

type EnvelopeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ResponseData contains the embedded response header and data.
	ResponseData []byte
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
}

func (r *EnvelopeResponse) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putBytes(r.ResponseData); err != nil {
		return err
	}

	pe.putInt16(r.ErrorCode)

	pe.putUVarint(0)
	return nil
}

func (r *EnvelopeResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.ResponseData, err = pd.getBytes(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *EnvelopeResponse) GetKey() int16 {
	return 58
}

func (r *EnvelopeResponse) GetVersion() int16 {
	return r.Version
}

func (r *EnvelopeResponse) GetHeaderVersion() int16 {
	return 1
}

func (r *EnvelopeResponse) IsValidVersion() bool {
	return r.Version == 0
}

func (r *EnvelopeResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
