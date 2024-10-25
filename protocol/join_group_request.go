// protocol has been generated from message format json - DO NOT EDIT
package protocol

// JoinGroupRequestProtocol contains the list of protocols that the member supports.
type JoinGroupRequestProtocol struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the protocol name.
	Name string
	// Metadata contains the protocol metadata.
	Metadata []byte
}

func (p *JoinGroupRequestProtocol) encode(pe packetEncoder, version int16) (err error) {
	p.Version = version
	if err := pe.putString(p.Name); err != nil {
		return err
	}

	if err := pe.putBytes(p.Metadata); err != nil {
		return err
	}

	if p.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (p *JoinGroupRequestProtocol) decode(pd packetDecoder, version int16) (err error) {
	p.Version = version
	if p.Name, err = pd.getString(); err != nil {
		return err
	}

	if p.Metadata, err = pd.getBytes(); err != nil {
		return err
	}

	if p.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type JoinGroupRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group identifier.
	GroupID string
	// SessionTimeoutMs contains the coordinator considers the consumer dead if it receives no heartbeat after this timeout in milliseconds.
	SessionTimeoutMs int32
	// RebalanceTimeoutMs contains the maximum time in milliseconds that the coordinator will wait for each member to rejoin when rebalancing the group.
	RebalanceTimeoutMs int32
	// MemberID contains the member id assigned by the group coordinator.
	MemberID string
	// GroupInstanceID contains the unique identifier of the consumer instance provided by end user.
	GroupInstanceID *string
	// ProtocolType contains the unique name the for class of protocols implemented by the group we want to join.
	ProtocolType string
	// Protocols contains the list of protocols that the member supports.
	Protocols []JoinGroupRequestProtocol
	// Reason contains the reason why the member (re-)joins the group.
	Reason *string
}

func (r *JoinGroupRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 6 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putString(r.GroupID); err != nil {
		return err
	}

	pe.putInt32(r.SessionTimeoutMs)

	if r.Version >= 1 {
		pe.putInt32(r.RebalanceTimeoutMs)
	}

	if err := pe.putString(r.MemberID); err != nil {
		return err
	}

	if r.Version >= 5 {
		if err := pe.putNullableString(r.GroupInstanceID); err != nil {
			return err
		}
	}

	if err := pe.putString(r.ProtocolType); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(r.Protocols)); err != nil {
		return err
	}
	for _, block := range r.Protocols {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 8 {
		if err := pe.putNullableString(r.Reason); err != nil {
			return err
		}
	}

	if r.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *JoinGroupRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 6 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.GroupID, err = pd.getString(); err != nil {
		return err
	}

	if r.SessionTimeoutMs, err = pd.getInt32(); err != nil {
		return err
	}

	if r.Version >= 1 {
		if r.RebalanceTimeoutMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.MemberID, err = pd.getString(); err != nil {
		return err
	}

	if r.Version >= 5 {
		if r.GroupInstanceID, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if r.ProtocolType, err = pd.getString(); err != nil {
		return err
	}

	var numProtocols int
	if numProtocols, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numProtocols > 0 {
		r.Protocols = make([]JoinGroupRequestProtocol, numProtocols)
		for i := 0; i < numProtocols; i++ {
			var block JoinGroupRequestProtocol
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Protocols[i] = block
		}
	}

	if r.Version >= 8 {
		if r.Reason, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if r.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *JoinGroupRequest) GetKey() int16 {
	return 11
}

func (r *JoinGroupRequest) GetVersion() int16 {
	return r.Version
}

func (r *JoinGroupRequest) GetHeaderVersion() int16 {
	if r.Version >= 6 {
		return 2
	}
	return 1
}

func (r *JoinGroupRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 9
}

func (r *JoinGroupRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
