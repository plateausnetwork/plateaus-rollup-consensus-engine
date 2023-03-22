// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// LotteryValidationMetaData contains all meta data concerning the LotteryValidation contract.
var LotteryValidationMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"payable\":false,\"inputs\":[]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"approved\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokenId\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"ApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"operator\",\"indexed\":true},{\"type\":\"bool\",\"name\":\"approved\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokenId\",\"indexed\":true}]},{\"type\":\"function\",\"name\":\"approve\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"to\"},{\"type\":\"uint256\",\"name\":\"tokenId\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"balanceOf\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"}],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"getApproved\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\"}],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"},{\"type\":\"address\",\"name\":\"operator\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"mint\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"imageBase64\"},{\"type\":\"string\",\"name\":\"hash\"},{\"type\":\"string\",\"name\":\"externalURL\"},{\"type\":\"uint256\",\"name\":\"minHeight\"},{\"type\":\"uint256\",\"name\":\"maxHeight\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"name\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"string\"}]},{\"type\":\"function\",\"name\":\"ownerOf\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\"}],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"from\"},{\"type\":\"address\",\"name\":\"to\"},{\"type\":\"uint256\",\"name\":\"tokenId\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"from\"},{\"type\":\"address\",\"name\":\"to\"},{\"type\":\"uint256\",\"name\":\"tokenId\"},{\"type\":\"bytes\",\"name\":\"data\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"operator\"},{\"type\":\"bool\",\"name\":\"approved\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"supportsInterface\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"bytes4\",\"name\":\"interfaceId\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"symbol\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"string\"}]},{\"type\":\"function\",\"name\":\"tokenURI\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\"}],\"outputs\":[{\"type\":\"string\"}]},{\"type\":\"function\",\"name\":\"transferFrom\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"from\"},{\"type\":\"address\",\"name\":\"to\"},{\"type\":\"uint256\",\"name\":\"tokenId\"}],\"outputs\":[]}]",
}

// LotteryValidationABI is the input ABI used to generate the binding from.
// Deprecated: Use LotteryValidationMetaData.ABI instead.
var LotteryValidationABI = LotteryValidationMetaData.ABI

// LotteryValidation is an auto generated Go binding around an Ethereum contract.
type LotteryValidation struct {
	LotteryValidationCaller     // Read-only binding to the contract
	LotteryValidationTransactor // Write-only binding to the contract
	LotteryValidationFilterer   // Log filterer for contract events
}

// LotteryValidationCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryValidationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryValidationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LotteryValidationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryValidationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LotteryValidationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryValidationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LotteryValidationSession struct {
	Contract     *LotteryValidation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LotteryValidationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LotteryValidationCallerSession struct {
	Contract *LotteryValidationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LotteryValidationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LotteryValidationTransactorSession struct {
	Contract     *LotteryValidationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LotteryValidationRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryValidationRaw struct {
	Contract *LotteryValidation // Generic contract binding to access the raw methods on
}

// LotteryValidationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryValidationCallerRaw struct {
	Contract *LotteryValidationCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryValidationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LotteryValidationTransactorRaw struct {
	Contract *LotteryValidationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLotteryValidation creates a new instance of LotteryValidation, bound to a specific deployed contract.
func NewLotteryValidation(address common.Address, backend bind.ContractBackend) (*LotteryValidation, error) {
	contract, err := bindLotteryValidation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LotteryValidation{LotteryValidationCaller: LotteryValidationCaller{contract: contract}, LotteryValidationTransactor: LotteryValidationTransactor{contract: contract}, LotteryValidationFilterer: LotteryValidationFilterer{contract: contract}}, nil
}

// NewLotteryValidationCaller creates a new read-only instance of LotteryValidation, bound to a specific deployed contract.
func NewLotteryValidationCaller(address common.Address, caller bind.ContractCaller) (*LotteryValidationCaller, error) {
	contract, err := bindLotteryValidation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryValidationCaller{contract: contract}, nil
}

// NewLotteryValidationTransactor creates a new write-only instance of LotteryValidation, bound to a specific deployed contract.
func NewLotteryValidationTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryValidationTransactor, error) {
	contract, err := bindLotteryValidation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryValidationTransactor{contract: contract}, nil
}

// NewLotteryValidationFilterer creates a new log filterer instance of LotteryValidation, bound to a specific deployed contract.
func NewLotteryValidationFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryValidationFilterer, error) {
	contract, err := bindLotteryValidation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryValidationFilterer{contract: contract}, nil
}

// bindLotteryValidation binds a generic wrapper to an already deployed contract.
func bindLotteryValidation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LotteryValidationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LotteryValidation *LotteryValidationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LotteryValidation.Contract.LotteryValidationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LotteryValidation *LotteryValidationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryValidation.Contract.LotteryValidationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LotteryValidation *LotteryValidationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LotteryValidation.Contract.LotteryValidationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LotteryValidation *LotteryValidationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LotteryValidation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LotteryValidation *LotteryValidationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LotteryValidation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LotteryValidation *LotteryValidationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LotteryValidation.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LotteryValidation *LotteryValidationCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LotteryValidation.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LotteryValidation *LotteryValidationSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LotteryValidation.Contract.BalanceOf(&_LotteryValidation.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LotteryValidation *LotteryValidationCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LotteryValidation.Contract.BalanceOf(&_LotteryValidation.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LotteryValidation *LotteryValidationCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LotteryValidation.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LotteryValidation *LotteryValidationSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LotteryValidation.Contract.GetApproved(&_LotteryValidation.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LotteryValidation *LotteryValidationCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LotteryValidation.Contract.GetApproved(&_LotteryValidation.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LotteryValidation *LotteryValidationCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LotteryValidation.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LotteryValidation *LotteryValidationSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LotteryValidation.Contract.IsApprovedForAll(&_LotteryValidation.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LotteryValidation *LotteryValidationCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LotteryValidation.Contract.IsApprovedForAll(&_LotteryValidation.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LotteryValidation *LotteryValidationCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LotteryValidation.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LotteryValidation *LotteryValidationSession) Name() (string, error) {
	return _LotteryValidation.Contract.Name(&_LotteryValidation.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LotteryValidation *LotteryValidationCallerSession) Name() (string, error) {
	return _LotteryValidation.Contract.Name(&_LotteryValidation.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LotteryValidation *LotteryValidationCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LotteryValidation.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LotteryValidation *LotteryValidationSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LotteryValidation.Contract.OwnerOf(&_LotteryValidation.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LotteryValidation *LotteryValidationCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LotteryValidation.Contract.OwnerOf(&_LotteryValidation.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LotteryValidation *LotteryValidationCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LotteryValidation.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LotteryValidation *LotteryValidationSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LotteryValidation.Contract.SupportsInterface(&_LotteryValidation.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LotteryValidation *LotteryValidationCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LotteryValidation.Contract.SupportsInterface(&_LotteryValidation.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LotteryValidation *LotteryValidationCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LotteryValidation.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LotteryValidation *LotteryValidationSession) Symbol() (string, error) {
	return _LotteryValidation.Contract.Symbol(&_LotteryValidation.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LotteryValidation *LotteryValidationCallerSession) Symbol() (string, error) {
	return _LotteryValidation.Contract.Symbol(&_LotteryValidation.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LotteryValidation *LotteryValidationCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LotteryValidation.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LotteryValidation *LotteryValidationSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LotteryValidation.Contract.TokenURI(&_LotteryValidation.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LotteryValidation *LotteryValidationCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LotteryValidation.Contract.TokenURI(&_LotteryValidation.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LotteryValidation *LotteryValidationTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LotteryValidation *LotteryValidationSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.Contract.Approve(&_LotteryValidation.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LotteryValidation *LotteryValidationTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.Contract.Approve(&_LotteryValidation.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x89ca53a8.
//
// Solidity: function mint(string imageBase64, string hash, string externalURL, uint256 minHeight, uint256 maxHeight) returns()
func (_LotteryValidation *LotteryValidationTransactor) Mint(opts *bind.TransactOpts, imageBase64 string, hash string, externalURL string, minHeight *big.Int, maxHeight *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.contract.Transact(opts, "mint", imageBase64, hash, externalURL, minHeight, maxHeight)
}

// Mint is a paid mutator transaction binding the contract method 0x89ca53a8.
//
// Solidity: function mint(string imageBase64, string hash, string externalURL, uint256 minHeight, uint256 maxHeight) returns()
func (_LotteryValidation *LotteryValidationSession) Mint(imageBase64 string, hash string, externalURL string, minHeight *big.Int, maxHeight *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.Contract.Mint(&_LotteryValidation.TransactOpts, imageBase64, hash, externalURL, minHeight, maxHeight)
}

// Mint is a paid mutator transaction binding the contract method 0x89ca53a8.
//
// Solidity: function mint(string imageBase64, string hash, string externalURL, uint256 minHeight, uint256 maxHeight) returns()
func (_LotteryValidation *LotteryValidationTransactorSession) Mint(imageBase64 string, hash string, externalURL string, minHeight *big.Int, maxHeight *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.Contract.Mint(&_LotteryValidation.TransactOpts, imageBase64, hash, externalURL, minHeight, maxHeight)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryValidation *LotteryValidationTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryValidation *LotteryValidationSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.Contract.SafeTransferFrom(&_LotteryValidation.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryValidation *LotteryValidationTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.Contract.SafeTransferFrom(&_LotteryValidation.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LotteryValidation *LotteryValidationTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LotteryValidation.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LotteryValidation *LotteryValidationSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LotteryValidation.Contract.SafeTransferFrom0(&_LotteryValidation.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LotteryValidation *LotteryValidationTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LotteryValidation.Contract.SafeTransferFrom0(&_LotteryValidation.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LotteryValidation *LotteryValidationTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LotteryValidation.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LotteryValidation *LotteryValidationSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LotteryValidation.Contract.SetApprovalForAll(&_LotteryValidation.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LotteryValidation *LotteryValidationTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LotteryValidation.Contract.SetApprovalForAll(&_LotteryValidation.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryValidation *LotteryValidationTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryValidation *LotteryValidationSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.Contract.TransferFrom(&_LotteryValidation.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LotteryValidation *LotteryValidationTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LotteryValidation.Contract.TransferFrom(&_LotteryValidation.TransactOpts, from, to, tokenId)
}

// LotteryValidationApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LotteryValidation contract.
type LotteryValidationApprovalIterator struct {
	Event *LotteryValidationApproval // Event containing the contract specifics and raw log

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
func (it *LotteryValidationApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryValidationApproval)
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
		it.Event = new(LotteryValidationApproval)
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
func (it *LotteryValidationApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryValidationApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryValidationApproval represents a Approval event raised by the LotteryValidation contract.
type LotteryValidationApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LotteryValidation *LotteryValidationFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LotteryValidationApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryValidation.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LotteryValidationApprovalIterator{contract: _LotteryValidation.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LotteryValidation *LotteryValidationFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LotteryValidationApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryValidation.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryValidationApproval)
				if err := _LotteryValidation.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LotteryValidation *LotteryValidationFilterer) ParseApproval(log types.Log) (*LotteryValidationApproval, error) {
	event := new(LotteryValidationApproval)
	if err := _LotteryValidation.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryValidationApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LotteryValidation contract.
type LotteryValidationApprovalForAllIterator struct {
	Event *LotteryValidationApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LotteryValidationApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryValidationApprovalForAll)
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
		it.Event = new(LotteryValidationApprovalForAll)
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
func (it *LotteryValidationApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryValidationApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryValidationApprovalForAll represents a ApprovalForAll event raised by the LotteryValidation contract.
type LotteryValidationApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LotteryValidation *LotteryValidationFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LotteryValidationApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LotteryValidation.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LotteryValidationApprovalForAllIterator{contract: _LotteryValidation.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LotteryValidation *LotteryValidationFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LotteryValidationApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LotteryValidation.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryValidationApprovalForAll)
				if err := _LotteryValidation.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LotteryValidation *LotteryValidationFilterer) ParseApprovalForAll(log types.Log) (*LotteryValidationApprovalForAll, error) {
	event := new(LotteryValidationApprovalForAll)
	if err := _LotteryValidation.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryValidationTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LotteryValidation contract.
type LotteryValidationTransferIterator struct {
	Event *LotteryValidationTransfer // Event containing the contract specifics and raw log

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
func (it *LotteryValidationTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryValidationTransfer)
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
		it.Event = new(LotteryValidationTransfer)
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
func (it *LotteryValidationTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryValidationTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryValidationTransfer represents a Transfer event raised by the LotteryValidation contract.
type LotteryValidationTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LotteryValidation *LotteryValidationFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LotteryValidationTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryValidation.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LotteryValidationTransferIterator{contract: _LotteryValidation.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LotteryValidation *LotteryValidationFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LotteryValidationTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LotteryValidation.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryValidationTransfer)
				if err := _LotteryValidation.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LotteryValidation *LotteryValidationFilterer) ParseTransfer(log types.Log) (*LotteryValidationTransfer, error) {
	event := new(LotteryValidationTransfer)
	if err := _LotteryValidation.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
