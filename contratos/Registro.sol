pragma solidity 0.5.11;

contract Registro {
    string public mensagem;

    constructor () public {
        mensagem = "Uma boa e pacifica morte para todos...";
    }

    function RegistrarMensagem(string memory _mensagem) public {
        mensagem = _mensagem;
        emit NovaMensagem(mensagem);
    }

    event NovaMensagem(
        string _novaMsg
    );
}
