// protocol has been generated from message format json - DO NOT EDIT
package protocol

type SaslHandshakeRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Mechanism contains the SASL mechanism chosen by the client.
	Mechanism string
}

func (r *SaslHandshakeRequest) encode(pe packetEncoder) (err error) {
	if err := pe.putString(r.Mechanism); err != nil {
		return err
	}

	return nil
}

func (r *SaslHandshakeRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Mechanism, err = pd.getString(); err != nil {
		return err
	}

	return nil
}

func (r *SaslHandshakeRequest) GetKey() int16 {
	return 17
}

func (r *SaslHandshakeRequest) GetVersion() int16 {
	return r.Version
}

func (r *SaslHandshakeRequest) GetHeaderVersion() int16 {
	return 1
}

func (r *SaslHandshakeRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 1
}

func (r *SaslHandshakeRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
