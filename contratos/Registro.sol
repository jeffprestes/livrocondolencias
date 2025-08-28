// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Registro {
    string public mensagem;

    constructor() {
        mensagem = "Uma boa e pacifica morte para todos...";
    }

    function RegistrarMensagem(string calldata _mensagem) public {
        mensagem = _mensagem;
        emit NovaMensagem(mensagem);
    }

    event NovaMensagem(
        string _novaMsg
    );
}
