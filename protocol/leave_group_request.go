// protocol has been generated from message format json - DO NOT EDIT
package protocol

// MemberIdentity contains a List of leaving member identities.
type MemberIdentity struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// MemberID contains the member ID to remove from the group.
	MemberID string
	// GroupInstanceID contains the group instance ID to remove from the group.
	GroupInstanceID *string
	// Reason contains the reason why the member left the group.
	Reason *string
}

func (m *MemberIdentity) encode(pe packetEncoder, version int16) (err error) {
	m.Version = version
	if m.Version >= 3 {
		if err := pe.putString(m.MemberID); err != nil {
			return err
		}
	}

	if m.Version >= 3 {
		if err := pe.putNullableString(m.GroupInstanceID); err != nil {
			return err
		}
	}

	if m.Version >= 5 {
		if err := pe.putNullableString(m.Reason); err != nil {
			return err
		}
	}

	if m.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (m *MemberIdentity) decode(pd packetDecoder, version int16) (err error) {
	m.Version = version
	if m.Version >= 3 {
		if m.MemberID, err = pd.getString(); err != nil {
			return err
		}
	}

	if m.Version >= 3 {
		if m.GroupInstanceID, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if m.Version >= 5 {
		if m.Reason, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if m.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type LeaveGroupRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the ID of the group to leave.
	GroupID string
	// MemberID contains the member ID to remove from the group.
	MemberID string
	// Members contains a List of leaving member identities.
	Members []MemberIdentity
}

func (r *LeaveGroupRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putString(r.GroupID); err != nil {
		return err
	}

	if r.Version >= 0 && r.Version <= 2 {
		if err := pe.putString(r.MemberID); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		if err := pe.putArrayLength(len(r.Members)); err != nil {
			return err
		}
		for _, block := range r.Members {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *LeaveGroupRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.GroupID, err = pd.getString(); err != nil {
		return err
	}

	if r.Version >= 0 && r.Version <= 2 {
		if r.MemberID, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		var numMembers int
		if numMembers, err = pd.getArrayLength(); err != nil {
			return err
		}
		r.Members = make([]MemberIdentity, numMembers)
		for i := 0; i < numMembers; i++ {
			var block MemberIdentity
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Members[i] = block
		}
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *LeaveGroupRequest) GetKey() int16 {
	return 13
}

func (r *LeaveGroupRequest) GetVersion() int16 {
	return r.Version
}

func (r *LeaveGroupRequest) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 2
	}
	return 1
}

func (r *LeaveGroupRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 5
}

func (r *LeaveGroupRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
