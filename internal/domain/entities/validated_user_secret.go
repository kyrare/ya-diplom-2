package entities

type ValidatedUserSecret struct {
	UserSecret
	isValidated bool
}

func NewValidatedUserSecret(secret *UserSecret) (*ValidatedUserSecret, error) {
	if err := secret.validate(); err != nil {
		return nil, err
	}

	return &ValidatedUserSecret{
		UserSecret:  *secret,
		isValidated: true,
	}, nil
}

func (vp *ValidatedUserSecret) IsValid() bool {
	return vp.isValidated
}
