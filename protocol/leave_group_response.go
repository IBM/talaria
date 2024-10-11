// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// MemberResponse contains a List of leaving member responses.
type MemberResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// MemberID contains the member ID to remove from the group.
	MemberID string
	// GroupInstanceID contains the group instance ID to remove from the group.
	GroupInstanceID *string
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
}

func (m *MemberResponse) encode(pe packetEncoder, version int16) (err error) {
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

	if m.Version >= 3 {
		pe.putInt16(m.ErrorCode)
	}

	if m.Version >= 4 {
		pe.putUVarint(0)
	}
	return nil
}

func (m *MemberResponse) decode(pd packetDecoder, version int16) (err error) {
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

	if m.Version >= 3 {
		if m.ErrorCode, err = pd.getInt16(); err != nil {
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

type LeaveGroupResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// Members contains a List of leaving member responses.
	Members []MemberResponse
}

func (r *LeaveGroupResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 4 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 1 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	pe.putInt16(r.ErrorCode)

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

func (r *LeaveGroupResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 4 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 1 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Version >= 3 {
		var numMembers int
		if numMembers, err = pd.getArrayLength(); err != nil {
			return err
		}
		r.Members = make([]MemberResponse, numMembers)
		for i := 0; i < numMembers; i++ {
			var block MemberResponse
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

func (r *LeaveGroupResponse) GetKey() int16 {
	return 13
}

func (r *LeaveGroupResponse) GetVersion() int16 {
	return r.Version
}

func (r *LeaveGroupResponse) GetHeaderVersion() int16 {
	if r.Version >= 4 {
		return 1
	}
	return 0
}

func (r *LeaveGroupResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 5
}

func (r *LeaveGroupResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *LeaveGroupResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
