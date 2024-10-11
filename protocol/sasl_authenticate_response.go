// protocol has been generated from message format json - DO NOT EDIT
package protocol

type SaslAuthenticateResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
	// AuthBytes contains the SASL authentication bytes from the server, as defined by the SASL mechanism.
	AuthBytes []byte
	// SessionLifetimeMs contains the SASL authentication bytes from the server, as defined by the SASL mechanism.
	SessionLifetimeMs int64
}

func (r *SaslAuthenticateResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt16(r.ErrorCode)

	if err := pe.putNullableString(r.ErrorMessage); err != nil {
		return err
	}

	if err := pe.putBytes(r.AuthBytes); err != nil {
		return err
	}

	if r.Version >= 1 {
		pe.putInt64(r.SessionLifetimeMs)
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *SaslAuthenticateResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.ErrorMessage, err = pd.getNullableString(); err != nil {
		return err
	}

	if r.AuthBytes, err = pd.getBytes(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.SessionLifetimeMs, err = pd.getInt64(); err != nil {
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

func (r *SaslAuthenticateResponse) GetKey() int16 {
	return 36
}

func (r *SaslAuthenticateResponse) GetVersion() int16 {
	return r.Version
}

func (r *SaslAuthenticateResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *SaslAuthenticateResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *SaslAuthenticateResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
