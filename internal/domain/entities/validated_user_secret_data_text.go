package entities

type ValidatedUserSecretDataText struct {
	UserSecretDataText
	isValidated bool
}

func NewValidateUserSecretText(secret *UserSecretDataText) (*ValidatedUserSecretDataText, error) {
	if err := secret.validate(); err != nil {
		return nil, err
	}

	return &ValidatedUserSecretDataText{
		UserSecretDataText: *secret,
		isValidated:        true,
	}, nil
}

func (us *ValidatedUserSecretDataText) IsValid() bool {
	return us.isValidated
}
