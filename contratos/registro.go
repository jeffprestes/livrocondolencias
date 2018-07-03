// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contratos

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// RegistroABI is the input ABI used to generate the binding from.
const RegistroABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mensagem\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_mensagem\",\"type\":\"string\"}],\"name\":\"RegistrarMensagem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_novaMsg\",\"type\":\"string\"}],\"name\":\"NovaMensagem\",\"type\":\"event\"}]"

// RegistroBin is the compiled bytecode used for deploying new contracts.
const RegistroBin = `0x608060405234801561001057600080fd5b506040805160608101825260268082527f556d6120626f612065207061636966696361206d6f727465207061726120746f602083019081527f646f732e2e2e0000000000000000000000000000000000000000000000000000929093019190915261007d91600091610083565b5061011e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c457805160ff19168380011785556100f1565b828001600101855582156100f1579182015b828111156100f15782518255916020019190600101906100d6565b506100fd929150610101565b5090565b61011b91905b808211156100fd5760008155600101610107565b90565b61034e8061012d6000396000f30060806040526004361061004b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663160c50638114610050578063c6c65728146100da575b600080fd5b34801561005c57600080fd5b50610065610135565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561009f578181015183820152602001610087565b50505050905090810190601f1680156100cc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156100e657600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101339436949293602493928401919081908401838280828437509497506101c39650505050505050565b005b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101bb5780601f10610190576101008083540402835291602001916101bb565b820191906000526020600020905b81548152906001019060200180831161019e57829003601f168201915b505050505081565b80516101d6906000906020840190610287565b5060408051602080825260008054600260001961010060018416150201909116049183018290527f3ea7d0fcf026f7eb4d355a230e84ee80e817b99ed1757333c7a24592915bc7c4939092918291820190849080156102765780601f1061024b57610100808354040283529160200191610276565b820191906000526020600020905b81548152906001019060200180831161025957829003601f168201915b50509250505060405180910390a150565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106102c857805160ff19168380011785556102f5565b828001600101855582156102f5579182015b828111156102f55782518255916020019190600101906102da565b50610301929150610305565b5090565b61031f91905b80821115610301576000815560010161030b565b905600a165627a7a723058203f824513aae94b22849882c2b8d6898ea79a51d467821007a99571f964ba63b30029`

// DeployRegistro deploys a new Ethereum contract, binding an instance of Registro to it.
func DeployRegistro(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registro, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistroABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistroBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registro{RegistroCaller: RegistroCaller{contract: contract}, RegistroTransactor: RegistroTransactor{contract: contract}, RegistroFilterer: RegistroFilterer{contract: contract}}, nil
}

// Registro is an auto generated Go binding around an Ethereum contract.
type Registro struct {
	RegistroCaller     // Read-only binding to the contract
	RegistroTransactor // Write-only binding to the contract
	RegistroFilterer   // Log filterer for contract events
}

// RegistroCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistroCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistroTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistroTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistroFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistroFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistroSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistroSession struct {
	Contract     *Registro         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistroCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistroCallerSession struct {
	Contract *RegistroCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistroTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistroTransactorSession struct {
	Contract     *RegistroTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistroRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistroRaw struct {
	Contract *Registro // Generic contract binding to access the raw methods on
}

// RegistroCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistroCallerRaw struct {
	Contract *RegistroCaller // Generic read-only contract binding to access the raw methods on
}

// RegistroTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistroTransactorRaw struct {
	Contract *RegistroTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistro creates a new instance of Registro, bound to a specific deployed contract.
func NewRegistro(address common.Address, backend bind.ContractBackend) (*Registro, error) {
	contract, err := bindRegistro(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registro{RegistroCaller: RegistroCaller{contract: contract}, RegistroTransactor: RegistroTransactor{contract: contract}, RegistroFilterer: RegistroFilterer{contract: contract}}, nil
}

// NewRegistroCaller creates a new read-only instance of Registro, bound to a specific deployed contract.
func NewRegistroCaller(address common.Address, caller bind.ContractCaller) (*RegistroCaller, error) {
	contract, err := bindRegistro(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistroCaller{contract: contract}, nil
}

// NewRegistroTransactor creates a new write-only instance of Registro, bound to a specific deployed contract.
func NewRegistroTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistroTransactor, error) {
	contract, err := bindRegistro(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistroTransactor{contract: contract}, nil
}

// NewRegistroFilterer creates a new log filterer instance of Registro, bound to a specific deployed contract.
func NewRegistroFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistroFilterer, error) {
	contract, err := bindRegistro(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistroFilterer{contract: contract}, nil
}

// bindRegistro binds a generic wrapper to an already deployed contract.
func bindRegistro(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistroABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registro *RegistroRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registro.Contract.RegistroCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registro *RegistroRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registro.Contract.RegistroTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registro *RegistroRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registro.Contract.RegistroTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registro *RegistroCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registro.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registro *RegistroTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registro.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registro *RegistroTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registro.Contract.contract.Transact(opts, method, params...)
}

// Mensagem is a free data retrieval call binding the contract method 0x160c5063.
//
// Solidity: function mensagem() constant returns(string)
func (_Registro *RegistroCaller) Mensagem(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Registro.contract.Call(opts, out, "mensagem")
	return *ret0, err
}

// Mensagem is a free data retrieval call binding the contract method 0x160c5063.
//
// Solidity: function mensagem() constant returns(string)
func (_Registro *RegistroSession) Mensagem() (string, error) {
	return _Registro.Contract.Mensagem(&_Registro.CallOpts)
}

// Mensagem is a free data retrieval call binding the contract method 0x160c5063.
//
// Solidity: function mensagem() constant returns(string)
func (_Registro *RegistroCallerSession) Mensagem() (string, error) {
	return _Registro.Contract.Mensagem(&_Registro.CallOpts)
}

// RegistrarMensagem is a paid mutator transaction binding the contract method 0xc6c65728.
//
// Solidity: function RegistrarMensagem(_mensagem string) returns()
func (_Registro *RegistroTransactor) RegistrarMensagem(opts *bind.TransactOpts, _mensagem string) (*types.Transaction, error) {
	return _Registro.contract.Transact(opts, "RegistrarMensagem", _mensagem)
}

// RegistrarMensagem is a paid mutator transaction binding the contract method 0xc6c65728.
//
// Solidity: function RegistrarMensagem(_mensagem string) returns()
func (_Registro *RegistroSession) RegistrarMensagem(_mensagem string) (*types.Transaction, error) {
	return _Registro.Contract.RegistrarMensagem(&_Registro.TransactOpts, _mensagem)
}

// RegistrarMensagem is a paid mutator transaction binding the contract method 0xc6c65728.
//
// Solidity: function RegistrarMensagem(_mensagem string) returns()
func (_Registro *RegistroTransactorSession) RegistrarMensagem(_mensagem string) (*types.Transaction, error) {
	return _Registro.Contract.RegistrarMensagem(&_Registro.TransactOpts, _mensagem)
}

// RegistroNovaMensagemIterator is returned from FilterNovaMensagem and is used to iterate over the raw logs and unpacked data for NovaMensagem events raised by the Registro contract.
type RegistroNovaMensagemIterator struct {
	Event *RegistroNovaMensagem // Event containing the contract specifics and raw log

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
func (it *RegistroNovaMensagemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistroNovaMensagem)
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
		it.Event = new(RegistroNovaMensagem)
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
func (it *RegistroNovaMensagemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistroNovaMensagemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistroNovaMensagem represents a NovaMensagem event raised by the Registro contract.
type RegistroNovaMensagem struct {
	NovaMsg string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNovaMensagem is a free log retrieval operation binding the contract event 0x3ea7d0fcf026f7eb4d355a230e84ee80e817b99ed1757333c7a24592915bc7c4.
//
// Solidity: e NovaMensagem(_novaMsg string)
func (_Registro *RegistroFilterer) FilterNovaMensagem(opts *bind.FilterOpts) (*RegistroNovaMensagemIterator, error) {

	logs, sub, err := _Registro.contract.FilterLogs(opts, "NovaMensagem")
	if err != nil {
		return nil, err
	}
	return &RegistroNovaMensagemIterator{contract: _Registro.contract, event: "NovaMensagem", logs: logs, sub: sub}, nil
}

// WatchNovaMensagem is a free log subscription operation binding the contract event 0x3ea7d0fcf026f7eb4d355a230e84ee80e817b99ed1757333c7a24592915bc7c4.
//
// Solidity: e NovaMensagem(_novaMsg string)
func (_Registro *RegistroFilterer) WatchNovaMensagem(opts *bind.WatchOpts, sink chan<- *RegistroNovaMensagem) (event.Subscription, error) {

	logs, sub, err := _Registro.contract.WatchLogs(opts, "NovaMensagem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistroNovaMensagem)
				if err := _Registro.contract.UnpackLog(event, "NovaMensagem", log); err != nil {
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
