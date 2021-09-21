package key_generators

type KeyGenerator func() string

var keyGenerators = make(map[string]KeyGenerator)

func RegisterKeyGenerator(name string, keyGenerator KeyGenerator) {
	keyGenerators[name] = keyGenerator
}

func GetKeyGenerator(keyGeneratorType string) KeyGenerator {
	return keyGenerators[keyGeneratorType]
}

