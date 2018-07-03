package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/jeffprestes/livrocondolencias/contratos"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	rede, err := ethclient.Dial("https://rinkeby.infura.io/QPF0qjGpH9OjFuuMrCse")
	if err != nil {
		log.Fatalf("Nao foi possível se conectar a rede do Infura. Erro: %s\n", err.Error())
	}
	/*
		Modelo abrindo a chave privada importada a sua instancia local do GETH
		arquivoComChavePrivada, err := os.Open(os.Getenv("ETH_ACCOUNT"))
		if err != nil {
			log.Fatalf("Nao foi possível abrir o arquivo com a chave privada. Erro: %s\n", err.Error())
		}
		auth, err := bind.NewTransactor(arquivoComChavePrivada, os.Getenv("ETH_ACCOUNT_PASS"))
		if err != nil {
			log.Fatalf("Nao foi possível gerar o transactor para assinar as transações. Erro: %s\n", err.Error())
		}
	*/
	prvKey, err := crypto.HexToECDSA(os.Getenv("ETH_ACCOUNT_PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Erro ao fazer o parse da chave privada . Erro: %s\n", err.Error())
	}
	auth := bind.NewKeyedTransactor(prvKey)
	endereco, _, _, err := contratos.DeployRegistro(auth, rede)
	if err != nil {
		log.Fatalf("Erro ao fazer o deploy do contrato na rede. Erro: %s\n", err.Error())
	}
	log.Printf("O endereço do contrato (deploy ainda pendente na rede) é: 0x%x\n", endereco)
}
