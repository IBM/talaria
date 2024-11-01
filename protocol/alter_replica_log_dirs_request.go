// protocol has been generated from message format json - DO NOT EDIT
package protocol

// AlterReplicaLogDirTopic contains the topics to add to the directory.
type AlterReplicaLogDirTopic struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Name contains the topic name.
	Name string
	// Partitions contains the partition indexes.
	Partitions []int32
}

func (t *AlterReplicaLogDirTopic) encode(pe packetEncoder, version int16) (err error) {
	t.Version = version
	if err := pe.putString(t.Name); err != nil {
		return err
	}

	if err := pe.putInt32Array(t.Partitions); err != nil {
		return err
	}

	if t.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (t *AlterReplicaLogDirTopic) decode(pd packetDecoder, version int16) (err error) {
	t.Version = version
	if t.Name, err = pd.getString(); err != nil {
		return err
	}

	if t.Partitions, err = pd.getInt32Array(); err != nil {
		return err
	}

	if t.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

// AlterReplicaLogDir contains the alterations to make for each directory.
type AlterReplicaLogDir struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Path contains the absolute directory path.
	Path string
	// Topics contains the topics to add to the directory.
	Topics []AlterReplicaLogDirTopic
}

func (d *AlterReplicaLogDir) encode(pe packetEncoder, version int16) (err error) {
	d.Version = version
	if err := pe.putString(d.Path); err != nil {
		return err
	}

	if err := pe.putArrayLength(len(d.Topics)); err != nil {
		return err
	}
	for _, block := range d.Topics {
		if err := block.encode(pe, d.Version); err != nil {
			return err
		}
	}

	if d.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (d *AlterReplicaLogDir) decode(pd packetDecoder, version int16) (err error) {
	d.Version = version
	if d.Path, err = pd.getString(); err != nil {
		return err
	}

	var numTopics int
	if numTopics, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numTopics > 0 {
		d.Topics = make([]AlterReplicaLogDirTopic, numTopics)
		for i := 0; i < numTopics; i++ {
			var block AlterReplicaLogDirTopic
			if err := block.decode(pd, d.Version); err != nil {
				return err
			}
			d.Topics[i] = block
		}
	}

	if d.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type AlterReplicaLogDirsRequest struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Dirs contains the alterations to make for each directory.
	Dirs []AlterReplicaLogDir
}

func (r *AlterReplicaLogDirsRequest) encode(pe packetEncoder) (err error) {
	if r.Version >= 2 {
		pe = FlexibleEncoderFrom(pe)
	}
	if err := pe.putArrayLength(len(r.Dirs)); err != nil {
		return err
	}
	for _, block := range r.Dirs {
		if err := block.encode(pe, r.Version); err != nil {
			return err
		}
	}

	if r.Version >= 2 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *AlterReplicaLogDirsRequest) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 2 {
		pd = FlexibleDecoderFrom(pd)
	}
	var numDirs int
	if numDirs, err = pd.getArrayLength(); err != nil {
		return err
	}
	if numDirs > 0 {
		r.Dirs = make([]AlterReplicaLogDir, numDirs)
		for i := 0; i < numDirs; i++ {
			var block AlterReplicaLogDir
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Dirs[i] = block
		}
	}

	if r.Version >= 2 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *AlterReplicaLogDirsRequest) GetKey() int16 {
	return 34
}

func (r *AlterReplicaLogDirsRequest) GetVersion() int16 {
	return r.Version
}

func (r *AlterReplicaLogDirsRequest) GetHeaderVersion() int16 {
	if r.Version >= 2 {
		return 2
	}
	return 1
}

func (r *AlterReplicaLogDirsRequest) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 2
}

func (r *AlterReplicaLogDirsRequest) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}
