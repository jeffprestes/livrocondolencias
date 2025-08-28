// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contratos

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContratosMetaData contains all meta data concerning the Contratos contract.
var ContratosMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_novaMsg\",\"type\":\"string\"}],\"name\":\"NovaMensagem\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_mensagem\",\"type\":\"string\"}],\"name\":\"RegistrarMensagem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mensagem\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405180606001604052806026815260200162000a1c602691395f90816200003a9190620002a5565b5062000389565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620000bd57607f821691505b602082108103620000d357620000d262000078565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620001377fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620000fa565b620001438683620000fa565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6200018d6200018762000181846200015b565b62000164565b6200015b565b9050919050565b5f819050919050565b620001a8836200016d565b620001c0620001b78262000194565b84845462000106565b825550505050565b5f90565b620001d6620001c8565b620001e38184846200019d565b505050565b5b818110156200020a57620001fe5f82620001cc565b600181019050620001e9565b5050565b601f82111562000259576200022381620000d9565b6200022e84620000eb565b810160208510156200023e578190505b620002566200024d85620000eb565b830182620001e8565b50505b505050565b5f82821c905092915050565b5f6200027b5f19846008026200025e565b1980831691505092915050565b5f6200029583836200026a565b9150826002028217905092915050565b620002b08262000041565b67ffffffffffffffff811115620002cc57620002cb6200004b565b5b620002d88254620000a5565b620002e58282856200020e565b5f60209050601f8311600181146200031b575f841562000306578287015190505b62000312858262000288565b86555062000381565b601f1984166200032b86620000d9565b5f5b8281101562000354578489015182556001820191506020850194506020810190506200032d565b8683101562000374578489015162000370601f8916826200026a565b8355505b6001600288020188555050505b505050505050565b61068580620003975f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063160c506314610038578063c6c6572814610056575b5f80fd5b610040610072565b60405161004d91906101d3565b60405180910390f35b610070600480360381019061006b919061025c565b6100fd565b005b5f805461007e906102d4565b80601f01602080910402602001604051908101604052809291908181526020018280546100aa906102d4565b80156100f55780601f106100cc576101008083540402835291602001916100f5565b820191905f5260205f20905b8154815290600101906020018083116100d857829003601f168201915b505050505081565b81815f918261010d9291906104e1565b507f3ea7d0fcf026f7eb4d355a230e84ee80e817b99ed1757333c7a24592915bc7c45f60405161013d919061062f565b60405180910390a15050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610180578082015181840152602081019050610165565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6101a582610149565b6101af8185610153565b93506101bf818560208601610163565b6101c88161018b565b840191505092915050565b5f6020820190508181035f8301526101eb818461019b565b905092915050565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261021c5761021b6101fb565b5b8235905067ffffffffffffffff811115610239576102386101ff565b5b60208301915083600182028301111561025557610254610203565b5b9250929050565b5f8060208385031215610272576102716101f3565b5b5f83013567ffffffffffffffff81111561028f5761028e6101f7565b5b61029b85828601610207565b92509250509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806102eb57607f821691505b6020821081036102fe576102fd6102a7565b5b50919050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103977fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261035c565b6103a1868361035c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6103e56103e06103db846103b9565b6103c2565b6103b9565b9050919050565b5f819050919050565b6103fe836103cb565b61041261040a826103ec565b848454610368565b825550505050565b5f90565b61042661041a565b6104318184846103f5565b505050565b5b81811015610454576104495f8261041e565b600181019050610437565b5050565b601f8211156104995761046a8161033b565b6104738461034d565b81016020851015610482578190505b61049661048e8561034d565b830182610436565b50505b505050565b5f82821c905092915050565b5f6104b95f198460080261049e565b1980831691505092915050565b5f6104d183836104aa565b9150826002028217905092915050565b6104eb8383610304565b67ffffffffffffffff8111156105045761050361030e565b5b61050e82546102d4565b610519828285610458565b5f601f831160018114610546575f8415610534578287013590505b61053e85826104c6565b8655506105a5565b601f1984166105548661033b565b5f5b8281101561057b57848901358255600182019150602085019450602081019050610556565b868310156105985784890135610594601f8916826104aa565b8355505b6001600288020188555050505b50505050505050565b5f81546105ba816102d4565b6105c48186610153565b9450600182165f81146105de57600181146105f457610626565b60ff198316865281151560200286019350610626565b6105fd8561033b565b5f5b8381101561061e578154818901526001820191506020810190506105ff565b808801955050505b50505092915050565b5f6020820190508181035f83015261064781846105ae565b90509291505056fea264697066735822122026bd3db5a938b1d9cb873377ef69bede2cebd4e9bf9e2cd8c3f0af53d0f01e4764736f6c63430008180033556d6120626f612065207061636966696361206d6f727465207061726120746f646f732e2e2e",
}

