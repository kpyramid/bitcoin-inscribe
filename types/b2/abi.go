// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package b2

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

// B2MetaData contains all meta data concerning the B2 contract.
var B2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"l1Address\",\"type\":\"string\"}],\"name\":\"lockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"l1Address\",\"type\":\"string\"}],\"name\":\"unLockEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"collection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getAddressNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"l1Address\",\"type\":\"string\"}],\"name\":\"getL1AvailableTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Address\",\"type\":\"address\"}],\"name\":\"getL2LockedTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isTokenLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"l1Address\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"setInitialData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenLockedAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"l1Address\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"l2Address\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"unLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"}],\"name\":\"updateSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// B2ABI is the input ABI used to generate the binding from.
// Deprecated: Use B2MetaData.ABI instead.
var B2ABI = B2MetaData.ABI

// B2 is an auto generated Go binding around an Ethereum contract.
type B2 struct {
	B2Caller     // Read-only binding to the contract
	B2Transactor // Write-only binding to the contract
	B2Filterer   // Log filterer for contract events
}

// B2Caller is an auto generated read-only Go binding around an Ethereum contract.
type B2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// B2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type B2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// B2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type B2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// B2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type B2Session struct {
	Contract     *B2               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// B2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type B2CallerSession struct {
	Contract *B2Caller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// B2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type B2TransactorSession struct {
	Contract     *B2Transactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// B2Raw is an auto generated low-level Go binding around an Ethereum contract.
type B2Raw struct {
	Contract *B2 // Generic contract binding to access the raw methods on
}

// B2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type B2CallerRaw struct {
	Contract *B2Caller // Generic read-only contract binding to access the raw methods on
}

// B2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type B2TransactorRaw struct {
	Contract *B2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewB2 creates a new instance of B2, bound to a specific deployed contract.
func NewB2(address common.Address, backend bind.ContractBackend) (*B2, error) {
	contract, err := bindB2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &B2{B2Caller: B2Caller{contract: contract}, B2Transactor: B2Transactor{contract: contract}, B2Filterer: B2Filterer{contract: contract}}, nil
}

// NewB2Caller creates a new read-only instance of B2, bound to a specific deployed contract.
func NewB2Caller(address common.Address, caller bind.ContractCaller) (*B2Caller, error) {
	contract, err := bindB2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &B2Caller{contract: contract}, nil
}

// NewB2Transactor creates a new write-only instance of B2, bound to a specific deployed contract.
func NewB2Transactor(address common.Address, transactor bind.ContractTransactor) (*B2Transactor, error) {
	contract, err := bindB2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &B2Transactor{contract: contract}, nil
}

// NewB2Filterer creates a new log filterer instance of B2, bound to a specific deployed contract.
func NewB2Filterer(address common.Address, filterer bind.ContractFilterer) (*B2Filterer, error) {
	contract, err := bindB2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &B2Filterer{contract: contract}, nil
}

// bindB2 binds a generic wrapper to an already deployed contract.
func bindB2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := B2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_B2 *B2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _B2.Contract.B2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_B2 *B2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _B2.Contract.B2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_B2 *B2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _B2.Contract.B2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_B2 *B2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _B2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_B2 *B2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _B2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_B2 *B2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _B2.Contract.contract.Transact(opts, method, params...)
}

