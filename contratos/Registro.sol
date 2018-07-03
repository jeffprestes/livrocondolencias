pragma solidity 0.4.24;

contract Registro {
    string public mensagem;

    constructor () public {
        mensagem = "Uma boa e pacifica morte para todos...";
    }

    function RegistrarMensagem(string _mensagem) public {
        mensagem = _mensagem;
    }
}