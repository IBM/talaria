// protocol has been generated from message format json - DO NOT EDIT
package protocol

import "time"

// Coordinator contains each coordinator result in the response
type Coordinator struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// Key contains the coordinator key.
	Key string
	// NodeID contains the node id.
	NodeID int32
	// Host contains the host name.
	Host string
	// Port contains the port.
	Port int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
}

func (c *Coordinator) encode(pe packetEncoder, version int16) (err error) {
	c.Version = version
	if c.Version >= 4 {
		if err := pe.putString(c.Key); err != nil {
			return err
		}
	}

	if c.Version >= 4 {
		pe.putInt32(c.NodeID)
	}

	if c.Version >= 4 {
		if err := pe.putString(c.Host); err != nil {
			return err
		}
	}

	if c.Version >= 4 {
		pe.putInt32(c.Port)
	}

	if c.Version >= 4 {
		pe.putInt16(c.ErrorCode)
	}

	if c.Version >= 4 {
		if err := pe.putNullableString(c.ErrorMessage); err != nil {
			return err
		}
	}

	if c.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (c *Coordinator) decode(pd packetDecoder, version int16) (err error) {
	c.Version = version
	if c.Version >= 4 {
		if c.Key, err = pd.getString(); err != nil {
			return err
		}
	}

	if c.Version >= 4 {
		if c.NodeID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if c.Version >= 4 {
		if c.Host, err = pd.getString(); err != nil {
			return err
		}
	}

	if c.Version >= 4 {
		if c.Port, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if c.Version >= 4 {
		if c.ErrorCode, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if c.Version >= 4 {
		if c.ErrorMessage, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if c.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

type FindCoordinatorResponse struct {
	// Version defines the protocol version to use for encode and decode
	Version int16
	// ThrottleTimeMs contains the duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	ThrottleTimeMs int32
	// ErrorCode contains the error code, or 0 if there was no error.
	ErrorCode int16
	// ErrorMessage contains the error message, or null if there was no error.
	ErrorMessage *string
	// NodeID contains the node id.
	NodeID int32
	// Host contains the host name.
	Host string
	// Port contains the port.
	Port int32
	// Coordinators contains each coordinator result in the response
	Coordinators []Coordinator
}

func (r *FindCoordinatorResponse) encode(pe packetEncoder) (err error) {
	if r.Version >= 3 {
		pe = FlexibleEncoderFrom(pe)
	}
	if r.Version >= 1 {
		pe.putInt32(r.ThrottleTimeMs)
	}

	if r.Version >= 0 && r.Version <= 3 {
		pe.putInt16(r.ErrorCode)
	}

	if r.Version >= 1 && r.Version <= 3 {
		if err := pe.putNullableString(r.ErrorMessage); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		pe.putInt32(r.NodeID)
	}

	if r.Version >= 0 && r.Version <= 3 {
		if err := pe.putString(r.Host); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		pe.putInt32(r.Port)
	}

	if r.Version >= 4 {
		if err := pe.putArrayLength(len(r.Coordinators)); err != nil {
			return err
		}
		for _, block := range r.Coordinators {
			if err := block.encode(pe, r.Version); err != nil {
				return err
			}
		}
	}

	if r.Version >= 3 {
		pe.putUVarint(0)
	}
	return nil
}

func (r *FindCoordinatorResponse) decode(pd packetDecoder, version int16) (err error) {
	r.Version = version
	if r.Version >= 3 {
		pd = FlexibleDecoderFrom(pd)
	}
	if r.Version >= 1 {
		if r.ThrottleTimeMs, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		if r.ErrorCode, err = pd.getInt16(); err != nil {
			return err
		}
	}

	if r.Version >= 1 && r.Version <= 3 {
		if r.ErrorMessage, err = pd.getNullableString(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		if r.NodeID, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		if r.Host, err = pd.getString(); err != nil {
			return err
		}
	}

	if r.Version >= 0 && r.Version <= 3 {
		if r.Port, err = pd.getInt32(); err != nil {
			return err
		}
	}

	if r.Version >= 4 {
		var numCoordinators int
		if numCoordinators, err = pd.getArrayLength(); err != nil {
			return err
		}
		r.Coordinators = make([]Coordinator, numCoordinators)
		for i := 0; i < numCoordinators; i++ {
			var block Coordinator
			if err := block.decode(pd, r.Version); err != nil {
				return err
			}
			r.Coordinators[i] = block
		}
	}

	if r.Version >= 3 {
		if _, err = pd.getEmptyTaggedFieldArray(); err != nil {
			return err
		}
	}
	return nil
}

func (r *FindCoordinatorResponse) GetKey() int16 {
	return 10
}

func (r *FindCoordinatorResponse) GetVersion() int16 {
	return r.Version
}

func (r *FindCoordinatorResponse) GetHeaderVersion() int16 {
	if r.Version >= 3 {
		return 1
	}
	return 0
}

func (r *FindCoordinatorResponse) IsValidVersion() bool {
	return r.Version >= 0 && r.Version <= 4
}

func (r *FindCoordinatorResponse) GetRequiredVersion() int16 {
	// TODO - it isn't possible to determine this from the message format json files
	return 0
}

func (r *FindCoordinatorResponse) throttleTime() time.Duration {
	return time.Duration(r.ThrottleTimeMs) * time.Millisecond
}
