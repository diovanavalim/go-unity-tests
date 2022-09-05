package addresses

import (
	"strings"
	"testing"
)

func TestAddressTypeRua(t *testing.T) {
	// t.Parallel() (Para rodar o teste em modo paralelo)

	address := "Rua Luiz Oscar de Carvalho"

	expectedAddressType := strings.Split(address, " ")[0]

	receivedAddressType := AddressType(address)

	if expectedAddressType != receivedAddressType {
		t.Errorf(
			"Received address type is different from expected address type. Expected %s, received %s",
			expectedAddressType,
			receivedAddressType,
		)
	}

	t.Log("Test Successfull")
}

func TestAddressTypeAvenida(t *testing.T) {
	// t.Parallel() (Para rodar o teste em modo paralelo)

	address := "Avenida Paraná"

	expectedAddressType := strings.Split(address, " ")[0]

	receivedAddressType := AddressType(address)

	if expectedAddressType != expectedAddressType {
		t.Errorf(
			"Received address type is different from expected address type. Expected %s, received %s",
			expectedAddressType,
			receivedAddressType,
		)
	}

	t.Log("Test Successfull")
}

func TestAddressTypeEstrada(t *testing.T) {
	address := "Estrada dos Açores"

	expectedAddressType := strings.Split(address, " ")[0]

	receivedAddressType := AddressType(address)

	if expectedAddressType != expectedAddressType {
		t.Errorf(
			"Received address type is different from expected address type. Expected %s, received %s",
			expectedAddressType,
			receivedAddressType,
		)
	}

	t.Log("Test Successfull")
}

func TestAddressTypeRodovia(t *testing.T) {
	address := "Rodovia SC 401"

	expectedAddressType := strings.Split(address, " ")[0]

	receivedAddressType := AddressType(address)

	if expectedAddressType != expectedAddressType {
		t.Errorf(
			"Received address type is different from expected address type. Expected %s, received %s",
			expectedAddressType,
			receivedAddressType,
		)
	}

	t.Log("Test Successfull")
}

func TestAddressTypeInvalido(t *testing.T) {
	address := "Apple MacBook Pro 2021"

	expectedAddressType := "Tipo inválido"

	receivedAddressType := AddressType(address)

	if expectedAddressType != expectedAddressType {
		t.Errorf(
			"Received address type is different from expected address type. Expected %s, received %s",
			expectedAddressType,
			receivedAddressType,
		)
	}

	t.Log("Test Successfull")
}

// Como rodar testes unitários?
// ./... indica para o compilador buscar arquivos de teste em todos os pacotes da aplicação.
// go test ./...
// go test ./... -v (para rodar em modo verbose)
// go test ./... --cover (para obter estatísticas de coverage)
// go test ./... -coverprofile coverage.txt (para obter estatísticas de coverage em um arquivo por linha de código)
// go tool cover --html=coverage.txt (para obter uma representação visual do coverage em linhas de código)
