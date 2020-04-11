/*
Package json implements the serializing and deserializing of user accounts.
*/
package json

import (
	"encoding/json"

	"github.com/BranDebs/Avocado-Backend/account"
	"github.com/pkg/errors"
)

var (
	ErrEncodeNilStruct    = errors.New("AccountSerializer.Encode: cannot encode nil struct")
	ErrDecodeNilByteSlice = errors.New("AccountSerializer.Decode: cannot decode nil bytes")
)

/*
AccountSerializer serializes/deserializes accounts in the json format.
*/
type AccountSerializer struct{}

func (*AccountSerializer) Encode(acct *account.Account) ([]byte, error) {
	if acct == nil {
		return nil, ErrEncodeNilStruct
	}
	return json.Marshal(&acct)
}

func (*AccountSerializer) Decode(data []byte) (*account.Account, error) {
	if data == nil {
		return nil, ErrDecodeNilByteSlice
	}
	var acc account.Account
	if err := json.Unmarshal(data, &acc); err != nil {
		return nil, errors.Wrap(err, "AccountSerializer.Decode")
	}
	return &acc, nil
}
