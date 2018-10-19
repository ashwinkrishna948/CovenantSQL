package types

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	hsp "github.com/CovenantSQL/HashStablePack/marshalhash"
)

// MarshalHash marshals for hash
func (z *TokenReceive) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 4
	o = append(o, 0x84, 0x84)
	if z.Signee == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.Signee.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x84)
	if z.Signature == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.Signature.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x84)
	if oTemp, err := z.TokenReceiveHeader.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x84)
	if oTemp, err := z.HeaderHash.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TokenReceive) Msgsize() (s int) {
	s = 1 + 7
	if z.Signee == nil {
		s += hsp.NilSize
	} else {
		s += z.Signee.Msgsize()
	}
	s += 10
	if z.Signature == nil {
		s += hsp.NilSize
	} else {
		s += z.Signature.Msgsize()
	}
	s += 19 + z.TokenReceiveHeader.Msgsize() + 11 + z.HeaderHash.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *TokenReceiveHeader) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 5
	o = append(o, 0x85, 0x85)
	if oTemp, err := z.Type.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x85)
	if oTemp, err := z.Amount.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x85)
	if oTemp, err := z.Nonce.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x85)
	if oTemp, err := z.Observer.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x85)
	if oTemp, err := z.Receiver.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TokenReceiveHeader) Msgsize() (s int) {
	s = 1 + 5 + z.Type.Msgsize() + 7 + z.Amount.Msgsize() + 6 + z.Nonce.Msgsize() + 9 + z.Observer.Msgsize() + 9 + z.Receiver.Msgsize()
	return
}