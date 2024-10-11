// protocol has been generated from message format json - DO NOT EDIT
package protocol

type FindCoordinatorRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Key contains the coordinator key.
	Key string
	// KeyType contains the coordinator key type. (Group, transaction, etc.)
	KeyType int8
	// CoordinatorKeys contains the coordinator keys.
	CoordinatorKeys []string
}

func (r *FindCoordinatorRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 0 && r.Version <= 3 {
		if err := pe.putString(r.Key); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		pe.putInt8(r.KeyType)
	}

	if r.Version >= 4 {
		if err := pe.putStringArray(r.CoordinatorKeys); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *FindCoordinatorRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 0 && r.Version <= 3 {
		if r.Key, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 1 {
		if r.KeyType, err = pd.getInt8(); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		if r.CoordinatorKeys, err = pd.getStringArray(); err != nil {
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

func (r *FindCoordinatorRequest) GetKey() int16 {
	return 10
}

func (r *FindCoordinatorRequest) GetVersion() int16 {
	return r.Version
}

func (r *FindCoordinatorRequest) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 2
	}
	return 1
}

func (r *FindCoordinatorRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *FindCoordinatorRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
