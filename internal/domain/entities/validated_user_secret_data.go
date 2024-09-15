package entities

type ValidatedUserDataSecret interface {
	UserSecretData
	IsValid() bool
}
