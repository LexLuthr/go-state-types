// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package account

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufState = []byte{129}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufState); err != nil {
		return err
	}

	// t.Address (address.Address) (struct)
	if err := t.Address.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) (err error) {
	*t = State{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Address (address.Address) (struct)

	{

		if err := t.Address.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Address: %w", err)
		}

	}
	return nil
}

var lengthBufAuthenticateMessageParams = []byte{130}

func (t *AuthenticateMessageParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufAuthenticateMessageParams); err != nil {
		return err
	}

	// t.Signature ([]uint8) (slice)
	if len(t.Signature) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Signature was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Signature))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Signature); err != nil {
		return err
	}

	// t.Message ([]uint8) (slice)
	if len(t.Message) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Message was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Message))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Message); err != nil {
		return err
	}

	return nil
}

func (t *AuthenticateMessageParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = AuthenticateMessageParams{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Signature ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Signature: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Signature = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Signature); err != nil {
		return err
	}

	// t.Message ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Message: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Message = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Message); err != nil {
		return err
	}

	return nil
}
