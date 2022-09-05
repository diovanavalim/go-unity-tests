package addresses

import "strings"

func AddressType(address string) string {
	validAddresses := []string{"rua", "avenida", "estrada", "rodovia"}

	lowerAddress := strings.ToLower(address)

	firstWord := strings.Split(lowerAddress, " ")[0]

	valid := false

	for _, addrType := range validAddresses {
		if addrType == firstWord {
			valid = true
		}
	}

	if valid {
		return strings.Title(firstWord)
	}

	return "Tipo inválido"
}