// Collection is a free data retrieval call binding the contract method 0x7de1e536.
//
// Solidity: function collection() view returns(address)
func (_B2 *B2Caller) Collection(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _B2.contract.Call(opts, &out, "collection")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collection is a free data retrieval call binding the contract method 0x7de1e536.
//
// Solidity: function collection() view returns(address)
func (_B2 *B2Session) Collection() (common.Address, error) {
	return _B2.Contract.Collection(&_B2.CallOpts)
}

// Collection is a free data retrieval call binding the contract method 0x7de1e536.
//
// Solidity: function collection() view returns(address)
func (_B2 *B2CallerSession) Collection() (common.Address, error) {
	return _B2.Contract.Collection(&_B2.CallOpts)
}

// GetAddressNonce is a free data retrieval call binding the contract method 0xa62e73ef.
//
// Solidity: function getAddressNonce(address _address) view returns(uint256)
func (_B2 *B2Caller) GetAddressNonce(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _B2.contract.Call(opts, &out, "getAddressNonce", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddressNonce is a free data retrieval call binding the contract method 0xa62e73ef.
//
// Solidity: function getAddressNonce(address _address) view returns(uint256)
func (_B2 *B2Session) GetAddressNonce(_address common.Address) (*big.Int, error) {
	return _B2.Contract.GetAddressNonce(&_B2.CallOpts, _address)
}

// GetAddressNonce is a free data retrieval call binding the contract method 0xa62e73ef.
//
// Solidity: function getAddressNonce(address _address) view returns(uint256)
func (_B2 *B2CallerSession) GetAddressNonce(_address common.Address) (*big.Int, error) {
	return _B2.Contract.GetAddressNonce(&_B2.CallOpts, _address)
}

// GetL1AvailableTokens is a free data retrieval call binding the contract method 0xf622a994.
//
// Solidity: function getL1AvailableTokens(string l1Address) view returns(uint256[])
func (_B2 *B2Caller) GetL1AvailableTokens(opts *bind.CallOpts, l1Address string) ([]*big.Int, error) {
	var out []interface{}
	err := _B2.contract.Call(opts, &out, "getL1AvailableTokens", l1Address)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetL1AvailableTokens is a free data retrieval call binding the contract method 0xf622a994.
//
// Solidity: function getL1AvailableTokens(string l1Address) view returns(uint256[])
func (_B2 *B2Session) GetL1AvailableTokens(l1Address string) ([]*big.Int, error) {
	return _B2.Contract.GetL1AvailableTokens(&_B2.CallOpts, l1Address)
}

// GetL1AvailableTokens is a free data retrieval call binding the contract method 0xf622a994.
//
// Solidity: function getL1AvailableTokens(string l1Address) view returns(uint256[])
func (_B2 *B2CallerSession) GetL1AvailableTokens(l1Address string) ([]*big.Int, error) {
	return _B2.Contract.GetL1AvailableTokens(&_B2.CallOpts, l1Address)
}

// GetL2LockedTokens is a free data retrieval call binding the contract method 0x4d931015.
//
// Solidity: function getL2LockedTokens(address l2Address) view returns(uint256[])
func (_B2 *B2Caller) GetL2LockedTokens(opts *bind.CallOpts, l2Address common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _B2.contract.Call(opts, &out, "getL2LockedTokens", l2Address)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetL2LockedTokens is a free data retrieval call binding the contract method 0x4d931015.
//
// Solidity: function getL2LockedTokens(address l2Address) view returns(uint256[])
func (_B2 *B2Session) GetL2LockedTokens(l2Address common.Address) ([]*big.Int, error) {
	return _B2.Contract.GetL2LockedTokens(&_B2.CallOpts, l2Address)
}

// GetL2LockedTokens is a free data retrieval call binding the contract method 0x4d931015.
//
// Solidity: function getL2LockedTokens(address l2Address) view returns(uint256[])
func (_B2 *B2CallerSession) GetL2LockedTokens(l2Address common.Address) ([]*big.Int, error) {
	return _B2.Contract.GetL2LockedTokens(&_B2.CallOpts, l2Address)
}

// IsTokenLocked is a free data retrieval call binding the contract method 0x276a28a3.
//
// Solidity: function isTokenLocked(uint256 ) view returns(bool)
func (_B2 *B2Caller) IsTokenLocked(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _B2.contract.Call(opts, &out, "isTokenLocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenLocked is a free data retrieval call binding the contract method 0x276a28a3.
//
// Solidity: function isTokenLocked(uint256 ) view returns(bool)
func (_B2 *B2Session) IsTokenLocked(arg0 *big.Int) (bool, error) {
	return _B2.Contract.IsTokenLocked(&_B2.CallOpts, arg0)
}

// IsTokenLocked is a free data retrieval call binding the contract method 0x276a28a3.
//
// Solidity: function isTokenLocked(uint256 ) view returns(bool)
func (_B2 *B2CallerSession) IsTokenLocked(arg0 *big.Int) (bool, error) {
	return _B2.Contract.IsTokenLocked(&_B2.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_B2 *B2Caller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _B2.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_B2 *B2Session) Nonce(arg0 common.Address) (*big.Int, error) {
	return _B2.Contract.Nonce(&_B2.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_B2 *B2CallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _B2.Contract.Nonce(&_B2.CallOpts, arg0)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_B2 *B2Caller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _B2.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_B2 *B2Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _B2.Contract.OnERC721Received(&_B2.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_B2 *B2CallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _B2.Contract.OnERC721Received(&_B2.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_B2 *B2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _B2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_B2 *B2Session) Owner() (common.Address, error) {
	return _B2.Contract.Owner(&_B2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_B2 *B2CallerSession) Owner() (common.Address, error) {
	return _B2.Contract.Owner(&_B2.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_B2 *B2Caller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _B2.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_B2 *B2Session) Signer() (common.Address, error) {
	return _B2.Contract.Signer(&_B2.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_B2 *B2CallerSession) Signer() (common.Address, error) {
	return _B2.Contract.Signer(&_B2.CallOpts)
}

// TokenLockedAddress is a free data retrieval call binding the contract method 0x9a1be046.
//
// Solidity: function tokenLockedAddress(uint256 ) view returns(string l1Address, address l2Address)
func (_B2 *B2Caller) TokenLockedAddress(opts *bind.CallOpts, arg0 *big.Int) (struct {
	L1Address string
	L2Address common.Address
}, error) {
	var out []interface{}
	err := _B2.contract.Call(opts, &out, "tokenLockedAddress", arg0)

	outstruct := new(struct {
		L1Address string
		L2Address common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.L1Address = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.L2Address = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// TokenLockedAddress is a free data retrieval call binding the contract method 0x9a1be046.
//
// Solidity: function tokenLockedAddress(uint256 ) view returns(string l1Address, address l2Address)
func (_B2 *B2Session) TokenLockedAddress(arg0 *big.Int) (struct {
	L1Address string
	L2Address common.Address
}, error) {
	return _B2.Contract.TokenLockedAddress(&_B2.CallOpts, arg0)
}

// TokenLockedAddress is a free data retrieval call binding the contract method 0x9a1be046.
//
// Solidity: function tokenLockedAddress(uint256 ) view returns(string l1Address, address l2Address)
func (_B2 *B2CallerSession) TokenLockedAddress(arg0 *big.Int) (struct {
	L1Address string
	L2Address common.Address
}, error) {
	return _B2.Contract.TokenLockedAddress(&_B2.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_B2 *B2Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _B2.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_B2 *B2Session) Initialize() (*types.Transaction, error) {
	return _B2.Contract.Initialize(&_B2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_B2 *B2TransactorSession) Initialize() (*types.Transaction, error) {
	return _B2.Contract.Initialize(&_B2.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xe1efb596.
//
// Solidity: function lock(string l1Address, uint256 tokenId) returns()
func (_B2 *B2Transactor) Lock(opts *bind.TransactOpts, l1Address string, tokenId *big.Int) (*types.Transaction, error) {
	return _B2.contract.Transact(opts, "lock", l1Address, tokenId)
}

// Lock is a paid mutator transaction binding the contract method 0xe1efb596.
//
// Solidity: function lock(string l1Address, uint256 tokenId) returns()
func (_B2 *B2Session) Lock(l1Address string, tokenId *big.Int) (*types.Transaction, error) {
	return _B2.Contract.Lock(&_B2.TransactOpts, l1Address, tokenId)
}

// Lock is a paid mutator transaction binding the contract method 0xe1efb596.
//
// Solidity: function lock(string l1Address, uint256 tokenId) returns()
func (_B2 *B2TransactorSession) Lock(l1Address string, tokenId *big.Int) (*types.Transaction, error) {
	return _B2.Contract.Lock(&_B2.TransactOpts, l1Address, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_B2 *B2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _B2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_B2 *B2Session) RenounceOwnership() (*types.Transaction, error) {
	return _B2.Contract.RenounceOwnership(&_B2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_B2 *B2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _B2.Contract.RenounceOwnership(&_B2.TransactOpts)
}

// SetInitialData is a paid mutator transaction binding the contract method 0xa43c22a2.
//
// Solidity: function setInitialData(address _collection, address _signer) returns()
func (_B2 *B2Transactor) SetInitialData(opts *bind.TransactOpts, _collection common.Address, _signer common.Address) (*types.Transaction, error) {
	return _B2.contract.Transact(opts, "setInitialData", _collection, _signer)
}

// SetInitialData is a paid mutator transaction binding the contract method 0xa43c22a2.
//
// Solidity: function setInitialData(address _collection, address _signer) returns()
func (_B2 *B2Session) SetInitialData(_collection common.Address, _signer common.Address) (*types.Transaction, error) {
	return _B2.Contract.SetInitialData(&_B2.TransactOpts, _collection, _signer)
}

// SetInitialData is a paid mutator transaction binding the contract method 0xa43c22a2.
//
// Solidity: function setInitialData(address _collection, address _signer) returns()
func (_B2 *B2TransactorSession) SetInitialData(_collection common.Address, _signer common.Address) (*types.Transaction, error) {
	return _B2.Contract.SetInitialData(&_B2.TransactOpts, _collection, _signer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_B2 *B2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _B2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_B2 *B2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _B2.Contract.TransferOwnership(&_B2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_B2 *B2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _B2.Contract.TransferOwnership(&_B2.TransactOpts, newOwner)
}

// UnLock is a paid mutator transaction binding the contract method 0x9849552b.
//
// Solidity: function unLock(uint256 tokenId, bytes signature) returns()
func (_B2 *B2Transactor) UnLock(opts *bind.TransactOpts, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _B2.contract.Transact(opts, "unLock", tokenId, signature)
}

// UnLock is a paid mutator transaction binding the contract method 0x9849552b.
//
// Solidity: function unLock(uint256 tokenId, bytes signature) returns()
func (_B2 *B2Session) UnLock(tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _B2.Contract.UnLock(&_B2.TransactOpts, tokenId, signature)
}

// UnLock is a paid mutator transaction binding the contract method 0x9849552b.
//
// Solidity: function unLock(uint256 tokenId, bytes signature) returns()
func (_B2 *B2TransactorSession) UnLock(tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _B2.Contract.UnLock(&_B2.TransactOpts, tokenId, signature)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address newSigner) returns()
func (_B2 *B2Transactor) UpdateSigner(opts *bind.TransactOpts, newSigner common.Address) (*types.Transaction, error) {
	return _B2.contract.Transact(opts, "updateSigner", newSigner)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address newSigner) returns()
func (_B2 *B2Session) UpdateSigner(newSigner common.Address) (*types.Transaction, error) {
	return _B2.Contract.UpdateSigner(&_B2.TransactOpts, newSigner)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address newSigner) returns()
func (_B2 *B2TransactorSession) UpdateSigner(newSigner common.Address) (*types.Transaction, error) {
	return _B2.Contract.UpdateSigner(&_B2.TransactOpts, newSigner)
}

// B2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the B2 contract.
type B2InitializedIterator struct {
	Event *B2Initialized // Event containing the contract specifics and raw log

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
func (it *B2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(B2Initialized)
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
		it.Event = new(B2Initialized)
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
func (it *B2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *B2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// B2Initialized represents a Initialized event raised by the B2 contract.
type B2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_B2 *B2Filterer) FilterInitialized(opts *bind.FilterOpts) (*B2InitializedIterator, error) {

	logs, sub, err := _B2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &B2InitializedIterator{contract: _B2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_B2 *B2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *B2Initialized) (event.Subscription, error) {

	logs, sub, err := _B2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(B2Initialized)
				if err := _B2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_B2 *B2Filterer) ParseInitialized(log types.Log) (*B2Initialized, error) {
	event := new(B2Initialized)
	if err := _B2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// B2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the B2 contract.
type B2OwnershipTransferredIterator struct {
	Event *B2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *B2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(B2OwnershipTransferred)
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
		it.Event = new(B2OwnershipTransferred)
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
func (it *B2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *B2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// B2OwnershipTransferred represents a OwnershipTransferred event raised by the B2 contract.
type B2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_B2 *B2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*B2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _B2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &B2OwnershipTransferredIterator{contract: _B2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_B2 *B2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *B2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _B2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(B2OwnershipTransferred)
				if err := _B2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_B2 *B2Filterer) ParseOwnershipTransferred(log types.Log) (*B2OwnershipTransferred, error) {
	event := new(B2OwnershipTransferred)
	if err := _B2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// B2LockEventIterator is returned from FilterLockEvent and is used to iterate over the raw logs and unpacked data for LockEvent events raised by the B2 contract.
type B2LockEventIterator struct {
	Event *B2LockEvent // Event containing the contract specifics and raw log

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
func (it *B2LockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(B2LockEvent)
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
		it.Event = new(B2LockEvent)
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
func (it *B2LockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *B2LockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// B2LockEvent represents a LockEvent event raised by the B2 contract.
type B2LockEvent struct {
	L2Address common.Address
	TokenId   *big.Int
	Nonce     *big.Int
	L1Address string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLockEvent is a free log retrieval operation binding the contract event 0xaea85ef757763ea5a60381b591933471cfc3af19cc8e30287cc85060199ab16b.
//
// Solidity: event lockEvent(address indexed l2Address, uint256 indexed tokenId, uint256 indexed nonce, string l1Address)
func (_B2 *B2Filterer) FilterLockEvent(opts *bind.FilterOpts, l2Address []common.Address, tokenId []*big.Int, nonce []*big.Int) (*B2LockEventIterator, error) {

	var l2AddressRule []interface{}
	for _, l2AddressItem := range l2Address {
		l2AddressRule = append(l2AddressRule, l2AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _B2.contract.FilterLogs(opts, "lockEvent", l2AddressRule, tokenIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &B2LockEventIterator{contract: _B2.contract, event: "lockEvent", logs: logs, sub: sub}, nil
}

// WatchLockEvent is a free log subscription operation binding the contract event 0xaea85ef757763ea5a60381b591933471cfc3af19cc8e30287cc85060199ab16b.
//
// Solidity: event lockEvent(address indexed l2Address, uint256 indexed tokenId, uint256 indexed nonce, string l1Address)
func (_B2 *B2Filterer) WatchLockEvent(opts *bind.WatchOpts, sink chan<- *B2LockEvent, l2Address []common.Address, tokenId []*big.Int, nonce []*big.Int) (event.Subscription, error) {

	var l2AddressRule []interface{}
	for _, l2AddressItem := range l2Address {
		l2AddressRule = append(l2AddressRule, l2AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _B2.contract.WatchLogs(opts, "lockEvent", l2AddressRule, tokenIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(B2LockEvent)
				if err := _B2.contract.UnpackLog(event, "lockEvent", log); err != nil {
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

// ParseLockEvent is a log parse operation binding the contract event 0xaea85ef757763ea5a60381b591933471cfc3af19cc8e30287cc85060199ab16b.
//
// Solidity: event lockEvent(address indexed l2Address, uint256 indexed tokenId, uint256 indexed nonce, string l1Address)
func (_B2 *B2Filterer) ParseLockEvent(log types.Log) (*B2LockEvent, error) {
	event := new(B2LockEvent)
	if err := _B2.contract.UnpackLog(event, "lockEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// B2UnLockEventIterator is returned from FilterUnLockEvent and is used to iterate over the raw logs and unpacked data for UnLockEvent events raised by the B2 contract.
type B2UnLockEventIterator struct {
	Event *B2UnLockEvent // Event containing the contract specifics and raw log

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
func (it *B2UnLockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(B2UnLockEvent)
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
		it.Event = new(B2UnLockEvent)
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
func (it *B2UnLockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *B2UnLockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// B2UnLockEvent represents a UnLockEvent event raised by the B2 contract.
type B2UnLockEvent struct {
	L2Address common.Address
	TokenId   *big.Int
	Nonce     *big.Int
	L1Address string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnLockEvent is a free log retrieval operation binding the contract event 0x06b724a116219967b7673783e60f242efd969688499501b4cfd82a1e54863206.
//
// Solidity: event unLockEvent(address indexed l2Address, uint256 indexed tokenId, uint256 indexed nonce, string l1Address)
func (_B2 *B2Filterer) FilterUnLockEvent(opts *bind.FilterOpts, l2Address []common.Address, tokenId []*big.Int, nonce []*big.Int) (*B2UnLockEventIterator, error) {

	var l2AddressRule []interface{}
	for _, l2AddressItem := range l2Address {
		l2AddressRule = append(l2AddressRule, l2AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _B2.contract.FilterLogs(opts, "unLockEvent", l2AddressRule, tokenIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &B2UnLockEventIterator{contract: _B2.contract, event: "unLockEvent", logs: logs, sub: sub}, nil
}

// WatchUnLockEvent is a free log subscription operation binding the contract event 0x06b724a116219967b7673783e60f242efd969688499501b4cfd82a1e54863206.
//
// Solidity: event unLockEvent(address indexed l2Address, uint256 indexed tokenId, uint256 indexed nonce, string l1Address)
func (_B2 *B2Filterer) WatchUnLockEvent(opts *bind.WatchOpts, sink chan<- *B2UnLockEvent, l2Address []common.Address, tokenId []*big.Int, nonce []*big.Int) (event.Subscription, error) {

	var l2AddressRule []interface{}
	for _, l2AddressItem := range l2Address {
		l2AddressRule = append(l2AddressRule, l2AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _B2.contract.WatchLogs(opts, "unLockEvent", l2AddressRule, tokenIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(B2UnLockEvent)
				if err := _B2.contract.UnpackLog(event, "unLockEvent", log); err != nil {
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

// ParseUnLockEvent is a log parse operation binding the contract event 0x06b724a116219967b7673783e60f242efd969688499501b4cfd82a1e54863206.
//
// Solidity: event unLockEvent(address indexed l2Address, uint256 indexed tokenId, uint256 indexed nonce, string l1Address)
func (_B2 *B2Filterer) ParseUnLockEvent(log types.Log) (*B2UnLockEvent, error) {
	event := new(B2UnLockEvent)
	if err := _B2.contract.UnpackLog(event, "unLockEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
