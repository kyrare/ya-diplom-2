package entities

type UserSecretDataFile struct {
}

var _ UserSecretData = &UserSecretDataFile{}

func NewUserSecretFile() *UserSecretDataFile {
	return &UserSecretDataFile{}
}

func (d *UserSecretDataFile) GetType() UserSecretType {
	return UserSecretFileType
}

func (d *UserSecretDataFile) GetData() ([]byte, error) {
	// todo
	panic("implement me")
	return []byte{}, nil
}

func (d *UserSecretDataFile) validate() error {
	//if d.FileName == "" {
	//	return errors.New("file name must not be empty")
	//}
	// todo
	panic("implement me")
	return nil
}
