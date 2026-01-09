// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swap

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

// SwapMetaData contains all meta data concerning the Swap contract.
var SwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_exchangeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"ExchangeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"setExchangeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"swapAforB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"swapBforA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokensSwapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getAmountOutAforB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getAmountOutBforA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenA\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenB\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SwapABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapMetaData.ABI instead.
var SwapABI = SwapMetaData.ABI

// Swap is an auto generated Go binding around an Ethereum contract.
type Swap struct {
	SwapCaller     // Read-only binding to the contract
	SwapTransactor // Write-only binding to the contract
	SwapFilterer   // Log filterer for contract events
}

// SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapSession struct {
	Contract     *Swap             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapCallerSession struct {
	Contract *SwapCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTransactorSession struct {
	Contract     *SwapTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRaw struct {
	Contract *Swap // Generic contract binding to access the raw methods on
}

// SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapCallerRaw struct {
	Contract *SwapCaller // Generic read-only contract binding to access the raw methods on
}

// SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTransactorRaw struct {
	Contract *SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwap creates a new instance of Swap, bound to a specific deployed contract.
func NewSwap(address common.Address, backend bind.ContractBackend) (*Swap, error) {
	contract, err := bindSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// NewSwapCaller creates a new read-only instance of Swap, bound to a specific deployed contract.
func NewSwapCaller(address common.Address, caller bind.ContractCaller) (*SwapCaller, error) {
	contract, err := bindSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapCaller{contract: contract}, nil
}

// NewSwapTransactor creates a new write-only instance of Swap, bound to a specific deployed contract.
func NewSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapTransactor, error) {
	contract, err := bindSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTransactor{contract: contract}, nil
}

// NewSwapFilterer creates a new log filterer instance of Swap, bound to a specific deployed contract.
func NewSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapFilterer, error) {
	contract, err := bindSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapFilterer{contract: contract}, nil
}

// bindSwap binds a generic wrapper to an already deployed contract.
func bindSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transact(opts, method, params...)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Swap *SwapCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "exchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Swap *SwapSession) ExchangeRate() (*big.Int, error) {
	return _Swap.Contract.ExchangeRate(&_Swap.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Swap *SwapCallerSession) ExchangeRate() (*big.Int, error) {
	return _Swap.Contract.ExchangeRate(&_Swap.CallOpts)
}

