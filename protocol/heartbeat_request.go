// protocol has been generated from message format json - DO NOT EDIT
package protocol

type HeartbeatRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group id.
	GroupID string
	// GenerationID contains the generation of the group.
	GenerationID int32
	// MemberID contains the member ID.
	MemberID string
	// GroupInstanceID contains the unique identifier of the consumer instance provided by end user.
	GroupInstanceID *string
}

func (r *HeartbeatRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putString(r.GroupID); err != nil {
		return err
	}

	pe.putInt32(r.GenerationID)

	if err := pe.putString(r.MemberID); err != nil {
		return err
	}

	if r.Version >= 3 {
		if err := pe.putNullableString(r.GroupInstanceID); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *HeartbeatRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.GroupID, err = pd.getString(); err != nil {
		return err
	}

	if r.GenerationID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.MemberID, err = pd.getString(); err != nil {
		return err
	}

	if r.Version >= 3 {
		if r.GroupInstanceID, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *HeartbeatRequest) GetKey() int16 {
	return 12
}

func (r *HeartbeatRequest) GetVersion() int16 {
	return r.Version
}

func (r *HeartbeatRequest) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 2
	}
	return 1
}

func (r *HeartbeatRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *HeartbeatRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
