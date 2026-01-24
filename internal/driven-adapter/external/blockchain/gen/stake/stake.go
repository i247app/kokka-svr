// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake

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

// StakeMetaData contains all meta data concerning the Stake contract.
var StakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"apy\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAllRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAPY\",\"type\":\"uint256\"}],\"name\":\"APYUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsDeposited\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newAPY\",\"type\":\"uint256\"}],\"name\":\"setAPY\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"apy\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"apyRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAllUserStakes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakedAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pendingRewardAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalStakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPoolAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeStartTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakeABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeMetaData.ABI instead.
var StakeABI = StakeMetaData.ABI

// Stake is an auto generated Go binding around an Ethereum contract.
type Stake struct {
	StakeCaller     // Read-only binding to the contract
	StakeTransactor // Write-only binding to the contract
	StakeFilterer   // Log filterer for contract events
}

// StakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeSession struct {
	Contract     *Stake            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeCallerSession struct {
	Contract *StakeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTransactorSession struct {
	Contract     *StakeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRaw struct {
	Contract *Stake // Generic contract binding to access the raw methods on
}

// StakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeCallerRaw struct {
	Contract *StakeCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTransactorRaw struct {
	Contract *StakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStake creates a new instance of Stake, bound to a specific deployed contract.
func NewStake(address common.Address, backend bind.ContractBackend) (*Stake, error) {
	contract, err := bindStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// NewStakeCaller creates a new read-only instance of Stake, bound to a specific deployed contract.
func NewStakeCaller(address common.Address, caller bind.ContractCaller) (*StakeCaller, error) {
	contract, err := bindStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCaller{contract: contract}, nil
}

// NewStakeTransactor creates a new write-only instance of Stake, bound to a specific deployed contract.
func NewStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTransactor, error) {
	contract, err := bindStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTransactor{contract: contract}, nil
}

// NewStakeFilterer creates a new log filterer instance of Stake, bound to a specific deployed contract.
func NewStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeFilterer, error) {
	contract, err := bindStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeFilterer{contract: contract}, nil
}

// bindStake binds a generic wrapper to an already deployed contract.
func bindStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.StakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Stake *StakeCaller) BASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Stake *StakeSession) BASISPOINTS() (*big.Int, error) {
	return _Stake.Contract.BASISPOINTS(&_Stake.CallOpts)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Stake *StakeCallerSession) BASISPOINTS() (*big.Int, error) {
	return _Stake.Contract.BASISPOINTS(&_Stake.CallOpts)
}

// SECONDSPERYEAR is a free data retrieval call binding the contract method 0xe6a69ab8.
//
// Solidity: function SECONDS_PER_YEAR() view returns(uint256)
func (_Stake *StakeCaller) SECONDSPERYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "SECONDS_PER_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSPERYEAR is a free data retrieval call binding the contract method 0xe6a69ab8.
//
// Solidity: function SECONDS_PER_YEAR() view returns(uint256)
func (_Stake *StakeSession) SECONDSPERYEAR() (*big.Int, error) {
	return _Stake.Contract.SECONDSPERYEAR(&_Stake.CallOpts)
}

// SECONDSPERYEAR is a free data retrieval call binding the contract method 0xe6a69ab8.
//
// Solidity: function SECONDS_PER_YEAR() view returns(uint256)
func (_Stake *StakeCallerSession) SECONDSPERYEAR() (*big.Int, error) {
	return _Stake.Contract.SECONDSPERYEAR(&_Stake.CallOpts)
}

