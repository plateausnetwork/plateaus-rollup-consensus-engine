// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Addr    common.Address
	Network string
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Id            *big.Int
	Subscriptions []Struct0
	Winners       []Struct0
	StartDate     *big.Int
	FinishDate    *big.Int
}

// LotteryMetaData contains all meta data concerning the Lottery contract.
var LotteryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"payable\":false,\"inputs\":[]},{\"type\":\"function\",\"name\":\"_validateNetwork\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"network\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"addNetwork\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"network\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"currentLottery\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"id\"},{\"type\":\"uint256\",\"name\":\"start_date\"},{\"type\":\"uint256\",\"name\":\"finish_date\"}]},{\"type\":\"function\",\"name\":\"getAllowedNetworks\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"string[]\"}]},{\"type\":\"function\",\"name\":\"getCurrentSub\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"tuple[]\",\"components\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"string\",\"name\":\"network\"}]}]},{\"type\":\"function\",\"name\":\"getCurrentWinners\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"string\"}]},{\"type\":\"function\",\"name\":\"getLotteries\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"tuple[]\",\"components\":[{\"type\":\"uint256\",\"name\":\"id\"},{\"type\":\"tuple[]\",\"name\":\"subscriptions\",\"components\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"string\",\"name\":\"network\"}]},{\"type\":\"tuple[]\",\"name\":\"winners\",\"components\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"string\",\"name\":\"network\"}]},{\"type\":\"uint256\",\"name\":\"start_date\"},{\"type\":\"uint256\",\"name\":\"finish_date\"}]}]},{\"type\":\"function\",\"name\":\"isClosed\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"isOpen\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"lastLottery\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"id\"},{\"type\":\"uint256\",\"name\":\"start_date\"},{\"type\":\"uint256\",\"name\":\"finish_date\"}]},{\"type\":\"function\",\"name\":\"lotteries\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"uint256\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"id\"},{\"type\":\"uint256\",\"name\":\"start_date\"},{\"type\":\"uint256\",\"name\":\"finish_date\"}]},{\"type\":\"function\",\"name\":\"removeNetwork\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"network\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"subscribe\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"string[]\",\"name\":\"networks\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"wasPicked\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"winner\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[]}]",
}

// LotteryABI is the input ABI used to generate the binding from.
// Deprecated: Use LotteryMetaData.ABI instead.
var LotteryABI = LotteryMetaData.ABI

// Lottery is an auto generated Go binding around an Ethereum contract.
type Lottery struct {
	LotteryCaller     // Read-only binding to the contract
	LotteryTransactor // Write-only binding to the contract
	LotteryFilterer   // Log filterer for contract events
}

// LotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LotterySession struct {
	Contract     *Lottery          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LotteryCallerSession struct {
	Contract *LotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LotteryTransactorSession struct {
	Contract     *LotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryRaw struct {
	Contract *Lottery // Generic contract binding to access the raw methods on
}

// LotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryCallerRaw struct {
	Contract *LotteryCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LotteryTransactorRaw struct {
	Contract *LotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLottery creates a new instance of Lottery, bound to a specific deployed contract.
func NewLottery(address common.Address, backend bind.ContractBackend) (*Lottery, error) {
	contract, err := bindLottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lottery{LotteryCaller: LotteryCaller{contract: contract}, LotteryTransactor: LotteryTransactor{contract: contract}, LotteryFilterer: LotteryFilterer{contract: contract}}, nil
}

// NewLotteryCaller creates a new read-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryCaller(address common.Address, caller bind.ContractCaller) (*LotteryCaller, error) {
	contract, err := bindLottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryCaller{contract: contract}, nil
}

// NewLotteryTransactor creates a new write-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryTransactor, error) {
	contract, err := bindLottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryTransactor{contract: contract}, nil
}

// NewLotteryFilterer creates a new log filterer instance of Lottery, bound to a specific deployed contract.
func NewLotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryFilterer, error) {
	contract, err := bindLottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryFilterer{contract: contract}, nil
}

// bindLottery binds a generic wrapper to an already deployed contract.
func bindLottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LotteryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.LotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transact(opts, method, params...)
}

// ValidateNetwork is a free data retrieval call binding the contract method 0xaa21f52c.
//
// Solidity: function _validateNetwork(string network) view returns(bool)
func (_Lottery *LotteryCaller) ValidateNetwork(opts *bind.CallOpts, network string) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "_validateNetwork", network)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateNetwork is a free data retrieval call binding the contract method 0xaa21f52c.
//
// Solidity: function _validateNetwork(string network) view returns(bool)
func (_Lottery *LotterySession) ValidateNetwork(network string) (bool, error) {
	return _Lottery.Contract.ValidateNetwork(&_Lottery.CallOpts, network)
}

