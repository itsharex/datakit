package tc

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/mdlayher/netlink"
)

const (
	tcaHfscUnspec = iota
	tcaHfscRsc
	tcaHfscFsc
	tcaHfscUsc
)

// Hfsc contains attributes of the hfsc class
type Hfsc struct {
	Rsc *ServiceCurve
	Fsc *ServiceCurve
	Usc *ServiceCurve
}

// unmarshalHfsc parses the Hfsc-encoded data and stores the result in the value pointed to by info.
func unmarshalHfsc(data []byte, info *Hfsc) error {
	ad, err := netlink.NewAttributeDecoder(data)
	if err != nil {
		return err
	}
	ad.ByteOrder = nativeEndian
	for ad.Next() {
		switch ad.Type() {
		case tcaHfscRsc:
			curve := &ServiceCurve{}
			if err := unmarshalStruct(ad.Bytes(), curve); err != nil {
				return err
			}
			info.Rsc = curve
		case tcaHfscFsc:
			curve := &ServiceCurve{}
			if err := unmarshalStruct(ad.Bytes(), curve); err != nil {
				return err
			}
			info.Fsc = curve
		case tcaHfscUsc:
			curve := &ServiceCurve{}
			if err := unmarshalStruct(ad.Bytes(), curve); err != nil {
				return err
			}
			info.Usc = curve
		default:
			return fmt.Errorf("unmarshalHfsc()\t%d\n\t%v", ad.Type(), ad.Bytes())
		}
	}
	return nil
}

// marshalHfsc returns the binary encoding of Hfsc
func marshalHfsc(info *Hfsc) ([]byte, error) {
	options := []tcOption{}

	if info == nil {
		return []byte{}, fmt.Errorf("Hfsc: %w", ErrNoArg)
	}

	// TODO: improve logic and check combinations

	if info.Rsc != nil {
		data, err := marshalStruct(info.Rsc)
		if err != nil {
			return []byte{}, err
		}
		options = append(options, tcOption{Interpretation: vtBytes, Type: tcaHfscRsc, Data: data})
	}
	if info.Fsc != nil {
		data, err := marshalStruct(info.Fsc)
		if err != nil {
			return []byte{}, err
		}
		options = append(options, tcOption{Interpretation: vtBytes, Type: tcaHfscFsc, Data: data})
	}
	if info.Usc != nil {
		data, err := marshalStruct(info.Usc)
		if err != nil {
			return []byte{}, err
		}
		options = append(options, tcOption{Interpretation: vtBytes, Type: tcaHfscUsc, Data: data})
	}

	return marshalAttributes(options)
}

// ServiceCurve from include/uapi/linux/pkt_sched.h
type ServiceCurve struct {
	M1 uint32
	D  uint32
	M2 uint32
}

// HfscQOpt contains attributes of the hfsc qdisc
type HfscQOpt struct {
	DefCls uint16
}

// unmarshalHfscQOpt parses the HfscQOpt-encoded data and stores the result in the value pointed to by info.
func unmarshalHfscQOpt(data []byte, info *HfscQOpt) error {

	info.DefCls = nativeEndian.Uint16(data)

	return nil
}

// marshalHfscQOpt returns the binary encoding of HfscQOpt
func marshalHfscQOpt(info *HfscQOpt) ([]byte, error) {
	if info == nil {
		return []byte{}, fmt.Errorf("HfscQOpt: %w", ErrNoArg)
	}

	data := bytes.NewBuffer(make([]byte, 0, 2))
	if err := binary.Write(data, nativeEndian, info.DefCls); err != nil {
		return []byte{}, err
	}
	return data.Bytes(), nil
}
