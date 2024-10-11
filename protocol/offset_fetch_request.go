// protocol has been generated from message format json - DO NOT EDIT
package protocol

// OffsetFetchRequestTopic contains each topic we would like to fetch offsets for, or null to fetch offsets for all topics.
type OffsetFetchRequestTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// PartitionIndexes contains the partition indexes we would like to fetch offsets for.
	PartitionIndexes []int32
}

func (t *OffsetFetchRequestTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 0 && t.Version <= 7 {
		if err := pe.putString(t.Name); err != nil {
			return err
		}
	}

	if t.Version >= 0 && t.Version <= 7 {
		if err := pe.putInt32Array(t.PartitionIndexes); err != nil {
			return err
		}
	}

	if t.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *OffsetFetchRequestTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 0 && t.Version <= 7 {
		if t.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 0 && t.Version <= 7 {
		if t.PartitionIndexes, err = pd.getInt32Array(); err != nil {
			return err
		}
	}

	if t.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// OffsetFetchRequestTopics contains each topic we would like to fetch offsets for, or null to fetch offsets for all topics.
type OffsetFetchRequestTopics struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// PartitionIndexes contains the partition indexes we would like to fetch offsets for.
	PartitionIndexes []int32
}

func (t *OffsetFetchRequestTopics) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 8 {
		if err := pe.putString(t.Name); err != nil {
			return err
		}
	}

	if t.Version >= 8 {
		if err := pe.putInt32Array(t.PartitionIndexes); err != nil {
			return err
		}
	}

	if t.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *OffsetFetchRequestTopics) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Version >= 8 {
		if t.Name, err = pd.getString(); err != nil {
			return err
		}
	}

	if t.Version >= 8 {
		if t.PartitionIndexes, err = pd.getInt32Array(); err != nil {
			return err
		}
	}

	if t.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// OffsetFetchRequestGroup contains each group we would like to fetch offsets for
type OffsetFetchRequestGroup struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// groupID contains the group ID.
	groupID string
	// Topics contains each topic we would like to fetch offsets for, or null to fetch offsets for all topics.
	Topics []OffsetFetchRequestTopics
}

func (g *OffsetFetchRequestGroup) encode(pe packetEncoder, version int16) (err error) {
	g.Version = version
	if g.Version >= 8 {
		if err := pe.putString(g.groupID); err != nil {
			return err
		}
	}

	if g.Version >= 8 {
		if err := pe.putArrayLength(len(g.Topics)); err != nil {
			return err
		}
		for _, block := range g.Topics {
			if err := block.encode(pe, g.Version); err != nil {
				return err
			}
		}
	}

	if g.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (g *OffsetFetchRequestGroup) decode(pd packetDecoder, version int16) (err error) {
	g.Version = version
	if g.Version >= 8 {
		if g.groupID, err = pd.getString(); err != nil {
			return err
		}
	}

	if g.Version >= 8 {
		var numTopics int
		if numTopics, err = pd.getArrayLength(); err != nil {
			return err
		}
		g.Topics = make([]OffsetFetchRequestTopics, numTopics)
		for i := 0; i < numTopics; i++ {
			var block OffsetFetchRequestTopics
			if err := block.decode(pd, g.Version); err != nil {
				return err
			}
			g.Topics[i] = block
		}
	}

	if g.Version >= 6 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type OffsetFetchRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// GroupID contains the group to fetch offsets for.
	GroupID string
	// Topics contains each topic we would like to fetch offsets for, or null to fetch offsets for all topics.
	Topics []OffsetFetchRequestTopic
	// Groups contains each group we would like to fetch offsets for
	Groups []OffsetFetchRequestGroup
	// RequireStable contains a Whether broker should hold on returning unstable offsets but set a retriable error code for the partitions.
	RequireStable bool
}

func (r *OffsetFetchRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 6 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 0 && r.Version <= 7 {
		if err := pe.putString(r.GroupID); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 7 {
		if err := pe.putArrayLength(len(r.Topics)); err != nil {
			return err
		}
		for _, block := range r.Topics {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 8 {
		if err := pe.putArrayLength(len(r.Groups)); err != nil {
			return err
		}
		for _, block := range r.Groups {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 7 {
		pe.putBool(r.RequireStable)
	}

	if r.Version >= 6 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *OffsetFetchRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 6 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 0 && r.Version <= 7 {
		if r.GroupID, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 7 {
		var numTopics int
		if numTopics, err = pd.getArrayLength(); err != nil {
			return err
		}
		r.Topics = make([]OffsetFetchRequestTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block OffsetFetchRequestTopic
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Topics[i] = block
		}
	}

	if r.Version >= 8 {
		var numGroups int
		if numGroups, err = pd.getArrayLength(); err != nil {
			return err
		}
		r.Groups = make([]OffsetFetchRequestGroup, numGroups)
		for i := 0; i < numGroups; i++ {
			var block OffsetFetchRequestGroup
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Groups[i] = block
		}
	}

	if r.Version >= 7 {
		if r.RequireStable, err = pd.getBool(); err != nil {
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

func (r *OffsetFetchRequest) GetKey() int16 {
	return 9
}

func (r *OffsetFetchRequest) GetVersion() int16 {
	return r.Version
}

func (r *OffsetFetchRequest) GetHeaderVersion() int16 {
	if r.Version >= 6 {
		return 2
	}
	return 1
}

func (r *OffsetFetchRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 8
}

func (r *OffsetFetchRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
