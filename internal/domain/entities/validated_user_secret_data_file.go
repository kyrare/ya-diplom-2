package entities

type ValidatedUserSecretDataFile struct {
	UserSecretDataFile
	isValidated bool
}

func NewValidateUserSecretFile(secret *UserSecretDataFile) (*ValidatedUserSecretDataFile, error) {
	if err := secret.validate(); err != nil {
		return nil, err
	}

	return &ValidatedUserSecretDataFile{
		UserSecretDataFile: *secret,
		isValidated:        true,
	}, nil
}

func (us *ValidatedUserSecretDataFile) IsValid() bool {
	return us.isValidated
}
