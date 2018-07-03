package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jeffprestes/livrocondolencias/contratos"
)

func main() {
	rede, err := ethclient.Dial(os.Getenv("ETH_RINKEBY_URL"))
	if err != nil {
		log.Fatalf("Nao foi possível se conectar a rede do Infura. Erro: %s\n", err.Error())
	}
	livroRegistro, err := contratos.NewRegistro(common.HexToAddress("0xE7402E8cDe5566c926f8103B793fa6eEEC13C2b5"), rede)
	if err != nil {
		log.Fatalf("Nao foi possível instanciar o contrato inteligente na rede do Infura. Erro: %s\n", err.Error())
	}
	ultimaMensagem, err := livroRegistro.Mensagem(nil)
	if err != nil {
		log.Fatalf("Nao foi possível obter a ultima mensagem registrada. Erro: %s\n", err.Error())
	}
	log.Println("Ultima mensagem de condolencias registrada: ", ultimaMensagem)
}
