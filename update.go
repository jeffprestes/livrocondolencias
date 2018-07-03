package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jeffprestes/livrocondolencias/contratos"
)

func main() {
	rede, err := ethclient.Dial(os.Getenv("ETH_RINKEBY_URL"))
	if err != nil {
		log.Fatalf("Nao foi possível se conectar a rede do Infura. Erro: %s\n", err.Error())
	}

	livroRegistro, err := contratos.NewRegistro(common.HexToAddress("0xe7402e8cde5566c926f8103b793fa6eeec13c2b5"), rede)
	if err != nil {
		log.Fatalf("Nao foi possível instanciar o contrato inteligente na rede do Infura. Erro: %s\n", err.Error())
	}

	prvKey, err := crypto.HexToECDSA(os.Getenv("ETH_ACCOUNT_PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Erro ao fazer o parse da chave privada . Erro: %s\n", err.Error())
	}
	auth := bind.NewKeyedTransactor(prvKey)

	opcoesTransacao := &bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}

	_, err = livroRegistro.RegistrarMensagem(opcoesTransacao, "Uma saudacao daqui do Eliseo..")
	if err != nil {
		log.Fatalf("Erro ao registrar a mensagem. Erro: %s\n", err.Error())
	}
	log.Println("Mensagem registrada. Aguarde a mineração...")
}
