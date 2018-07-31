pragma solidity 0.4.24;

contract Testamento {
    
    struct Testamentario {
        string nome;
        uint heranca;
        mapping(address => Beneficiario) beneficiarios;
    }
    
    struct Beneficiario {
        uint percentual;
    }
    
    constructor() public payable {}
    
    mapping(address => Testamentario) public testamentarios;
    
    function aderir(string _name) public {
        Testamentario memory t = Testamentario(_name, 0);
        testamentarios[msg.sender] = t;
    }   
    
    function adicionarHeranca(address _recebedor) public payable {
        Testamentario tmt = testamentarios[_recebedor]; 
        tmt.heranca = tmt.heranca + msg.value;
    }
    
    function adicionarBeneficiario(address _beneficiario, uint _percentual) public {
        Testamentario tmt = testamentarios[msg.sender];
        Beneficiario memory benf = Beneficiario(_percentual);
        tmt.beneficiarios[_beneficiario] = benf;
    }
} 
    
