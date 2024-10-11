// protocol has been generated from message format json - DO NOT EDIT
package protocol

type SaslHandshakeResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// Mechanisms contains the mechanisms enabled in the server.
	Mechanisms []string
}

func (r *SaslHandshakeResponse) encode(pe packetEncoder) (err error) {
	pe.putInt16(r.ErrorCode)

	if err := pe.putStringArray(r.Mechanisms); err != nil {
		return err
	}

	return nil
}

func (r *SaslHandshakeResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Mechanisms, err = pd.getStringArray(); err != nil {
		return err
	}

	return nil
}

func (r *SaslHandshakeResponse) GetKey() int16 {
	return 17
}

func (r *SaslHandshakeResponse) GetVersion() int16 {
	return r.Version
}

func (r *SaslHandshakeResponse) GetHeaderVersion() int16 {
	return 0
}

func (r *SaslHandshakeResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *SaslHandshakeResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
