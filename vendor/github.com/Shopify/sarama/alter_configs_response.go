package sarama

import "time"

// AlterConfigsResponse is a response type for alter config
type AlterConfigsResponse struct {
	ThrottleTime time.Duration
	Resources    []*AlterConfigsResourceResponse
}

// AlterConfigsResourceResponse is a response type for alter config resource
type AlterConfigsResourceResponse struct {
	ErrorCode int16
	ErrorMsg  string
	Type      ConfigResourceType
	Name      string
}

func (a *AlterConfigsResponse) encode(pe packetEncoder) error {
	pe.putInt32(int32(a.ThrottleTime / time.Millisecond))

	if err := pe.putArrayLength(len(a.Resources)); err != nil {
		return err
	}

	for _, v := range a.Resources {
		if err := v.encode(pe); err != nil {
			return err
		}
	}

	return nil
}

func (a *AlterConfigsResponse) decode(pd packetDecoder, version int16) error {
	throttleTime, err := pd.getInt32()
	if err != nil {
		return err
	}
	a.ThrottleTime = time.Duration(throttleTime) * time.Millisecond

	responseCount, err := pd.getArrayLength()
	if err != nil {
		return err
	}

	a.Resources = make([]*AlterConfigsResourceResponse, responseCount)

	for i := range a.Resources {
		a.Resources[i] = new(AlterConfigsResourceResponse)

		if err := a.Resources[i].decode(pd, version); err != nil {
			return err
		}
	}

	return nil
}

func (a *AlterConfigsResourceResponse) encode(pe packetEncoder) error {
	pe.putInt16(a.ErrorCode)
	err := pe.putString(a.ErrorMsg)
	if err != nil {
		return err
	}
	pe.putInt8(int8(a.Type))
	err = pe.putString(a.Name)
	if err != nil {
		return err
	}
	return nil
}

func (a *AlterConfigsResourceResponse) decode(pd packetDecoder, version int16) error {
	errCode, err := pd.getInt16()
	if err != nil {
		return err
	}
	a.ErrorCode = errCode

	e, err := pd.getString()
	if err != nil {
		return err
	}
	a.ErrorMsg = e

	t, err := pd.getInt8()
	if err != nil {
		return err
	}
	a.Type = ConfigResourceType(t)

	name, err := pd.getString()
	if err != nil {
		return err
	}
	a.Name = name

	return nil
}

func (a *AlterConfigsResponse) key() int16 {
	return 32
}

func (a *AlterConfigsResponse) version() int16 {
	return 0
}

func (a *AlterConfigsResponse) headerVersion() int16 {
	return 0
}

func (a *AlterConfigsResponse) requiredVersion() KafkaVersion {
	return V0_11_0_0
}
