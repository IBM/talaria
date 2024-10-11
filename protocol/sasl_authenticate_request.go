// protocol has been generated from message format json - DO NOT EDIT
package protocol

type SaslAuthenticateRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// AuthBytes contains the SASL authentication bytes from the client, as defined by the SASL mechanism.
	AuthBytes []byte
}

func (r *SaslAuthenticateRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putBytes(r.AuthBytes); err != nil {
		return err
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *SaslAuthenticateRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.AuthBytes, err = pd.getBytes(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *SaslAuthenticateRequest) GetKey() int16 {
	return 36
}

func (r *SaslAuthenticateRequest) GetVersion() int16 {
	return r.Version
}

func (r *SaslAuthenticateRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *SaslAuthenticateRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *SaslAuthenticateRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
