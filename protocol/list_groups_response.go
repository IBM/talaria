// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// ListedGroup contains each group in the response.
type ListedGroup struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group ID.
	GroupID string
	// ProtocolType contains the group protocol type.
	ProtocolType string
	// GroupState contains the group state name.
	GroupState string
}

func (g *ListedGroup) encode(pe packetEncoder, version int16) (err error) {
	g.Version = version
	if err := pe.putString(g.GroupID); err != nil {
		return err
	}

	if err := pe.putString(g.ProtocolType); err != nil {
		return err
	}

	if g.Version >= 4 {
		if err := pe.putString(g.GroupState); err != nil {
			return err
		}
	}

	if g.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (g *ListedGroup) decode(pd packetDecoder, version int16) (err error) {
	g.Version = version
	if g.GroupID, err = pd.getString(); err != nil {
		return err
	}

	if g.ProtocolType, err = pd.getString(); err != nil {
		return err
	}

	if g.Version >= 4 {
		if g.GroupState, err = pd.getString(); err != nil {
			return err
		}
	}

	if g.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type ListGroupsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// Groups contains each group in the response.
	Groups []ListedGroup
}

func (r *ListGroupsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 1 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	pe.putInt16(r.ErrorCode)

	if err := pe.putArrayLength(len(r.Groups)); err != nil {
		return err
	}
	for _, block := range r.Groups {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *ListGroupsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
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

	var numGroups int
	if numGroups, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Groups = make([]ListedGroup, numGroups)
	for i := 0; i < numGroups; i++ {
		var block ListedGroup
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Groups[i] = block
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ListGroupsResponse) GetKey() int16 {
	return 16
}

func (r *ListGroupsResponse) GetVersion() int16 {
	return r.Version
}

func (r *ListGroupsResponse) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 1
	}
	return 0
}

func (r *ListGroupsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *ListGroupsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *ListGroupsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
