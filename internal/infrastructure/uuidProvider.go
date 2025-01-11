package infrastructure

import uuidGoogle "github.com/google/uuid"

type UuidProvider struct{}

func (*UuidProvider) Random() []byte {
	uuid := uuidGoogle.New()

	bytes := [16]byte(uuid)

	return bytes[:]
}

func (*UuidProvider) ToString(bytes []byte) string {
	uuid, err := uuidGoogle.FromBytes(bytes)

	if err != nil {
		panic(err.Error())
	}

	return uuid.String()
}

func (*UuidProvider) FromString(text string) ([]byte, error) {
	uuid, err := uuidGoogle.Parse(text)

	if err != nil {
		return nil, err
	}

	bytes := [16]byte(uuid)

	return bytes[:], nil
}