// ContratosABI is the input ABI used to generate the binding from.
// Deprecated: Use ContratosMetaData.ABI instead.
var ContratosABI = ContratosMetaData.ABI

// ContratosBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContratosMetaData.Bin instead.
var ContratosBin = ContratosMetaData.Bin

// DeployContratos deploys a new Ethereum contract, binding an instance of Contratos to it.
func DeployContratos(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contratos, error) {
	parsed, err := ContratosMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContratosBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contratos{ContratosCaller: ContratosCaller{contract: contract}, ContratosTransactor: ContratosTransactor{contract: contract}, ContratosFilterer: ContratosFilterer{contract: contract}}, nil
}

// Contratos is an auto generated Go binding around an Ethereum contract.
type Contratos struct {
	ContratosCaller     // Read-only binding to the contract
	ContratosTransactor // Write-only binding to the contract
	ContratosFilterer   // Log filterer for contract events
}

// ContratosCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContratosCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContratosTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContratosTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContratosFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContratosFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContratosSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContratosSession struct {
	Contract     *Contratos        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContratosCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContratosCallerSession struct {
	Contract *ContratosCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContratosTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContratosTransactorSession struct {
	Contract     *ContratosTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContratosRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContratosRaw struct {
	Contract *Contratos // Generic contract binding to access the raw methods on
}

// ContratosCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContratosCallerRaw struct {
	Contract *ContratosCaller // Generic read-only contract binding to access the raw methods on
}

// ContratosTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContratosTransactorRaw struct {
	Contract *ContratosTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContratos creates a new instance of Contratos, bound to a specific deployed contract.
func NewContratos(address common.Address, backend bind.ContractBackend) (*Contratos, error) {
	contract, err := bindContratos(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contratos{ContratosCaller: ContratosCaller{contract: contract}, ContratosTransactor: ContratosTransactor{contract: contract}, ContratosFilterer: ContratosFilterer{contract: contract}}, nil
}

// NewContratosCaller creates a new read-only instance of Contratos, bound to a specific deployed contract.
func NewContratosCaller(address common.Address, caller bind.ContractCaller) (*ContratosCaller, error) {
	contract, err := bindContratos(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContratosCaller{contract: contract}, nil
}

// NewContratosTransactor creates a new write-only instance of Contratos, bound to a specific deployed contract.
func NewContratosTransactor(address common.Address, transactor bind.ContractTransactor) (*ContratosTransactor, error) {
	contract, err := bindContratos(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContratosTransactor{contract: contract}, nil
}

// NewContratosFilterer creates a new log filterer instance of Contratos, bound to a specific deployed contract.
func NewContratosFilterer(address common.Address, filterer bind.ContractFilterer) (*ContratosFilterer, error) {
	contract, err := bindContratos(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContratosFilterer{contract: contract}, nil
}

// bindContratos binds a generic wrapper to an already deployed contract.
func bindContratos(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContratosMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contratos *ContratosRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contratos.Contract.ContratosCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contratos *ContratosRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contratos.Contract.ContratosTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contratos *ContratosRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contratos.Contract.ContratosTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contratos *ContratosCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contratos.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contratos *ContratosTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contratos.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contratos *ContratosTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contratos.Contract.contract.Transact(opts, method, params...)
}

// Mensagem is a free data retrieval call binding the contract method 0x160c5063.
//
// Solidity: function mensagem() view returns(string)
func (_Contratos *ContratosCaller) Mensagem(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contratos.contract.Call(opts, &out, "mensagem")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Mensagem is a free data retrieval call binding the contract method 0x160c5063.
//
// Solidity: function mensagem() view returns(string)
func (_Contratos *ContratosSession) Mensagem() (string, error) {
	return _Contratos.Contract.Mensagem(&_Contratos.CallOpts)
}

// Mensagem is a free data retrieval call binding the contract method 0x160c5063.
//
// Solidity: function mensagem() view returns(string)
func (_Contratos *ContratosCallerSession) Mensagem() (string, error) {
	return _Contratos.Contract.Mensagem(&_Contratos.CallOpts)
}

// RegistrarMensagem is a paid mutator transaction binding the contract method 0xc6c65728.
//
// Solidity: function RegistrarMensagem(string _mensagem) returns()
func (_Contratos *ContratosTransactor) RegistrarMensagem(opts *bind.TransactOpts, _mensagem string) (*types.Transaction, error) {
	return _Contratos.contract.Transact(opts, "RegistrarMensagem", _mensagem)
}

// RegistrarMensagem is a paid mutator transaction binding the contract method 0xc6c65728.
//
// Solidity: function RegistrarMensagem(string _mensagem) returns()
func (_Contratos *ContratosSession) RegistrarMensagem(_mensagem string) (*types.Transaction, error) {
	return _Contratos.Contract.RegistrarMensagem(&_Contratos.TransactOpts, _mensagem)
}

// RegistrarMensagem is a paid mutator transaction binding the contract method 0xc6c65728.
//
// Solidity: function RegistrarMensagem(string _mensagem) returns()
func (_Contratos *ContratosTransactorSession) RegistrarMensagem(_mensagem string) (*types.Transaction, error) {
	return _Contratos.Contract.RegistrarMensagem(&_Contratos.TransactOpts, _mensagem)
}

// ContratosNovaMensagemIterator is returned from FilterNovaMensagem and is used to iterate over the raw logs and unpacked data for NovaMensagem events raised by the Contratos contract.
type ContratosNovaMensagemIterator struct {
	Event *ContratosNovaMensagem // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContratosNovaMensagemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContratosNovaMensagem)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContratosNovaMensagem)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContratosNovaMensagemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContratosNovaMensagemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContratosNovaMensagem represents a NovaMensagem event raised by the Contratos contract.
type ContratosNovaMensagem struct {
	NovaMsg string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNovaMensagem is a free log retrieval operation binding the contract event 0x3ea7d0fcf026f7eb4d355a230e84ee80e817b99ed1757333c7a24592915bc7c4.
//
// Solidity: event NovaMensagem(string _novaMsg)
func (_Contratos *ContratosFilterer) FilterNovaMensagem(opts *bind.FilterOpts) (*ContratosNovaMensagemIterator, error) {

	logs, sub, err := _Contratos.contract.FilterLogs(opts, "NovaMensagem")
	if err != nil {
		return nil, err
	}
	return &ContratosNovaMensagemIterator{contract: _Contratos.contract, event: "NovaMensagem", logs: logs, sub: sub}, nil
}

// WatchNovaMensagem is a free log subscription operation binding the contract event 0x3ea7d0fcf026f7eb4d355a230e84ee80e817b99ed1757333c7a24592915bc7c4.
//
// Solidity: event NovaMensagem(string _novaMsg)
func (_Contratos *ContratosFilterer) WatchNovaMensagem(opts *bind.WatchOpts, sink chan<- *ContratosNovaMensagem) (event.Subscription, error) {

	logs, sub, err := _Contratos.contract.WatchLogs(opts, "NovaMensagem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContratosNovaMensagem)
				if err := _Contratos.contract.UnpackLog(event, "NovaMensagem", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNovaMensagem is a log parse operation binding the contract event 0x3ea7d0fcf026f7eb4d355a230e84ee80e817b99ed1757333c7a24592915bc7c4.
//
// Solidity: event NovaMensagem(string _novaMsg)
func (_Contratos *ContratosFilterer) ParseNovaMensagem(log types.Log) (*ContratosNovaMensagem, error) {
	event := new(ContratosNovaMensagem)
	if err := _Contratos.contract.UnpackLog(event, "NovaMensagem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
