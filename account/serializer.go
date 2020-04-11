package account

type Serializer interface {
	Decode([]byte) (*Account, error)
	Encode(*Account) ([]byte, error)
}