// GetAmountOutAforB is a free data retrieval call binding the contract method 0x234e85c6.
//
// Solidity: function getAmountOutAforB(uint256 amountIn) view returns(uint256)
func (_Swap *SwapCaller) GetAmountOutAforB(opts *bind.CallOpts, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "getAmountOutAforB", amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOutAforB is a free data retrieval call binding the contract method 0x234e85c6.
//
// Solidity: function getAmountOutAforB(uint256 amountIn) view returns(uint256)
func (_Swap *SwapSession) GetAmountOutAforB(amountIn *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetAmountOutAforB(&_Swap.CallOpts, amountIn)
}

// GetAmountOutAforB is a free data retrieval call binding the contract method 0x234e85c6.
//
// Solidity: function getAmountOutAforB(uint256 amountIn) view returns(uint256)
func (_Swap *SwapCallerSession) GetAmountOutAforB(amountIn *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetAmountOutAforB(&_Swap.CallOpts, amountIn)
}

// GetAmountOutBforA is a free data retrieval call binding the contract method 0x0e6a22b0.
//
// Solidity: function getAmountOutBforA(uint256 amountIn) view returns(uint256)
func (_Swap *SwapCaller) GetAmountOutBforA(opts *bind.CallOpts, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "getAmountOutBforA", amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOutBforA is a free data retrieval call binding the contract method 0x0e6a22b0.
//
// Solidity: function getAmountOutBforA(uint256 amountIn) view returns(uint256)
func (_Swap *SwapSession) GetAmountOutBforA(amountIn *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetAmountOutBforA(&_Swap.CallOpts, amountIn)
}

// GetAmountOutBforA is a free data retrieval call binding the contract method 0x0e6a22b0.
//
// Solidity: function getAmountOutBforA(uint256 amountIn) view returns(uint256)
func (_Swap *SwapCallerSession) GetAmountOutBforA(amountIn *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetAmountOutBforA(&_Swap.CallOpts, amountIn)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 reserveA, uint256 reserveB)
func (_Swap *SwapCaller) GetReserves(opts *bind.CallOpts) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		ReserveA *big.Int
		ReserveB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 reserveA, uint256 reserveB)
func (_Swap *SwapSession) GetReserves() (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Swap.Contract.GetReserves(&_Swap.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 reserveA, uint256 reserveB)
func (_Swap *SwapCallerSession) GetReserves() (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _Swap.Contract.GetReserves(&_Swap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapSession) Owner() (common.Address, error) {
	return _Swap.Contract.Owner(&_Swap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapCallerSession) Owner() (common.Address, error) {
	return _Swap.Contract.Owner(&_Swap.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Swap *SwapCaller) TokenA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "tokenA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Swap *SwapSession) TokenA() (common.Address, error) {
	return _Swap.Contract.TokenA(&_Swap.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Swap *SwapCallerSession) TokenA() (common.Address, error) {
	return _Swap.Contract.TokenA(&_Swap.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Swap *SwapCaller) TokenB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "tokenB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Swap *SwapSession) TokenB() (common.Address, error) {
	return _Swap.Contract.TokenB(&_Swap.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Swap *SwapCallerSession) TokenB() (common.Address, error) {
	return _Swap.Contract.TokenB(&_Swap.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address token, uint256 amount) returns()
func (_Swap *SwapTransactor) AddLiquidity(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "addLiquidity", token, amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address token, uint256 amount) returns()
func (_Swap *SwapSession) AddLiquidity(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AddLiquidity(&_Swap.TransactOpts, token, amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address token, uint256 amount) returns()
func (_Swap *SwapTransactorSession) AddLiquidity(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AddLiquidity(&_Swap.TransactOpts, token, amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address token, uint256 amount) returns()
func (_Swap *SwapTransactor) RemoveLiquidity(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "removeLiquidity", token, amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address token, uint256 amount) returns()
func (_Swap *SwapSession) RemoveLiquidity(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.RemoveLiquidity(&_Swap.TransactOpts, token, amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address token, uint256 amount) returns()
func (_Swap *SwapTransactorSession) RemoveLiquidity(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.RemoveLiquidity(&_Swap.TransactOpts, token, amount)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 newRate) returns()
func (_Swap *SwapTransactor) SetExchangeRate(opts *bind.TransactOpts, newRate *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "setExchangeRate", newRate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 newRate) returns()
func (_Swap *SwapSession) SetExchangeRate(newRate *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SetExchangeRate(&_Swap.TransactOpts, newRate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 newRate) returns()
func (_Swap *SwapTransactorSession) SetExchangeRate(newRate *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SetExchangeRate(&_Swap.TransactOpts, newRate)
}

// SwapAforB is a paid mutator transaction binding the contract method 0xe4f1f70a.
//
// Solidity: function swapAforB(uint256 amountIn) returns()
func (_Swap *SwapTransactor) SwapAforB(opts *bind.TransactOpts, amountIn *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "swapAforB", amountIn)
}

// SwapAforB is a paid mutator transaction binding the contract method 0xe4f1f70a.
//
// Solidity: function swapAforB(uint256 amountIn) returns()
func (_Swap *SwapSession) SwapAforB(amountIn *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SwapAforB(&_Swap.TransactOpts, amountIn)
}

// SwapAforB is a paid mutator transaction binding the contract method 0xe4f1f70a.
//
// Solidity: function swapAforB(uint256 amountIn) returns()
func (_Swap *SwapTransactorSession) SwapAforB(amountIn *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SwapAforB(&_Swap.TransactOpts, amountIn)
}

// SwapBforA is a paid mutator transaction binding the contract method 0x8014a7aa.
//
// Solidity: function swapBforA(uint256 amountIn) returns()
func (_Swap *SwapTransactor) SwapBforA(opts *bind.TransactOpts, amountIn *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "swapBforA", amountIn)
}

// SwapBforA is a paid mutator transaction binding the contract method 0x8014a7aa.
//
// Solidity: function swapBforA(uint256 amountIn) returns()
func (_Swap *SwapSession) SwapBforA(amountIn *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SwapBforA(&_Swap.TransactOpts, amountIn)
}

// SwapBforA is a paid mutator transaction binding the contract method 0x8014a7aa.
//
// Solidity: function swapBforA(uint256 amountIn) returns()
func (_Swap *SwapTransactorSession) SwapBforA(amountIn *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SwapBforA(&_Swap.TransactOpts, amountIn)
}

// SwapExchangeRateUpdatedIterator is returned from FilterExchangeRateUpdated and is used to iterate over the raw logs and unpacked data for ExchangeRateUpdated events raised by the Swap contract.
type SwapExchangeRateUpdatedIterator struct {
	Event *SwapExchangeRateUpdated // Event containing the contract specifics and raw log

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
func (it *SwapExchangeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapExchangeRateUpdated)
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
		it.Event = new(SwapExchangeRateUpdated)
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
func (it *SwapExchangeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapExchangeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapExchangeRateUpdated represents a ExchangeRateUpdated event raised by the Swap contract.
type SwapExchangeRateUpdated struct {
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateUpdated is a free log retrieval operation binding the contract event 0x388f446e9526fe5c9af20a5919b342370c8a7c0cb05245afe1e545658fa3cdba.
//
// Solidity: event ExchangeRateUpdated(uint256 newRate)
func (_Swap *SwapFilterer) FilterExchangeRateUpdated(opts *bind.FilterOpts) (*SwapExchangeRateUpdatedIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "ExchangeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &SwapExchangeRateUpdatedIterator{contract: _Swap.contract, event: "ExchangeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchExchangeRateUpdated is a free log subscription operation binding the contract event 0x388f446e9526fe5c9af20a5919b342370c8a7c0cb05245afe1e545658fa3cdba.
//
// Solidity: event ExchangeRateUpdated(uint256 newRate)
func (_Swap *SwapFilterer) WatchExchangeRateUpdated(opts *bind.WatchOpts, sink chan<- *SwapExchangeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "ExchangeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapExchangeRateUpdated)
				if err := _Swap.contract.UnpackLog(event, "ExchangeRateUpdated", log); err != nil {
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

// ParseExchangeRateUpdated is a log parse operation binding the contract event 0x388f446e9526fe5c9af20a5919b342370c8a7c0cb05245afe1e545658fa3cdba.
//
// Solidity: event ExchangeRateUpdated(uint256 newRate)
func (_Swap *SwapFilterer) ParseExchangeRateUpdated(log types.Log) (*SwapExchangeRateUpdated, error) {
	event := new(SwapExchangeRateUpdated)
	if err := _Swap.contract.UnpackLog(event, "ExchangeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the Swap contract.
type SwapLiquidityAddedIterator struct {
	Event *SwapLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *SwapLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapLiquidityAdded)
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
		it.Event = new(SwapLiquidityAdded)
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
func (it *SwapLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapLiquidityAdded represents a LiquidityAdded event raised by the Swap contract.
type SwapLiquidityAdded struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0xc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088.
//
// Solidity: event LiquidityAdded(address token, uint256 amount)
func (_Swap *SwapFilterer) FilterLiquidityAdded(opts *bind.FilterOpts) (*SwapLiquidityAddedIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return &SwapLiquidityAddedIterator{contract: _Swap.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0xc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088.
//
// Solidity: event LiquidityAdded(address token, uint256 amount)
func (_Swap *SwapFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *SwapLiquidityAdded) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapLiquidityAdded)
				if err := _Swap.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0xc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088.
//
// Solidity: event LiquidityAdded(address token, uint256 amount)
func (_Swap *SwapFilterer) ParseLiquidityAdded(log types.Log) (*SwapLiquidityAdded, error) {
	event := new(SwapLiquidityAdded)
	if err := _Swap.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapLiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the Swap contract.
type SwapLiquidityRemovedIterator struct {
	Event *SwapLiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *SwapLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapLiquidityRemoved)
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
		it.Event = new(SwapLiquidityRemoved)
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
func (it *SwapLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapLiquidityRemoved represents a LiquidityRemoved event raised by the Swap contract.
type SwapLiquidityRemoved struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemoved is a free log retrieval operation binding the contract event 0xc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719.
//
// Solidity: event LiquidityRemoved(address token, uint256 amount)
func (_Swap *SwapFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts) (*SwapLiquidityRemovedIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "LiquidityRemoved")
	if err != nil {
		return nil, err
	}
	return &SwapLiquidityRemovedIterator{contract: _Swap.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0xc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719.
//
// Solidity: event LiquidityRemoved(address token, uint256 amount)
func (_Swap *SwapFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *SwapLiquidityRemoved) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "LiquidityRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapLiquidityRemoved)
				if err := _Swap.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

// ParseLiquidityRemoved is a log parse operation binding the contract event 0xc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719.
//
// Solidity: event LiquidityRemoved(address token, uint256 amount)
func (_Swap *SwapFilterer) ParseLiquidityRemoved(log types.Log) (*SwapLiquidityRemoved, error) {
	event := new(SwapLiquidityRemoved)
	if err := _Swap.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapTokensSwappedIterator is returned from FilterTokensSwapped and is used to iterate over the raw logs and unpacked data for TokensSwapped events raised by the Swap contract.
type SwapTokensSwappedIterator struct {
	Event *SwapTokensSwapped // Event containing the contract specifics and raw log

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
func (it *SwapTokensSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapTokensSwapped)
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
		it.Event = new(SwapTokensSwapped)
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
func (it *SwapTokensSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapTokensSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapTokensSwapped represents a TokensSwapped event raised by the Swap contract.
type SwapTokensSwapped struct {
	User      common.Address
	FromToken common.Address
	ToToken   common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensSwapped is a free log retrieval operation binding the contract event 0xad56699d0f375866eb895ed27203058a36a713382aaded78eb6b67da266d4332.
//
// Solidity: event TokensSwapped(address indexed user, address indexed fromToken, address indexed toToken, uint256 amountIn, uint256 amountOut)
func (_Swap *SwapFilterer) FilterTokensSwapped(opts *bind.FilterOpts, user []common.Address, fromToken []common.Address, toToken []common.Address) (*SwapTokensSwappedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "TokensSwapped", userRule, fromTokenRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return &SwapTokensSwappedIterator{contract: _Swap.contract, event: "TokensSwapped", logs: logs, sub: sub}, nil
}

// WatchTokensSwapped is a free log subscription operation binding the contract event 0xad56699d0f375866eb895ed27203058a36a713382aaded78eb6b67da266d4332.
//
// Solidity: event TokensSwapped(address indexed user, address indexed fromToken, address indexed toToken, uint256 amountIn, uint256 amountOut)
func (_Swap *SwapFilterer) WatchTokensSwapped(opts *bind.WatchOpts, sink chan<- *SwapTokensSwapped, user []common.Address, fromToken []common.Address, toToken []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "TokensSwapped", userRule, fromTokenRule, toTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapTokensSwapped)
				if err := _Swap.contract.UnpackLog(event, "TokensSwapped", log); err != nil {
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

// ParseTokensSwapped is a log parse operation binding the contract event 0xad56699d0f375866eb895ed27203058a36a713382aaded78eb6b67da266d4332.
//
// Solidity: event TokensSwapped(address indexed user, address indexed fromToken, address indexed toToken, uint256 amountIn, uint256 amountOut)
func (_Swap *SwapFilterer) ParseTokensSwapped(log types.Log) (*SwapTokensSwapped, error) {
	event := new(SwapTokensSwapped)
	if err := _Swap.contract.UnpackLog(event, "TokensSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
