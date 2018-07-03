package contratos

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

//TestDeployRegistro testa o deploy do smart contract registro
func TestDeployRegistro(t *testing.T) {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(100000000000)}
	blockchain := backends.NewSimulatedBackend(alloc)

	endereco, _, _, err := DeployRegistro(auth, blockchain)
	blockchain.Commit()

	if err != nil {
		t.Fatalf("Houve falha em fazer o deploy do contrato Registro: %+v", err)
	}

	if len(endereco.Bytes()) == 0 {
		t.Error("Um endereco valido era esperado. Recebido um endereco vazio")
	}
}
