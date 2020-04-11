package json

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/BranDebs/Avocado-Backend/account"
	"github.com/pkg/errors"
)

func TestAccountEncode(t *testing.T) {
	s := AccountSerializer{}

	type output struct {
		data []byte
		err  error
	}

	time := time.Date(2018, 2, 8, 0, 0, 0, 0, time.Local)

	tests := []struct {
		name string
		in   *account.Account
		want output
	}{
		{
			"nil account struct",
			nil,
			output{nil, ErrEncodeNilStruct},
		},
		{
			"account with email",
			&account.Account{Email: "test@email.com"},
			output{[]byte(`{"email":"test@email.com","name":"","created_at":null}`), nil},
		},
		{
			"account with all data",
			&account.Account{Email: "test@email.com", Name: "test", CreatedAt: &time},
			output{[]byte(`{"email":"test@email.com","name":"test","created_at":"2018-02-08T00:00:00+08:00"}`), nil}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, err := s.Encode(test.in)
			if err != nil && err != test.want.err {
				t.Errorf("want (%v) got (%v)", test.want.err, err)
				return
			}
			if !bytes.Equal(data, test.want.data) {
				t.Errorf("want (%s) got (%s)", string(test.want.data), string(data))
				return
			}
		})
	}
}

func TestAccountDecode(t *testing.T) {
	s := AccountSerializer{}

	type output struct {
		account *account.Account
		err     error
	}

	time := time.Date(2018, 2, 8, 0, 0, 0, 0, time.Local)

	tests := []struct {
		name string
		in   []byte
		want output
	}{
		{"nil byte slice", nil, output{err: ErrDecodeNilByteSlice}},
		{"empty byte slice", []byte{}, output{err: errors.New("AccountSerializer.Decode: unexpected end of JSON input")}},
		{
			"proper Account byte slice",
			[]byte(`{"email":"test@email.com","name":"test","created_at":"2018-02-08T00:00:00+08:00"}`),
			output{
				account: &account.Account{
					Email:     "test@email.com",
					Name:      "test",
					CreatedAt: &time,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			acc, err := s.Decode(test.in)
			if err != nil && err.Error() != test.want.err.Error() {
				t.Errorf("want err (%v) got err (%v)", test.want.err, err)
				return
			}

			if !reflect.DeepEqual(test.want.account, acc) {
				t.Errorf("want (%v) got (%v)", *test.want.account, *acc)
				return
			}
		})
	}
}
