// protocol has been generated from message format json - DO NOT EDIT
package protocol

type UpdateMetadataResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
}

func (r *UpdateMetadataResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 6 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt16(r.ErrorCode)

	if r.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *UpdateMetadataResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 6 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *UpdateMetadataResponse) GetKey() int16 {
	return 6
}

func (r *UpdateMetadataResponse) GetVersion() int16 {
	return r.Version
}

func (r *UpdateMetadataResponse) GetHeaderVersion() int16 {
	if r.Version >= 6 {
		return 1
	}
	return 0
}

func (r *UpdateMetadataResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 8
}

func (r *UpdateMetadataResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
