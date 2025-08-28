package contratos

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
)

func newTestBackend(t *testing.T) (*bind.TransactOpts, *simulated.Backend) {
	key, _ := crypto.GenerateKey()
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	auth.GasLimit = uint64(3000000)

	alloc := make(core.GenesisAlloc)
	balance, _ := new(big.Int).SetString("10000000000000000000", 10)
	alloc[auth.From] = core.GenesisAccount{Balance: balance} // 10 ether

	// Configure the chain for Shanghai
	chainConfig := *params.AllEthashProtocolChanges
	chainConfig.ChainID = big.NewInt(1337) // Simulated backend default
	shanghaiTime := uint64(0)
	chainConfig.ShanghaiTime = &shanghaiTime

	// Create the simulated backend
	backend := simulated.NewBackend(alloc, func(nodeConf *node.Config, ethConf *ethconfig.Config) {
		ethConf.Genesis = &core.Genesis{
			Config:   &chainConfig,
			Alloc:    alloc,
			GasLimit: 10000000,
		}
	})

	return auth, backend
}

//TestDeployContratos testa o deploy do smart contract registro
func TestDeployContratos(t *testing.T) {
	auth, backend := newTestBackend(t)

	endereco, _, _, err := DeployContratos(auth, backend.Client())
	if err != nil {
		t.Fatalf("Houve falha em fazer o deploy do contrato Contratos: %+v", err)
	}
	backend.Commit()

	if len(endereco.Bytes()) == 0 {
		t.Error("Um endereco valido era esperado. Recebido um endereco vazio")
	}
}

func TestGetContratos(t *testing.T) {
	auth, backend := newTestBackend(t)

	_, _, contrato, err := DeployContratos(auth, backend.Client())
	if err != nil {
		t.Fatalf("Houve falha em fazer o deploy do contrato Contratos: %+v", err)
	}
	backend.Commit()

	if msg, _ := contrato.Mensagem(nil); msg != "Uma boa e pacifica morte para todos..." {
		t.Errorf("Esperava-se a mensagem: Uma boa e pacifica morte para todos... e o retorno foi: %s\n", msg)
	}
}

func TestSetContratos(t *testing.T) {
	auth, backend := newTestBackend(t)

	_, _, contrato, err := DeployContratos(auth, backend.Client())
	if err != nil {
		t.Fatalf("Houve falha em fazer o deploy do contrato Contratos: %+v", err)
	}
	backend.Commit()

	opcoesTransacao := &bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		Value:    nil,
		GasLimit: uint64(3000000),
	}
	_, err = contrato.RegistrarMensagem(opcoesTransacao, "Uma passagem pacifica...")
	if err != nil {
		t.Fatalf("Houve falha em registrar a mensagem: %+v", err)
	}
	backend.Commit()

	if msg, _ := contrato.Mensagem(nil); msg != "Uma passagem pacifica..." {
		t.Errorf("Esperava-se a mensagem: Uma passagem pacifica... e o retorno foi: %s\n", msg)
	}
}
