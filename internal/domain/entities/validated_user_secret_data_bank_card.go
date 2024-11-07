package entities

type ValidatedUserSecretDataBankCard struct {
	UserSecretDataBankCard

	isValidated bool
}

func NewValidateUserSecretBankCard(secret *UserSecretDataBankCard) (*ValidatedUserSecretDataBankCard, error) {
	if err := secret.validate(); err != nil {
		return nil, err
	}

	return &ValidatedUserSecretDataBankCard{
		UserSecretDataBankCard: *secret,
		isValidated:            true,
	}, nil
}

func (us *ValidatedUserSecretDataBankCard) IsValid() bool {
	return us.isValidated
}
