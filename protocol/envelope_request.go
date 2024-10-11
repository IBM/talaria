// protocol has been generated from message format json - DO NOT EDIT
package protocol

type EnvelopeRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// RequestData contains the embedded request header and data.
	RequestData []byte
	// RequestPrincipal contains a Value of the initial client principal when the request is redirected by a broker.
	RequestPrincipal []byte
	// ClientHostAddress contains the original client's address in bytes.
	ClientHostAddress []byte
}

func (r *EnvelopeRequest) encode(pe packetEncoder) (err error) {
	pe = FlexibleEncoderFrom(pe)
	if err := pe.putBytes(r.RequestData); err != nil {
		return err
	}

	if err := pe.putBytes(r.RequestPrincipal); err != nil {
		return err
	}

	if err := pe.putBytes(r.ClientHostAddress); err != nil {
		return err
	}

	pe.putUVarint(0)
	return nil
}

func (r *EnvelopeRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	pd = FlexibleDecoderFrom(pd)
	if r.RequestData, err = pd.getBytes(); err != nil {
		return err
	}

	if r.RequestPrincipal, err = pd.getBytes(); err != nil {
		return err
	}

	if r.ClientHostAddress, err = pd.getBytes(); err != nil {
		return err
	}

	if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
		return err
	}
	return nil
}

func (r *EnvelopeRequest) GetKey() int16 {
	return 58
}

func (r *EnvelopeRequest) GetVersion() int16 {
	return r.Version
}

func (r *EnvelopeRequest) GetHeaderVersion() int16 {
	return 2
}

func (r *EnvelopeRequest) IsValidVersion() bool {
	return r.Version == 0
}

func (r *EnvelopeRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
