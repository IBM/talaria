// protocol has been generated from message format json - DO NOT EDIT
package protocol

// SyncGroupRequestAssignment contains each assignment.
type SyncGroupRequestAssignment struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// MemberID contains the ID of the member to assign.
	MemberID string
	// Assignment contains the member assignment.
	Assignment []byte
}

func (a *SyncGroupRequestAssignment) encode(pe packetEncoder, version int16) (err error) {
	a.Version = version
	if err := pe.putString(a.MemberID); err != nil {
		return err
	}

	if err := pe.putBytes(a.Assignment); err != nil {
		return err
	}

	if a.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (a *SyncGroupRequestAssignment) decode(pd packetDecoder, version int16) (err error) {
	a.Version = version
	if a.MemberID, err = pd.getString(); err != nil {
		return err
	}

	if a.Assignment, err = pd.getBytes(); err != nil {
		return err
	}

	if a.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type SyncGroupRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the unique group identifier.
	GroupID string
	// GenerationID contains the generation of the group.
	GenerationID int32
	// MemberID contains the member ID assigned by the group.
	MemberID string
	// GroupInstanceID contains the unique identifier of the consumer instance provided by end user.
	GroupInstanceID *string
	// ProtocolType contains the group protocol type.
	ProtocolType *string
	// ProtocolName contains the group protocol name.
	ProtocolName *string
	// Assignments contains each assignment.
	Assignments []SyncGroupRequestAssignment
}

func (r *SyncGroupRequest) encode(pe packetEncoder) (err error) {
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

	if r.Version >= 5 {
		if err := pe.putNullableString(r.ProtocolType); err != nil {
			return err
		}
	}

	if r.Version >= 5 {
		if err := pe.putNullableString(r.ProtocolName); err != nil {
			return err
		}
	}

	if err := pe.putArrayLength(len(r.Assignments)); err != nil {
		return err
	}
	for _, block := range r.Assignments {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *SyncGroupRequest) decode(pd packetDecoder, version int16) (err error) {
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

	if r.Version >= 5 {
		if r.ProtocolType, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if r.Version >= 5 {
		if r.ProtocolName, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	var numAssignments int
	if numAssignments, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numAssignments > 0 {
		r.Assignments = make([]SyncGroupRequestAssignment, numAssignments)
		for i := 0; i < numAssignments; i++ {
			var block SyncGroupRequestAssignment
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Assignments[i] = block
		}
	}

	if r.Version >= 4 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *SyncGroupRequest) GetKey() int16 {
	return 14
}

func (r *SyncGroupRequest) GetVersion() int16 {
	return r.Version
}

func (r *SyncGroupRequest) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 2
	}
	return 1
}

func (r *SyncGroupRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 5
}

func (r *SyncGroupRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
