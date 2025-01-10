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
		panic("bytes is invalid")
	}

	return uuid.String()
}

func (*UuidProvider) FromString(text string) []byte {
	uuid := uuidGoogle.MustParse(text)

	bytes := [16]byte(uuid)

	return bytes[:]
}