// ApyRates is a free data retrieval call binding the contract method 0x91d315c9.
//
// Solidity: function apyRates(address ) view returns(uint256)
func (_Stake *StakeCaller) ApyRates(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "apyRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApyRates is a free data retrieval call binding the contract method 0x91d315c9.
//
// Solidity: function apyRates(address ) view returns(uint256)
func (_Stake *StakeSession) ApyRates(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.ApyRates(&_Stake.CallOpts, arg0)
}

// ApyRates is a free data retrieval call binding the contract method 0x91d315c9.
//
// Solidity: function apyRates(address ) view returns(uint256)
func (_Stake *StakeCallerSession) ApyRates(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.ApyRates(&_Stake.CallOpts, arg0)
}

// GetAllUserStakes is a free data retrieval call binding the contract method 0x8bd7ca57.
//
// Solidity: function getAllUserStakes(address user) view returns(address[] tokens, uint256[] stakedAmounts, uint256[] pendingRewardAmounts)
func (_Stake *StakeCaller) GetAllUserStakes(opts *bind.CallOpts, user common.Address) (struct {
	Tokens               []common.Address
	StakedAmounts        []*big.Int
	PendingRewardAmounts []*big.Int
}, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "getAllUserStakes", user)

	outstruct := new(struct {
		Tokens               []common.Address
		StakedAmounts        []*big.Int
		PendingRewardAmounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.StakedAmounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.PendingRewardAmounts = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAllUserStakes is a free data retrieval call binding the contract method 0x8bd7ca57.
//
// Solidity: function getAllUserStakes(address user) view returns(address[] tokens, uint256[] stakedAmounts, uint256[] pendingRewardAmounts)
func (_Stake *StakeSession) GetAllUserStakes(user common.Address) (struct {
	Tokens               []common.Address
	StakedAmounts        []*big.Int
	PendingRewardAmounts []*big.Int
}, error) {
	return _Stake.Contract.GetAllUserStakes(&_Stake.CallOpts, user)
}

// GetAllUserStakes is a free data retrieval call binding the contract method 0x8bd7ca57.
//
// Solidity: function getAllUserStakes(address user) view returns(address[] tokens, uint256[] stakedAmounts, uint256[] pendingRewardAmounts)
func (_Stake *StakeCallerSession) GetAllUserStakes(user common.Address) (struct {
	Tokens               []common.Address
	StakedAmounts        []*big.Int
	PendingRewardAmounts []*big.Int
}, error) {
	return _Stake.Contract.GetAllUserStakes(&_Stake.CallOpts, user)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[])
func (_Stake *StakeCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[])
func (_Stake *StakeSession) GetSupportedTokens() ([]common.Address, error) {
	return _Stake.Contract.GetSupportedTokens(&_Stake.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[])
func (_Stake *StakeCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _Stake.Contract.GetSupportedTokens(&_Stake.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns(uint256 totalStakedAmount, uint256 rewardPoolAmount, bool isSupported)
func (_Stake *StakeCaller) GetTokenInfo(opts *bind.CallOpts, token common.Address) (struct {
	TotalStakedAmount *big.Int
	RewardPoolAmount  *big.Int
	IsSupported       bool
}, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "getTokenInfo", token)

	outstruct := new(struct {
		TotalStakedAmount *big.Int
		RewardPoolAmount  *big.Int
		IsSupported       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalStakedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardPoolAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsSupported = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns(uint256 totalStakedAmount, uint256 rewardPoolAmount, bool isSupported)
func (_Stake *StakeSession) GetTokenInfo(token common.Address) (struct {
	TotalStakedAmount *big.Int
	RewardPoolAmount  *big.Int
	IsSupported       bool
}, error) {
	return _Stake.Contract.GetTokenInfo(&_Stake.CallOpts, token)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address token) view returns(uint256 totalStakedAmount, uint256 rewardPoolAmount, bool isSupported)
func (_Stake *StakeCallerSession) GetTokenInfo(token common.Address) (struct {
	TotalStakedAmount *big.Int
	RewardPoolAmount  *big.Int
	IsSupported       bool
}, error) {
	return _Stake.Contract.GetTokenInfo(&_Stake.CallOpts, token)
}

// GetUserStake is a free data retrieval call binding the contract method 0x3710d4c7.
//
// Solidity: function getUserStake(address token, address user) view returns(uint256 stakedAmount, uint256 pendingReward, uint256 stakeStartTime)
func (_Stake *StakeCaller) GetUserStake(opts *bind.CallOpts, token common.Address, user common.Address) (struct {
	StakedAmount   *big.Int
	PendingReward  *big.Int
	StakeStartTime *big.Int
}, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "getUserStake", token, user)

	outstruct := new(struct {
		StakedAmount   *big.Int
		PendingReward  *big.Int
		StakeStartTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PendingReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StakeStartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserStake is a free data retrieval call binding the contract method 0x3710d4c7.
//
// Solidity: function getUserStake(address token, address user) view returns(uint256 stakedAmount, uint256 pendingReward, uint256 stakeStartTime)
func (_Stake *StakeSession) GetUserStake(token common.Address, user common.Address) (struct {
	StakedAmount   *big.Int
	PendingReward  *big.Int
	StakeStartTime *big.Int
}, error) {
	return _Stake.Contract.GetUserStake(&_Stake.CallOpts, token, user)
}

// GetUserStake is a free data retrieval call binding the contract method 0x3710d4c7.
//
// Solidity: function getUserStake(address token, address user) view returns(uint256 stakedAmount, uint256 pendingReward, uint256 stakeStartTime)
func (_Stake *StakeCallerSession) GetUserStake(token common.Address, user common.Address) (struct {
	StakedAmount   *big.Int
	PendingReward  *big.Int
	StakeStartTime *big.Int
}, error) {
	return _Stake.Contract.GetUserStake(&_Stake.CallOpts, token, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCallerSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0x80ac8228.
//
// Solidity: function pendingRewards(address token, address user) view returns(uint256)
func (_Stake *StakeCaller) PendingRewards(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "pendingRewards", token, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0x80ac8228.
//
// Solidity: function pendingRewards(address token, address user) view returns(uint256)
func (_Stake *StakeSession) PendingRewards(token common.Address, user common.Address) (*big.Int, error) {
	return _Stake.Contract.PendingRewards(&_Stake.CallOpts, token, user)
}

// PendingRewards is a free data retrieval call binding the contract method 0x80ac8228.
//
// Solidity: function pendingRewards(address token, address user) view returns(uint256)
func (_Stake *StakeCallerSession) PendingRewards(token common.Address, user common.Address) (*big.Int, error) {
	return _Stake.Contract.PendingRewards(&_Stake.CallOpts, token, user)
}

// RewardPool is a free data retrieval call binding the contract method 0x1bcda198.
//
// Solidity: function rewardPool(address ) view returns(uint256)
func (_Stake *StakeCaller) RewardPool(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "rewardPool", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPool is a free data retrieval call binding the contract method 0x1bcda198.
//
// Solidity: function rewardPool(address ) view returns(uint256)
func (_Stake *StakeSession) RewardPool(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.RewardPool(&_Stake.CallOpts, arg0)
}

// RewardPool is a free data retrieval call binding the contract method 0x1bcda198.
//
// Solidity: function rewardPool(address ) view returns(uint256)
func (_Stake *StakeCallerSession) RewardPool(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.RewardPool(&_Stake.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0xa4e47b66.
//
// Solidity: function stakes(address , address ) view returns(uint256 amount, uint256 startTime, uint256 lastClaimTime)
func (_Stake *StakeCaller) Stakes(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Amount        *big.Int
	StartTime     *big.Int
	LastClaimTime *big.Int
}, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "stakes", arg0, arg1)

	outstruct := new(struct {
		Amount        *big.Int
		StartTime     *big.Int
		LastClaimTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastClaimTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0xa4e47b66.
//
// Solidity: function stakes(address , address ) view returns(uint256 amount, uint256 startTime, uint256 lastClaimTime)
func (_Stake *StakeSession) Stakes(arg0 common.Address, arg1 common.Address) (struct {
	Amount        *big.Int
	StartTime     *big.Int
	LastClaimTime *big.Int
}, error) {
	return _Stake.Contract.Stakes(&_Stake.CallOpts, arg0, arg1)
}

// Stakes is a free data retrieval call binding the contract method 0xa4e47b66.
//
// Solidity: function stakes(address , address ) view returns(uint256 amount, uint256 startTime, uint256 lastClaimTime)
func (_Stake *StakeCallerSession) Stakes(arg0 common.Address, arg1 common.Address) (struct {
	Amount        *big.Int
	StartTime     *big.Int
	LastClaimTime *big.Int
}, error) {
	return _Stake.Contract.Stakes(&_Stake.CallOpts, arg0, arg1)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_Stake *StakeCaller) SupportedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "supportedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_Stake *StakeSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _Stake.Contract.SupportedTokens(&_Stake.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_Stake *StakeCallerSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _Stake.Contract.SupportedTokens(&_Stake.CallOpts, arg0)
}

// TokenList is a free data retrieval call binding the contract method 0x9ead7222.
//
// Solidity: function tokenList(uint256 ) view returns(address)
func (_Stake *StakeCaller) TokenList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "tokenList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenList is a free data retrieval call binding the contract method 0x9ead7222.
//
// Solidity: function tokenList(uint256 ) view returns(address)
func (_Stake *StakeSession) TokenList(arg0 *big.Int) (common.Address, error) {
	return _Stake.Contract.TokenList(&_Stake.CallOpts, arg0)
}

// TokenList is a free data retrieval call binding the contract method 0x9ead7222.
//
// Solidity: function tokenList(uint256 ) view returns(address)
func (_Stake *StakeCallerSession) TokenList(arg0 *big.Int) (common.Address, error) {
	return _Stake.Contract.TokenList(&_Stake.CallOpts, arg0)
}

// TotalStaked is a free data retrieval call binding the contract method 0x9bfd8d61.
//
// Solidity: function totalStaked(address ) view returns(uint256)
func (_Stake *StakeCaller) TotalStaked(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "totalStaked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x9bfd8d61.
//
// Solidity: function totalStaked(address ) view returns(uint256)
func (_Stake *StakeSession) TotalStaked(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.TotalStaked(&_Stake.CallOpts, arg0)
}

// TotalStaked is a free data retrieval call binding the contract method 0x9bfd8d61.
//
// Solidity: function totalStaked(address ) view returns(uint256)
func (_Stake *StakeCallerSession) TotalStaked(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.TotalStaked(&_Stake.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0xaf81c5b9.
//
// Solidity: function addToken(address token, uint256 apy) returns()
func (_Stake *StakeTransactor) AddToken(opts *bind.TransactOpts, token common.Address, apy *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "addToken", token, apy)
}

// AddToken is a paid mutator transaction binding the contract method 0xaf81c5b9.
//
// Solidity: function addToken(address token, uint256 apy) returns()
func (_Stake *StakeSession) AddToken(token common.Address, apy *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.AddToken(&_Stake.TransactOpts, token, apy)
}

// AddToken is a paid mutator transaction binding the contract method 0xaf81c5b9.
//
// Solidity: function addToken(address token, uint256 apy) returns()
func (_Stake *StakeTransactorSession) AddToken(token common.Address, apy *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.AddToken(&_Stake.TransactOpts, token, apy)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x0b83a727.
//
// Solidity: function claimAllRewards() returns()
func (_Stake *StakeTransactor) ClaimAllRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "claimAllRewards")
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x0b83a727.
//
// Solidity: function claimAllRewards() returns()
func (_Stake *StakeSession) ClaimAllRewards() (*types.Transaction, error) {
	return _Stake.Contract.ClaimAllRewards(&_Stake.TransactOpts)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x0b83a727.
//
// Solidity: function claimAllRewards() returns()
func (_Stake *StakeTransactorSession) ClaimAllRewards() (*types.Transaction, error) {
	return _Stake.Contract.ClaimAllRewards(&_Stake.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address token) returns()
func (_Stake *StakeTransactor) ClaimRewards(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "claimRewards", token)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address token) returns()
func (_Stake *StakeSession) ClaimRewards(token common.Address) (*types.Transaction, error) {
	return _Stake.Contract.ClaimRewards(&_Stake.TransactOpts, token)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address token) returns()
func (_Stake *StakeTransactorSession) ClaimRewards(token common.Address) (*types.Transaction, error) {
	return _Stake.Contract.ClaimRewards(&_Stake.TransactOpts, token)
}

// DepositRewards is a paid mutator transaction binding the contract method 0x97ad1cce.
//
// Solidity: function depositRewards(address token, uint256 amount) returns()
func (_Stake *StakeTransactor) DepositRewards(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "depositRewards", token, amount)
}

// DepositRewards is a paid mutator transaction binding the contract method 0x97ad1cce.
//
// Solidity: function depositRewards(address token, uint256 amount) returns()
func (_Stake *StakeSession) DepositRewards(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.DepositRewards(&_Stake.TransactOpts, token, amount)
}

// DepositRewards is a paid mutator transaction binding the contract method 0x97ad1cce.
//
// Solidity: function depositRewards(address token, uint256 amount) returns()
func (_Stake *StakeTransactorSession) DepositRewards(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.DepositRewards(&_Stake.TransactOpts, token, amount)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_Stake *StakeTransactor) RemoveToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "removeToken", token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_Stake *StakeSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _Stake.Contract.RemoveToken(&_Stake.TransactOpts, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_Stake *StakeTransactorSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _Stake.Contract.RemoveToken(&_Stake.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stake.Contract.RenounceOwnership(&_Stake.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stake.Contract.RenounceOwnership(&_Stake.TransactOpts)
}

// SetAPY is a paid mutator transaction binding the contract method 0xc4160fd3.
//
// Solidity: function setAPY(address token, uint256 newAPY) returns()
func (_Stake *StakeTransactor) SetAPY(opts *bind.TransactOpts, token common.Address, newAPY *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setAPY", token, newAPY)
}

// SetAPY is a paid mutator transaction binding the contract method 0xc4160fd3.
//
// Solidity: function setAPY(address token, uint256 newAPY) returns()
func (_Stake *StakeSession) SetAPY(token common.Address, newAPY *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetAPY(&_Stake.TransactOpts, token, newAPY)
}

// SetAPY is a paid mutator transaction binding the contract method 0xc4160fd3.
//
// Solidity: function setAPY(address token, uint256 newAPY) returns()
func (_Stake *StakeTransactorSession) SetAPY(token common.Address, newAPY *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetAPY(&_Stake.TransactOpts, token, newAPY)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address token, uint256 amount) returns()
func (_Stake *StakeTransactor) Stake(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "stake", token, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address token, uint256 amount) returns()
func (_Stake *StakeSession) Stake(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Stake(&_Stake.TransactOpts, token, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address token, uint256 amount) returns()
func (_Stake *StakeTransactorSession) Stake(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Stake(&_Stake.TransactOpts, token, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.TransferOwnership(&_Stake.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.TransferOwnership(&_Stake.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_Stake *StakeTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "withdraw", token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_Stake *StakeSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Withdraw(&_Stake.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address token, uint256 amount) returns()
func (_Stake *StakeTransactorSession) Withdraw(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Withdraw(&_Stake.TransactOpts, token, amount)
}

// StakeAPYUpdatedIterator is returned from FilterAPYUpdated and is used to iterate over the raw logs and unpacked data for APYUpdated events raised by the Stake contract.
type StakeAPYUpdatedIterator struct {
	Event *StakeAPYUpdated // Event containing the contract specifics and raw log

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
func (it *StakeAPYUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeAPYUpdated)
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
		it.Event = new(StakeAPYUpdated)
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
func (it *StakeAPYUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeAPYUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeAPYUpdated represents a APYUpdated event raised by the Stake contract.
type StakeAPYUpdated struct {
	Token  common.Address
	NewAPY *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAPYUpdated is a free log retrieval operation binding the contract event 0x95046b7fdb9bcf09cdb8b5e270934d955f51f771912290ff9a0020045195eee5.
//
// Solidity: event APYUpdated(address indexed token, uint256 newAPY)
func (_Stake *StakeFilterer) FilterAPYUpdated(opts *bind.FilterOpts, token []common.Address) (*StakeAPYUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "APYUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &StakeAPYUpdatedIterator{contract: _Stake.contract, event: "APYUpdated", logs: logs, sub: sub}, nil
}

// WatchAPYUpdated is a free log subscription operation binding the contract event 0x95046b7fdb9bcf09cdb8b5e270934d955f51f771912290ff9a0020045195eee5.
//
// Solidity: event APYUpdated(address indexed token, uint256 newAPY)
func (_Stake *StakeFilterer) WatchAPYUpdated(opts *bind.WatchOpts, sink chan<- *StakeAPYUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "APYUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeAPYUpdated)
				if err := _Stake.contract.UnpackLog(event, "APYUpdated", log); err != nil {
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

// ParseAPYUpdated is a log parse operation binding the contract event 0x95046b7fdb9bcf09cdb8b5e270934d955f51f771912290ff9a0020045195eee5.
//
// Solidity: event APYUpdated(address indexed token, uint256 newAPY)
func (_Stake *StakeFilterer) ParseAPYUpdated(log types.Log) (*StakeAPYUpdated, error) {
	event := new(StakeAPYUpdated)
	if err := _Stake.contract.UnpackLog(event, "APYUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stake contract.
type StakeOwnershipTransferredIterator struct {
	Event *StakeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeOwnershipTransferred)
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
		it.Event = new(StakeOwnershipTransferred)
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
func (it *StakeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeOwnershipTransferred represents a OwnershipTransferred event raised by the Stake contract.
type StakeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stake *StakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakeOwnershipTransferredIterator{contract: _Stake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stake *StakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeOwnershipTransferred)
				if err := _Stake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stake *StakeFilterer) ParseOwnershipTransferred(log types.Log) (*StakeOwnershipTransferred, error) {
	event := new(StakeOwnershipTransferred)
	if err := _Stake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the Stake contract.
type StakeRewardsClaimedIterator struct {
	Event *StakeRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *StakeRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRewardsClaimed)
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
		it.Event = new(StakeRewardsClaimed)
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
func (it *StakeRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRewardsClaimed represents a RewardsClaimed event raised by the Stake contract.
type StakeRewardsClaimed struct {
	User   common.Address
	Token  common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed token, uint256 reward)
func (_Stake *StakeFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*StakeRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "RewardsClaimed", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &StakeRewardsClaimedIterator{contract: _Stake.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed token, uint256 reward)
func (_Stake *StakeFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *StakeRewardsClaimed, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "RewardsClaimed", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRewardsClaimed)
				if err := _Stake.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed token, uint256 reward)
func (_Stake *StakeFilterer) ParseRewardsClaimed(log types.Log) (*StakeRewardsClaimed, error) {
	event := new(StakeRewardsClaimed)
	if err := _Stake.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRewardsDepositedIterator is returned from FilterRewardsDeposited and is used to iterate over the raw logs and unpacked data for RewardsDeposited events raised by the Stake contract.
type StakeRewardsDepositedIterator struct {
	Event *StakeRewardsDeposited // Event containing the contract specifics and raw log

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
func (it *StakeRewardsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRewardsDeposited)
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
		it.Event = new(StakeRewardsDeposited)
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
func (it *StakeRewardsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRewardsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRewardsDeposited represents a RewardsDeposited event raised by the Stake contract.
type StakeRewardsDeposited struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsDeposited is a free log retrieval operation binding the contract event 0xb8b27d0db504fa5d914f1fd330347096e88d5ff94b6c612d32797e7c12a8f66f.
//
// Solidity: event RewardsDeposited(address indexed token, uint256 amount)
func (_Stake *StakeFilterer) FilterRewardsDeposited(opts *bind.FilterOpts, token []common.Address) (*StakeRewardsDepositedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "RewardsDeposited", tokenRule)
	if err != nil {
		return nil, err
	}
	return &StakeRewardsDepositedIterator{contract: _Stake.contract, event: "RewardsDeposited", logs: logs, sub: sub}, nil
}

// WatchRewardsDeposited is a free log subscription operation binding the contract event 0xb8b27d0db504fa5d914f1fd330347096e88d5ff94b6c612d32797e7c12a8f66f.
//
// Solidity: event RewardsDeposited(address indexed token, uint256 amount)
func (_Stake *StakeFilterer) WatchRewardsDeposited(opts *bind.WatchOpts, sink chan<- *StakeRewardsDeposited, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "RewardsDeposited", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRewardsDeposited)
				if err := _Stake.contract.UnpackLog(event, "RewardsDeposited", log); err != nil {
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

// ParseRewardsDeposited is a log parse operation binding the contract event 0xb8b27d0db504fa5d914f1fd330347096e88d5ff94b6c612d32797e7c12a8f66f.
//
// Solidity: event RewardsDeposited(address indexed token, uint256 amount)
func (_Stake *StakeFilterer) ParseRewardsDeposited(log types.Log) (*StakeRewardsDeposited, error) {
	event := new(StakeRewardsDeposited)
	if err := _Stake.contract.UnpackLog(event, "RewardsDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Stake contract.
type StakeStakedIterator struct {
	Event *StakeStaked // Event containing the contract specifics and raw log

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
func (it *StakeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeStaked)
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
		it.Event = new(StakeStaked)
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
func (it *StakeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeStaked represents a Staked event raised by the Stake contract.
type StakeStaked struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed user, address indexed token, uint256 amount)
func (_Stake *StakeFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*StakeStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Staked", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &StakeStakedIterator{contract: _Stake.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed user, address indexed token, uint256 amount)
func (_Stake *StakeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakeStaked, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Staked", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeStaked)
				if err := _Stake.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed user, address indexed token, uint256 amount)
func (_Stake *StakeFilterer) ParseStaked(log types.Log) (*StakeStaked, error) {
	event := new(StakeStaked)
	if err := _Stake.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the Stake contract.
type StakeTokenAddedIterator struct {
	Event *StakeTokenAdded // Event containing the contract specifics and raw log

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
func (it *StakeTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeTokenAdded)
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
		it.Event = new(StakeTokenAdded)
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
func (it *StakeTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeTokenAdded represents a TokenAdded event raised by the Stake contract.
type StakeTokenAdded struct {
	Token common.Address
	Apy   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0xf4c563a3ea86ff1f4275e8c207df0375a51963f2b831b7bf4da8be938d92876c.
//
// Solidity: event TokenAdded(address indexed token, uint256 indexed apy)
func (_Stake *StakeFilterer) FilterTokenAdded(opts *bind.FilterOpts, token []common.Address, apy []*big.Int) (*StakeTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var apyRule []interface{}
	for _, apyItem := range apy {
		apyRule = append(apyRule, apyItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "TokenAdded", tokenRule, apyRule)
	if err != nil {
		return nil, err
	}
	return &StakeTokenAddedIterator{contract: _Stake.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0xf4c563a3ea86ff1f4275e8c207df0375a51963f2b831b7bf4da8be938d92876c.
//
// Solidity: event TokenAdded(address indexed token, uint256 indexed apy)
func (_Stake *StakeFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *StakeTokenAdded, token []common.Address, apy []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var apyRule []interface{}
	for _, apyItem := range apy {
		apyRule = append(apyRule, apyItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "TokenAdded", tokenRule, apyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeTokenAdded)
				if err := _Stake.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0xf4c563a3ea86ff1f4275e8c207df0375a51963f2b831b7bf4da8be938d92876c.
//
// Solidity: event TokenAdded(address indexed token, uint256 indexed apy)
func (_Stake *StakeFilterer) ParseTokenAdded(log types.Log) (*StakeTokenAdded, error) {
	event := new(StakeTokenAdded)
	if err := _Stake.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the Stake contract.
type StakeTokenRemovedIterator struct {
	Event *StakeTokenRemoved // Event containing the contract specifics and raw log

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
func (it *StakeTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeTokenRemoved)
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
		it.Event = new(StakeTokenRemoved)
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
func (it *StakeTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeTokenRemoved represents a TokenRemoved event raised by the Stake contract.
type StakeTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_Stake *StakeFilterer) FilterTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*StakeTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &StakeTokenRemovedIterator{contract: _Stake.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_Stake *StakeFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *StakeTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeTokenRemoved)
				if err := _Stake.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed token)
func (_Stake *StakeFilterer) ParseTokenRemoved(log types.Log) (*StakeTokenRemoved, error) {
	event := new(StakeTokenRemoved)
	if err := _Stake.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Stake contract.
type StakeWithdrawnIterator struct {
	Event *StakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeWithdrawn)
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
		it.Event = new(StakeWithdrawn)
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
func (it *StakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeWithdrawn represents a Withdrawn event raised by the Stake contract.
type StakeWithdrawn struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed user, address indexed token, uint256 amount)
func (_Stake *StakeFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*StakeWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Withdrawn", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &StakeWithdrawnIterator{contract: _Stake.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed user, address indexed token, uint256 amount)
func (_Stake *StakeFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *StakeWithdrawn, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Withdrawn", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeWithdrawn)
				if err := _Stake.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed user, address indexed token, uint256 amount)
func (_Stake *StakeFilterer) ParseWithdrawn(log types.Log) (*StakeWithdrawn, error) {
	event := new(StakeWithdrawn)
	if err := _Stake.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
