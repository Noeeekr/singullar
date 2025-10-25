package commandutil

type RegisterToken uint8

var token RegisterToken = 0

// "RegisterRegisterToken" allows to register a token through different files without having to worry about value conflict
func RegisterConfigurationToken() RegisterToken {
	token++
	return RegisterToken(token - 1)
}
