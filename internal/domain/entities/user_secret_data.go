package entities

type UserSecretData interface {
	validate() error
	GetType() UserSecretType
	GetData() ([]byte, error)
}
