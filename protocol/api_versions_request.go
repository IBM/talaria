// protocol has been generated from message format json - DO NOT EDIT
package protocol

type ApiVersionsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ClientSoftwareName contains the name of the client.
	ClientSoftwareName string
	// ClientSoftwareVersion contains the version of the client.
	ClientSoftwareVersion string
}

func (r *ApiVersionsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 3 {
		if err := pe.putString(r.ClientSoftwareName); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if err := pe.putString(r.ClientSoftwareVersion); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ApiVersionsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 3 {
		if r.ClientSoftwareName, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if r.ClientSoftwareVersion, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ApiVersionsRequest) GetKey() int16 {
	return 18
}

func (r *ApiVersionsRequest) GetVersion() int16 {
	return r.Version
}

func (r *ApiVersionsRequest) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 2
	}
	return 1
}

func (r *ApiVersionsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 3
}

func (r *ApiVersionsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
