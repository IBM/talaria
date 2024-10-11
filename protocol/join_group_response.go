// protocol has been generated from message format json - DO NOT EDIT
package protocol

import (
	"fmt"
	"time"
)

// JoinGroupResponseMember contains a
type JoinGroupResponseMember struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// MemberID contains the group member ID.
	MemberID string
	// GroupInstanceID contains the unique identifier of the consumer instance provided by end user.
	GroupInstanceID *string
	// Metadata contains the group member metadata.
	Metadata []byte
}

func (m *JoinGroupResponseMember) encode(pe packetEncoder, version int16) (err error) {
	m.Version = version
	if err := pe.putString(m.MemberID); err != nil {
		return err
	}

	if m.Version >= 5 {
		if err := pe.putNullableString(m.GroupInstanceID); err != nil {
			return err
		}
	}

	if err := pe.putBytes(m.Metadata); err != nil {
		return err
	}

	if m.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (m *JoinGroupResponseMember) decode(pd packetDecoder, version int16) (err error) {
	m.Version = version
	if m.MemberID, err = pd.getString(); err != nil {
		return err
	}

	if m.Version >= 5 {
		if m.GroupInstanceID, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if m.Metadata, err = pd.getBytes(); err != nil {
		return err
	}

	if m.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type JoinGroupResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// GenerationID contains the generation ID of the group.
	GenerationID int32
	// ProtocolType contains the group protocol name.
	ProtocolType *string
	// ProtocolName contains the group protocol selected by the coordinator.
	ProtocolName *string
	// Leader contains the leader of the group.
	Leader string
	// SkipAssignment contains a True if the leader must skip running the assignment.
	SkipAssignment bool
	// MemberID contains the member ID assigned by the group coordinator.
	MemberID string
	// Members contains a
	Members []JoinGroupResponseMember
}

func (r *JoinGroupResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 6 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 2 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	pe.putInt16(r.ErrorCode)

	pe.putInt32(r.GenerationID)

	if r.Version >= 7 {
		if err := pe.putNullableString(r.ProtocolType); err != nil {
			return err
		}
	}

	if r.Version >= 7 {
		if err := pe.putNullableString(r.ProtocolName); err != nil {
			return err
		}
	} else {
		if r.ProtocolName == nil {
			return fmt.Errorf("String field, ProtocolName, must not be nil in version %d", r.Version)
		}
		if err := pe.putString(*r.ProtocolName); err != nil {
			return err
		}
	}

	if err := pe.putString(r.Leader); err != nil {
		return err
	}

	if r.Version >= 9 {
		pe.putBool(r.SkipAssignment)
	}

	if err := pe.putString(r.MemberID); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Members)); err != nil {
		return err
	}
	for _, block := range r.Members {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *JoinGroupResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 6 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 2 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.GenerationID, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 7 {
		if r.ProtocolType, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if r.Version >= 7 {
		if r.ProtocolName, err = pd.getNullableString(); err != nil {
			return err
		}
	} else {
		if *r.ProtocolName, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Leader, err = pd.getString(); err != nil {
		return err
	}

	if r.Version >= 9 {
		if r.SkipAssignment, err = pd.getBool(); err != nil {
			return err
		}
	}

	if r.MemberID, err = pd.getString(); err != nil {
		return err
	}

	var numMembers int
	if numMembers, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Members = make([]JoinGroupResponseMember, numMembers)
	for i := 0; i < numMembers; i++ {
		var block JoinGroupResponseMember
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Members[i] = block
	}

	if r.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *JoinGroupResponse) GetKey() int16 {
	return 11
}

func (r *JoinGroupResponse) GetVersion() int16 {
	return r.Version
}

func (r *JoinGroupResponse) GetHeaderVersion() int16 {
	if r.Version >= 6 {
		return 1
	}
	return 0
}

func (r *JoinGroupResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 9
}

func (r *JoinGroupResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *JoinGroupResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
