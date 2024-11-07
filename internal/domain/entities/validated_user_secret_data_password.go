package entities

type ValidatedUserSecretDataPassword struct {
	UserSecretDataPassword
	isValidated bool
}

func NewValidateUserSecretPassword(secret *UserSecretDataPassword) (*ValidatedUserSecretDataPassword, error) {
	if err := secret.validate(); err != nil {
		return nil, err
	}

	return &ValidatedUserSecretDataPassword{
		UserSecretDataPassword: *secret,
		isValidated:            true,
	}, nil
}

func (us *ValidatedUserSecretDataPassword) IsValid() bool {
	return us.isValidated
}
