// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// DeletableGroupResult contains the deletion results
type DeletableGroupResult struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group id
	GroupID string
	// ErrorCode contains the deletion error, or 0 if the deletion succeeded.
	ErrorCode int16
}

func (r *DeletableGroupResult) encode(pe packetEncoder, version int16) (err error) {
	r.Version = version
	if err := pe.putString(r.GroupID); err != nil {
		return err
	}

	pe.putInt16(r.ErrorCode)

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DeletableGroupResult) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.GroupID, err = pd.getString(); err != nil {
		return err
	}

	if r.ErrorCode, err = pd.getInt16(); err != nil {
		return err
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteGroupsResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// Results contains the deletion results
	Results []DeletableGroupResult
}

func (r *DeleteGroupsResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	pe.putInt32(r.ThrottleTimeMs)

	if err := pe.putArrayLength(len(r.Results)); err != nil {
		return err
	}
	for _, block := range r.Results {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *DeleteGroupsResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
		return err
	}

	var numResults int
	if numResults, err = pd.getArrayLength(); err != nil {
		return err
	}
	r.Results = make([]DeletableGroupResult, numResults)
	for i := 0; i < numResults; i++ {
		var block DeletableGroupResult
		if err := block.decode(pd, r.Version); err != nil {
			return err
		}
		r.Results[i] = block
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *DeleteGroupsResponse) GetKey() int16 {
	return 42
}

func (r *DeleteGroupsResponse) GetVersion() int16 {
	return r.Version
}

func (r *DeleteGroupsResponse) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 1
	}
	return 0
}

func (r *DeleteGroupsResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *DeleteGroupsResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *DeleteGroupsResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
