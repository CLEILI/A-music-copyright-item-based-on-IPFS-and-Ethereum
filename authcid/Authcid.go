// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Authcid

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

// AuthcidMetaData contains all meta data concerning the Authcid contract.
var AuthcidMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"queryauth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061072a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063131a068014610046578063d5e943b914610062578063e4db2ceb14610092575b600080fd5b610060600480360381019061005b919061038f565b6100c2565b005b61007c6004803603810190610077919061038f565b6101d4565b604051610089919061041d565b60405180910390f35b6100ac60048036038101906100a79190610579565b6102cd565b6040516100b9919061041d565b60405180910390f35b600073ffffffffffffffffffffffffffffffffffffffff16600083836040516100ec9291906105f2565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610171576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016890610668565b60405180910390fd5b33600083836040516101849291906105f2565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff16600084846040516101ff9291906105f2565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610284576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027b906106d4565b60405180910390fd5b600083836040516102969291906105f2565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261034f5761034e61032a565b5b8235905067ffffffffffffffff81111561036c5761036b61032f565b5b60208301915083600182028301111561038857610387610334565b5b9250929050565b600080602083850312156103a6576103a5610320565b5b600083013567ffffffffffffffff8111156103c4576103c3610325565b5b6103d085828601610339565b92509250509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610407826103dc565b9050919050565b610417816103fc565b82525050565b6000602082019050610432600083018461040e565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6104868261043d565b810181811067ffffffffffffffff821117156104a5576104a461044e565b5b80604052505050565b60006104b8610316565b90506104c4828261047d565b919050565b600067ffffffffffffffff8211156104e4576104e361044e565b5b6104ed8261043d565b9050602081019050919050565b82818337600083830152505050565b600061051c610517846104c9565b6104ae565b90508281526020810184848401111561053857610537610438565b5b6105438482856104fa565b509392505050565b600082601f8301126105605761055f61032a565b5b8135610570848260208601610509565b91505092915050565b60006020828403121561058f5761058e610320565b5b600082013567ffffffffffffffff8111156105ad576105ac610325565b5b6105b98482850161054b565b91505092915050565b600081905092915050565b60006105d983856105c2565b93506105e68385846104fa565b82840190509392505050565b60006105ff8284866105cd565b91508190509392505050565b600082825260208201905092915050565b7f74686520636964206861732074686520617574686f7200000000000000000000600082015250565b600061065260168361060b565b915061065d8261061c565b602082019050919050565b6000602082019050818103600083015261068181610645565b9050919050565b7f7468652063696420646f6e27742068617665206120617574686f720000000000600082015250565b60006106be601b8361060b565b91506106c982610688565b602082019050919050565b600060208201905081810360008301526106ed816106b1565b905091905056fea2646970667358221220ecbbf1f0a01e46165c5d5b9a165daeccebd574da58f12d4e5a0f11b197a868cc64736f6c63430008120033",
}

// AuthcidABI is the input ABI used to generate the binding from.
// Deprecated: Use AuthcidMetaData.ABI instead.
var AuthcidABI = AuthcidMetaData.ABI

// AuthcidBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuthcidMetaData.Bin instead.
var AuthcidBin = AuthcidMetaData.Bin

// DeployAuthcid deploys a new Ethereum contract, binding an instance of Authcid to it.
func DeployAuthcid(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Authcid, error) {
	parsed, err := AuthcidMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuthcidBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Authcid{AuthcidCaller: AuthcidCaller{contract: contract}, AuthcidTransactor: AuthcidTransactor{contract: contract}, AuthcidFilterer: AuthcidFilterer{contract: contract}}, nil
}

// Authcid is an auto generated Go binding around an Ethereum contract.
type Authcid struct {
	AuthcidCaller     // Read-only binding to the contract
	AuthcidTransactor // Write-only binding to the contract
	AuthcidFilterer   // Log filterer for contract events
}

// AuthcidCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthcidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthcidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthcidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthcidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthcidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthcidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthcidSession struct {
	Contract     *Authcid          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthcidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthcidCallerSession struct {
	Contract *AuthcidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AuthcidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthcidTransactorSession struct {
	Contract     *AuthcidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AuthcidRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthcidRaw struct {
	Contract *Authcid // Generic contract binding to access the raw methods on
}

// AuthcidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthcidCallerRaw struct {
	Contract *AuthcidCaller // Generic read-only contract binding to access the raw methods on
}

// AuthcidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthcidTransactorRaw struct {
	Contract *AuthcidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthcid creates a new instance of Authcid, bound to a specific deployed contract.
func NewAuthcid(address common.Address, backend bind.ContractBackend) (*Authcid, error) {
	contract, err := bindAuthcid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Authcid{AuthcidCaller: AuthcidCaller{contract: contract}, AuthcidTransactor: AuthcidTransactor{contract: contract}, AuthcidFilterer: AuthcidFilterer{contract: contract}}, nil
}

// NewAuthcidCaller creates a new read-only instance of Authcid, bound to a specific deployed contract.
func NewAuthcidCaller(address common.Address, caller bind.ContractCaller) (*AuthcidCaller, error) {
	contract, err := bindAuthcid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthcidCaller{contract: contract}, nil
}

// NewAuthcidTransactor creates a new write-only instance of Authcid, bound to a specific deployed contract.
func NewAuthcidTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthcidTransactor, error) {
	contract, err := bindAuthcid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthcidTransactor{contract: contract}, nil
}

// NewAuthcidFilterer creates a new log filterer instance of Authcid, bound to a specific deployed contract.
func NewAuthcidFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthcidFilterer, error) {
	contract, err := bindAuthcid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthcidFilterer{contract: contract}, nil
}

// bindAuthcid binds a generic wrapper to an already deployed contract.
func bindAuthcid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuthcidMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authcid *AuthcidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Authcid.Contract.AuthcidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authcid *AuthcidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authcid.Contract.AuthcidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authcid *AuthcidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authcid.Contract.AuthcidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authcid *AuthcidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Authcid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authcid *AuthcidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authcid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authcid *AuthcidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authcid.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xe4db2ceb.
//
// Solidity: function auth(string ) view returns(address)
func (_Authcid *AuthcidCaller) Auth(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Authcid.contract.Call(opts, &out, "auth", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xe4db2ceb.
//
// Solidity: function auth(string ) view returns(address)
func (_Authcid *AuthcidSession) Auth(arg0 string) (common.Address, error) {
	return _Authcid.Contract.Auth(&_Authcid.CallOpts, arg0)
}

// Auth is a free data retrieval call binding the contract method 0xe4db2ceb.
//
// Solidity: function auth(string ) view returns(address)
func (_Authcid *AuthcidCallerSession) Auth(arg0 string) (common.Address, error) {
	return _Authcid.Contract.Auth(&_Authcid.CallOpts, arg0)
}

// Queryauth is a free data retrieval call binding the contract method 0xd5e943b9.
//
// Solidity: function queryauth(string cid) view returns(address)
func (_Authcid *AuthcidCaller) Queryauth(opts *bind.CallOpts, cid string) (common.Address, error) {
	var out []interface{}
	err := _Authcid.contract.Call(opts, &out, "queryauth", cid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Queryauth is a free data retrieval call binding the contract method 0xd5e943b9.
//
// Solidity: function queryauth(string cid) view returns(address)
func (_Authcid *AuthcidSession) Queryauth(cid string) (common.Address, error) {
	return _Authcid.Contract.Queryauth(&_Authcid.CallOpts, cid)
}

// Queryauth is a free data retrieval call binding the contract method 0xd5e943b9.
//
// Solidity: function queryauth(string cid) view returns(address)
func (_Authcid *AuthcidCallerSession) Queryauth(cid string) (common.Address, error) {
	return _Authcid.Contract.Queryauth(&_Authcid.CallOpts, cid)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string cid) returns()
func (_Authcid *AuthcidTransactor) Store(opts *bind.TransactOpts, cid string) (*types.Transaction, error) {
	return _Authcid.contract.Transact(opts, "store", cid)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string cid) returns()
func (_Authcid *AuthcidSession) Store(cid string) (*types.Transaction, error) {
	return _Authcid.Contract.Store(&_Authcid.TransactOpts, cid)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string cid) returns()
func (_Authcid *AuthcidTransactorSession) Store(cid string) (*types.Transaction, error) {
	return _Authcid.Contract.Store(&_Authcid.TransactOpts, cid)
}