// ValidateNetwork is a free data retrieval call binding the contract method 0xaa21f52c.
//
// Solidity: function _validateNetwork(string network) view returns(bool)
func (_Lottery *LotteryCallerSession) ValidateNetwork(network string) (bool, error) {
	return _Lottery.Contract.ValidateNetwork(&_Lottery.CallOpts, network)
}

// CurrentLottery is a free data retrieval call binding the contract method 0x2bd56b06.
//
// Solidity: function currentLottery() view returns(uint256 id, uint256 start_date, uint256 finish_date)
func (_Lottery *LotteryCaller) CurrentLottery(opts *bind.CallOpts) (struct {
	Id         *big.Int
	StartDate  *big.Int
	FinishDate *big.Int
}, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "currentLottery")

	outstruct := new(struct {
		Id         *big.Int
		StartDate  *big.Int
		FinishDate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FinishDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentLottery is a free data retrieval call binding the contract method 0x2bd56b06.
//
// Solidity: function currentLottery() view returns(uint256 id, uint256 start_date, uint256 finish_date)
func (_Lottery *LotterySession) CurrentLottery() (struct {
	Id         *big.Int
	StartDate  *big.Int
	FinishDate *big.Int
}, error) {
	return _Lottery.Contract.CurrentLottery(&_Lottery.CallOpts)
}

// CurrentLottery is a free data retrieval call binding the contract method 0x2bd56b06.
//
// Solidity: function currentLottery() view returns(uint256 id, uint256 start_date, uint256 finish_date)
func (_Lottery *LotteryCallerSession) CurrentLottery() (struct {
	Id         *big.Int
	StartDate  *big.Int
	FinishDate *big.Int
}, error) {
	return _Lottery.Contract.CurrentLottery(&_Lottery.CallOpts)
}

// GetAllowedNetworks is a free data retrieval call binding the contract method 0x6d6217e3.
//
// Solidity: function getAllowedNetworks() view returns(string[])
func (_Lottery *LotteryCaller) GetAllowedNetworks(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getAllowedNetworks")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllowedNetworks is a free data retrieval call binding the contract method 0x6d6217e3.
//
// Solidity: function getAllowedNetworks() view returns(string[])
func (_Lottery *LotterySession) GetAllowedNetworks() ([]string, error) {
	return _Lottery.Contract.GetAllowedNetworks(&_Lottery.CallOpts)
}

// GetAllowedNetworks is a free data retrieval call binding the contract method 0x6d6217e3.
//
// Solidity: function getAllowedNetworks() view returns(string[])
func (_Lottery *LotteryCallerSession) GetAllowedNetworks() ([]string, error) {
	return _Lottery.Contract.GetAllowedNetworks(&_Lottery.CallOpts)
}

// GetCurrentSub is a free data retrieval call binding the contract method 0xaa180910.
//
// Solidity: function getCurrentSub() view returns((address,string)[])
func (_Lottery *LotteryCaller) GetCurrentSub(opts *bind.CallOpts) ([]Struct0, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getCurrentSub")

	if err != nil {
		return *new([]Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct0)).(*[]Struct0)

	return out0, err

}

// GetCurrentSub is a free data retrieval call binding the contract method 0xaa180910.
//
// Solidity: function getCurrentSub() view returns((address,string)[])
func (_Lottery *LotterySession) GetCurrentSub() ([]Struct0, error) {
	return _Lottery.Contract.GetCurrentSub(&_Lottery.CallOpts)
}

// GetCurrentSub is a free data retrieval call binding the contract method 0xaa180910.
//
// Solidity: function getCurrentSub() view returns((address,string)[])
func (_Lottery *LotteryCallerSession) GetCurrentSub() ([]Struct0, error) {
	return _Lottery.Contract.GetCurrentSub(&_Lottery.CallOpts)
}

// GetCurrentWinners is a free data retrieval call binding the contract method 0xa3bc9336.
//
// Solidity: function getCurrentWinners() view returns(string)
func (_Lottery *LotteryCaller) GetCurrentWinners(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getCurrentWinners")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCurrentWinners is a free data retrieval call binding the contract method 0xa3bc9336.
//
// Solidity: function getCurrentWinners() view returns(string)
func (_Lottery *LotterySession) GetCurrentWinners() (string, error) {
	return _Lottery.Contract.GetCurrentWinners(&_Lottery.CallOpts)
}

// GetCurrentWinners is a free data retrieval call binding the contract method 0xa3bc9336.
//
// Solidity: function getCurrentWinners() view returns(string)
func (_Lottery *LotteryCallerSession) GetCurrentWinners() (string, error) {
	return _Lottery.Contract.GetCurrentWinners(&_Lottery.CallOpts)
}

// GetLotteries is a free data retrieval call binding the contract method 0xbcc82ec4.
//
// Solidity: function getLotteries() view returns((uint256,(address,string)[],(address,string)[],uint256,uint256)[])
func (_Lottery *LotteryCaller) GetLotteries(opts *bind.CallOpts) ([]Struct1, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getLotteries")

	if err != nil {
		return *new([]Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct1)).(*[]Struct1)

	return out0, err

}

// GetLotteries is a free data retrieval call binding the contract method 0xbcc82ec4.
//
// Solidity: function getLotteries() view returns((uint256,(address,string)[],(address,string)[],uint256,uint256)[])
func (_Lottery *LotterySession) GetLotteries() ([]Struct1, error) {
	return _Lottery.Contract.GetLotteries(&_Lottery.CallOpts)
}

// GetLotteries is a free data retrieval call binding the contract method 0xbcc82ec4.
//
// Solidity: function getLotteries() view returns((uint256,(address,string)[],(address,string)[],uint256,uint256)[])
func (_Lottery *LotteryCallerSession) GetLotteries() ([]Struct1, error) {
	return _Lottery.Contract.GetLotteries(&_Lottery.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool)
func (_Lottery *LotteryCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "isClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool)
func (_Lottery *LotterySession) IsClosed() (bool, error) {
	return _Lottery.Contract.IsClosed(&_Lottery.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool)
func (_Lottery *LotteryCallerSession) IsClosed() (bool, error) {
	return _Lottery.Contract.IsClosed(&_Lottery.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_Lottery *LotteryCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_Lottery *LotterySession) IsOpen() (bool, error) {
	return _Lottery.Contract.IsOpen(&_Lottery.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_Lottery *LotteryCallerSession) IsOpen() (bool, error) {
	return _Lottery.Contract.IsOpen(&_Lottery.CallOpts)
}

// LastLottery is a free data retrieval call binding the contract method 0xaf0b4b27.
//
// Solidity: function lastLottery() view returns(uint256 id, uint256 start_date, uint256 finish_date)
func (_Lottery *LotteryCaller) LastLottery(opts *bind.CallOpts) (struct {
	Id         *big.Int
	StartDate  *big.Int
	FinishDate *big.Int
}, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "lastLottery")

	outstruct := new(struct {
		Id         *big.Int
		StartDate  *big.Int
		FinishDate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FinishDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LastLottery is a free data retrieval call binding the contract method 0xaf0b4b27.
//
// Solidity: function lastLottery() view returns(uint256 id, uint256 start_date, uint256 finish_date)
func (_Lottery *LotterySession) LastLottery() (struct {
	Id         *big.Int
	StartDate  *big.Int
	FinishDate *big.Int
}, error) {
	return _Lottery.Contract.LastLottery(&_Lottery.CallOpts)
}

// LastLottery is a free data retrieval call binding the contract method 0xaf0b4b27.
//
// Solidity: function lastLottery() view returns(uint256 id, uint256 start_date, uint256 finish_date)
func (_Lottery *LotteryCallerSession) LastLottery() (struct {
	Id         *big.Int
	StartDate  *big.Int
	FinishDate *big.Int
}, error) {
	return _Lottery.Contract.LastLottery(&_Lottery.CallOpts)
}

// Lotteries is a free data retrieval call binding the contract method 0x1398e076.
//
// Solidity: function lotteries(uint256 ) view returns(uint256 id, uint256 start_date, uint256 finish_date)
func (_Lottery *LotteryCaller) Lotteries(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id         *big.Int
	StartDate  *big.Int
	FinishDate *big.Int
}, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "lotteries", arg0)

	outstruct := new(struct {
		Id         *big.Int
		StartDate  *big.Int
		FinishDate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FinishDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Lotteries is a free data retrieval call binding the contract method 0x1398e076.
//
// Solidity: function lotteries(uint256 ) view returns(uint256 id, uint256 start_date, uint256 finish_date)
func (_Lottery *LotterySession) Lotteries(arg0 *big.Int) (struct {
	Id         *big.Int
	StartDate  *big.Int
	FinishDate *big.Int
}, error) {
	return _Lottery.Contract.Lotteries(&_Lottery.CallOpts, arg0)
}

// Lotteries is a free data retrieval call binding the contract method 0x1398e076.
//
// Solidity: function lotteries(uint256 ) view returns(uint256 id, uint256 start_date, uint256 finish_date)
func (_Lottery *LotteryCallerSession) Lotteries(arg0 *big.Int) (struct {
	Id         *big.Int
	StartDate  *big.Int
	FinishDate *big.Int
}, error) {
	return _Lottery.Contract.Lotteries(&_Lottery.CallOpts, arg0)
}

// WasPicked is a free data retrieval call binding the contract method 0xd78afb7c.
//
// Solidity: function wasPicked() view returns(bool)
func (_Lottery *LotteryCaller) WasPicked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "wasPicked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasPicked is a free data retrieval call binding the contract method 0xd78afb7c.
//
// Solidity: function wasPicked() view returns(bool)
func (_Lottery *LotterySession) WasPicked() (bool, error) {
	return _Lottery.Contract.WasPicked(&_Lottery.CallOpts)
}

// WasPicked is a free data retrieval call binding the contract method 0xd78afb7c.
//
// Solidity: function wasPicked() view returns(bool)
func (_Lottery *LotteryCallerSession) WasPicked() (bool, error) {
	return _Lottery.Contract.WasPicked(&_Lottery.CallOpts)
}

// AddNetwork is a paid mutator transaction binding the contract method 0x91ef7357.
//
// Solidity: function addNetwork(string network) returns()
func (_Lottery *LotteryTransactor) AddNetwork(opts *bind.TransactOpts, network string) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "addNetwork", network)
}

// AddNetwork is a paid mutator transaction binding the contract method 0x91ef7357.
//
// Solidity: function addNetwork(string network) returns()
func (_Lottery *LotterySession) AddNetwork(network string) (*types.Transaction, error) {
	return _Lottery.Contract.AddNetwork(&_Lottery.TransactOpts, network)
}

// AddNetwork is a paid mutator transaction binding the contract method 0x91ef7357.
//
// Solidity: function addNetwork(string network) returns()
func (_Lottery *LotteryTransactorSession) AddNetwork(network string) (*types.Transaction, error) {
	return _Lottery.Contract.AddNetwork(&_Lottery.TransactOpts, network)
}

// RemoveNetwork is a paid mutator transaction binding the contract method 0x3d30bb4e.
//
// Solidity: function removeNetwork(string network) returns()
func (_Lottery *LotteryTransactor) RemoveNetwork(opts *bind.TransactOpts, network string) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "removeNetwork", network)
}

// RemoveNetwork is a paid mutator transaction binding the contract method 0x3d30bb4e.
//
// Solidity: function removeNetwork(string network) returns()
func (_Lottery *LotterySession) RemoveNetwork(network string) (*types.Transaction, error) {
	return _Lottery.Contract.RemoveNetwork(&_Lottery.TransactOpts, network)
}

// RemoveNetwork is a paid mutator transaction binding the contract method 0x3d30bb4e.
//
// Solidity: function removeNetwork(string network) returns()
func (_Lottery *LotteryTransactorSession) RemoveNetwork(network string) (*types.Transaction, error) {
	return _Lottery.Contract.RemoveNetwork(&_Lottery.TransactOpts, network)
}

// Subscribe is a paid mutator transaction binding the contract method 0x576adb3c.
//
// Solidity: function subscribe(string[] networks) returns()
func (_Lottery *LotteryTransactor) Subscribe(opts *bind.TransactOpts, networks []string) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "subscribe", networks)
}

// Subscribe is a paid mutator transaction binding the contract method 0x576adb3c.
//
// Solidity: function subscribe(string[] networks) returns()
func (_Lottery *LotterySession) Subscribe(networks []string) (*types.Transaction, error) {
	return _Lottery.Contract.Subscribe(&_Lottery.TransactOpts, networks)
}

// Subscribe is a paid mutator transaction binding the contract method 0x576adb3c.
//
// Solidity: function subscribe(string[] networks) returns()
func (_Lottery *LotteryTransactorSession) Subscribe(networks []string) (*types.Transaction, error) {
	return _Lottery.Contract.Subscribe(&_Lottery.TransactOpts, networks)
}

// Winner is a paid mutator transaction binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() returns()
func (_Lottery *LotteryTransactor) Winner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "winner")
}

// Winner is a paid mutator transaction binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() returns()
func (_Lottery *LotterySession) Winner() (*types.Transaction, error) {
	return _Lottery.Contract.Winner(&_Lottery.TransactOpts)
}

// Winner is a paid mutator transaction binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() returns()
func (_Lottery *LotteryTransactorSession) Winner() (*types.Transaction, error) {
	return _Lottery.Contract.Winner(&_Lottery.TransactOpts)
}
