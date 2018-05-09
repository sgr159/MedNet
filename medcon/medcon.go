// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package medcon

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MedconABI is the input ABI used to generate the binding from.
const MedconABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint256\"}],\"name\":\"addPatient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"}],\"name\":\"isExistingUser\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint256\"},{\"name\":\"pharma_id\",\"type\":\"string\"}],\"name\":\"addPharma\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint256\"},{\"name\":\"medical_id\",\"type\":\"string\"}],\"name\":\"addDoctor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MedconBin is the compiled bytecode used for deploying new contracts.
const MedconBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a0319909116179055610c998061003b6000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630a72fa4381146100715780632bbfb0ba1461010857806341c0e1b51461016d578063c3d122da14610180578063cc7ad2ff1461025e575b600080fd5b341561007c57600080fd5b61010660046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061033c92505050565b005b341561011357600080fd5b61015960046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061054a95505050505050565b604051901515815260200160405180910390f35b341561017857600080fd5b6101066105c4565b341561018b57600080fd5b61010660046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061060595505050505050565b341561026957600080fd5b61010660046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061083e95505050505050565b6003836040518082805190602001908083835b6020831061036e5780518252601f19909201916020918201910161034f565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390206003015460ff610100909104161515600114156103be57600080fd5b6103cb82848360016109f1565b6003846040518082805190602001908083835b602083106103fd5780518252601f1990920191602091820191016103de565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181908051610445929160200190610a83565b50602082015181600101908051610460929160200190610a83565b506040820151816002015560608201518160030160006101000a81548160ff0219169083600381111561048f57fe5b02179055506080820151600390910180549115156101000261ff0019909216919091179055506104bd610a42565b6002846040518082805190602001908083835b602083106104ef5780518252601f1990920191602091820191016104d0565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181600101556020820151600291909101805460ff191691151591909117905550505050565b60006003826040518082805190602001908083835b6020831061057e5780518252601f19909201916020918201910161055f565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902060030154610100900460ff1692915050565b6000543373ffffffffffffffffffffffffffffffffffffffff908116911614156106035760005473ffffffffffffffffffffffffffffffffffffffff16ff5b565b6003846040518082805190602001908083835b602083106106375780518252601f199092019160209182019101610618565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390206003015460ff6101009091041615156001141561068757600080fd5b61069483858460036109f1565b6003856040518082805190602001908083835b602083106106c65780518252601f1990920191602091820191016106a7565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902081518190805161070e929160200190610a83565b50602082015181600101908051610729929160200190610a83565b506040820151816002015560608201518160030160006101000a81548160ff0219169083600381111561075857fe5b02179055506080820151600390910180549115156101000261ff00199092169190911790555061078781610a65565b6004856040518082805190602001908083835b602083106107b95780518252601f19909201916020918201910161079a565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181908051610801929160200190610a83565b5060208201518160010190805161081c929160200190610afd565b506040820151600291909101805460ff19169115159190911790555050505050565b6003846040518082805190602001908083835b602083106108705780518252601f199092019160209182019101610851565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390206003015460ff610100909104161515600114156108c057600080fd5b6108cd83858460006109f1565b6003856040518082805190602001908083835b602083106108ff5780518252601f1990920191602091820191016108e0565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181908051610947929160200190610a83565b50602082015181600101908051610962929160200190610a83565b506040820151816002015560608201518160030160006101000a81548160ff0219169083600381111561099157fe5b02179055506080820151600390910180549115156101000261ff0019909216919091179055506109c081610a65565b600185604051808280519060200190808383602083106107b95780518252601f19909201916020918201910161079a565b6109f9610b55565b610a01610b55565b848152602081018690526040810184905260608101836003811115610a2257fe5b90816003811115610a2f57fe5b9052506001608082015295945050505050565b610a4a610b9c565b610a52610b9c565b60016020820152600081529050805b5090565b610a6d610bb3565b610a75610bb3565b918252506001604082015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610ac457805160ff1916838001178555610af1565b82800160010185558215610af1579182015b82811115610af1578251825591602001919060010190610ad6565b50610a61929150610bd4565b828054828255906000526020600020908101928215610b49579160200282015b82811115610b4957825182908051610b39929160200190610a83565b5091602001919060010190610b1d565b50610a61929150610bf1565b60a060405190810160405280610b69610c14565b8152602001610b76610c14565b81526020016000815260200160006003811115610b8f57fe5b8152600060209091015290565b604080519081016040526000808252602082015290565b606060405190810160405280610bc7610c14565b8152602001610b8f610c14565b610bee91905b80821115610a615760008155600101610bda565b90565b610bee91905b80821115610a61576000610c0b8282610c26565b50600101610bf7565b60206040519081016040526000815290565b50805460018160011615610100020316600290046000825580601f10610c4c5750610c6a565b601f016020900490600052602060002090810190610c6a9190610bd4565b505600a165627a7a72305820c63e10db7af5adad0e5c31e32c47fcff434badf8700a5b6074ddb205a6a7e1ae0029`

// DeployMedcon deploys a new Ethereum contract, binding an instance of Medcon to it.
func DeployMedcon(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Medcon, error) {
	parsed, err := abi.JSON(strings.NewReader(MedconABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MedconBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Medcon{MedconCaller: MedconCaller{contract: contract}, MedconTransactor: MedconTransactor{contract: contract}}, nil
}

// Medcon is an auto generated Go binding around an Ethereum contract.
type Medcon struct {
	MedconCaller     // Read-only binding to the contract
	MedconTransactor // Write-only binding to the contract
}

// MedconCaller is an auto generated read-only Go binding around an Ethereum contract.
type MedconCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedconTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MedconTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MedconSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MedconSession struct {
	Contract     *Medcon           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MedconCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MedconCallerSession struct {
	Contract *MedconCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MedconTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MedconTransactorSession struct {
	Contract     *MedconTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MedconRaw is an auto generated low-level Go binding around an Ethereum contract.
type MedconRaw struct {
	Contract *Medcon // Generic contract binding to access the raw methods on
}

// MedconCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MedconCallerRaw struct {
	Contract *MedconCaller // Generic read-only contract binding to access the raw methods on
}

// MedconTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MedconTransactorRaw struct {
	Contract *MedconTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMedcon creates a new instance of Medcon, bound to a specific deployed contract.
func NewMedcon(address common.Address, backend bind.ContractBackend) (*Medcon, error) {
	contract, err := bindMedcon(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Medcon{MedconCaller: MedconCaller{contract: contract}, MedconTransactor: MedconTransactor{contract: contract}}, nil
}

// NewMedconCaller creates a new read-only instance of Medcon, bound to a specific deployed contract.
func NewMedconCaller(address common.Address, caller bind.ContractCaller) (*MedconCaller, error) {
	contract, err := bindMedcon(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MedconCaller{contract: contract}, nil
}

// NewMedconTransactor creates a new write-only instance of Medcon, bound to a specific deployed contract.
func NewMedconTransactor(address common.Address, transactor bind.ContractTransactor) (*MedconTransactor, error) {
	contract, err := bindMedcon(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MedconTransactor{contract: contract}, nil
}

// bindMedcon binds a generic wrapper to an already deployed contract.
func bindMedcon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MedconABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Medcon *MedconRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Medcon.Contract.MedconCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Medcon *MedconRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medcon.Contract.MedconTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Medcon *MedconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Medcon.Contract.MedconTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Medcon *MedconCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Medcon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Medcon *MedconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medcon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Medcon *MedconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Medcon.Contract.contract.Transact(opts, method, params...)
}

// IsExistingUser is a free data retrieval call binding the contract method 0x2bbfb0ba.
//
// Solidity: function isExistingUser(username string) constant returns(res bool)
func (_Medcon *MedconCaller) IsExistingUser(opts *bind.CallOpts, username string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Medcon.contract.Call(opts, out, "isExistingUser", username)
	return *ret0, err
}

// IsExistingUser is a free data retrieval call binding the contract method 0x2bbfb0ba.
//
// Solidity: function isExistingUser(username string) constant returns(res bool)
func (_Medcon *MedconSession) IsExistingUser(username string) (bool, error) {
	return _Medcon.Contract.IsExistingUser(&_Medcon.CallOpts, username)
}

// IsExistingUser is a free data retrieval call binding the contract method 0x2bbfb0ba.
//
// Solidity: function isExistingUser(username string) constant returns(res bool)
func (_Medcon *MedconCallerSession) IsExistingUser(username string) (bool, error) {
	return _Medcon.Contract.IsExistingUser(&_Medcon.CallOpts, username)
}

// AddDoctor is a paid mutator transaction binding the contract method 0xcc7ad2ff.
//
// Solidity: function addDoctor(username string, name string, password uint256, medical_id string) returns()
func (_Medcon *MedconTransactor) AddDoctor(opts *bind.TransactOpts, username string, name string, password *big.Int, medical_id string) (*types.Transaction, error) {
	return _Medcon.contract.Transact(opts, "addDoctor", username, name, password, medical_id)
}

// AddDoctor is a paid mutator transaction binding the contract method 0xcc7ad2ff.
//
// Solidity: function addDoctor(username string, name string, password uint256, medical_id string) returns()
func (_Medcon *MedconSession) AddDoctor(username string, name string, password *big.Int, medical_id string) (*types.Transaction, error) {
	return _Medcon.Contract.AddDoctor(&_Medcon.TransactOpts, username, name, password, medical_id)
}

// AddDoctor is a paid mutator transaction binding the contract method 0xcc7ad2ff.
//
// Solidity: function addDoctor(username string, name string, password uint256, medical_id string) returns()
func (_Medcon *MedconTransactorSession) AddDoctor(username string, name string, password *big.Int, medical_id string) (*types.Transaction, error) {
	return _Medcon.Contract.AddDoctor(&_Medcon.TransactOpts, username, name, password, medical_id)
}

// AddPatient is a paid mutator transaction binding the contract method 0x0a72fa43.
//
// Solidity: function addPatient(username string, name string, password uint256) returns()
func (_Medcon *MedconTransactor) AddPatient(opts *bind.TransactOpts, username string, name string, password *big.Int) (*types.Transaction, error) {
	return _Medcon.contract.Transact(opts, "addPatient", username, name, password)
}

// AddPatient is a paid mutator transaction binding the contract method 0x0a72fa43.
//
// Solidity: function addPatient(username string, name string, password uint256) returns()
func (_Medcon *MedconSession) AddPatient(username string, name string, password *big.Int) (*types.Transaction, error) {
	return _Medcon.Contract.AddPatient(&_Medcon.TransactOpts, username, name, password)
}

// AddPatient is a paid mutator transaction binding the contract method 0x0a72fa43.
//
// Solidity: function addPatient(username string, name string, password uint256) returns()
func (_Medcon *MedconTransactorSession) AddPatient(username string, name string, password *big.Int) (*types.Transaction, error) {
	return _Medcon.Contract.AddPatient(&_Medcon.TransactOpts, username, name, password)
}

// AddPharma is a paid mutator transaction binding the contract method 0xc3d122da.
//
// Solidity: function addPharma(username string, name string, password uint256, pharma_id string) returns()
func (_Medcon *MedconTransactor) AddPharma(opts *bind.TransactOpts, username string, name string, password *big.Int, pharma_id string) (*types.Transaction, error) {
	return _Medcon.contract.Transact(opts, "addPharma", username, name, password, pharma_id)
}

// AddPharma is a paid mutator transaction binding the contract method 0xc3d122da.
//
// Solidity: function addPharma(username string, name string, password uint256, pharma_id string) returns()
func (_Medcon *MedconSession) AddPharma(username string, name string, password *big.Int, pharma_id string) (*types.Transaction, error) {
	return _Medcon.Contract.AddPharma(&_Medcon.TransactOpts, username, name, password, pharma_id)
}

// AddPharma is a paid mutator transaction binding the contract method 0xc3d122da.
//
// Solidity: function addPharma(username string, name string, password uint256, pharma_id string) returns()
func (_Medcon *MedconTransactorSession) AddPharma(username string, name string, password *big.Int, pharma_id string) (*types.Transaction, error) {
	return _Medcon.Contract.AddPharma(&_Medcon.TransactOpts, username, name, password, pharma_id)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Medcon *MedconTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Medcon.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Medcon *MedconSession) Kill() (*types.Transaction, error) {
	return _Medcon.Contract.Kill(&_Medcon.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Medcon *MedconTransactorSession) Kill() (*types.Transaction, error) {
	return _Medcon.Contract.Kill(&_Medcon.TransactOpts)
}

// MortalABI is the input ABI used to generate the binding from.
const MortalABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MortalBin is the compiled bytecode used for deploying new contracts.
const MortalBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905560c18061003a6000396000f300606060405260043610603e5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341c0e1b581146043575b600080fd5b3415604d57600080fd5b60536055565b005b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116141560935760005473ffffffffffffffffffffffffffffffffffffffff16ff5b5600a165627a7a723058209c1099298addbbb08b966aaad56d8f3ea3f238fdb9135c9a72bc609cbf5e09c60029`

// DeployMortal deploys a new Ethereum contract, binding an instance of Mortal to it.
func DeployMortal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mortal, error) {
	parsed, err := abi.JSON(strings.NewReader(MortalABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MortalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mortal{MortalCaller: MortalCaller{contract: contract}, MortalTransactor: MortalTransactor{contract: contract}}, nil
}

// Mortal is an auto generated Go binding around an Ethereum contract.
type Mortal struct {
	MortalCaller     // Read-only binding to the contract
	MortalTransactor // Write-only binding to the contract
}

// MortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type MortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MortalSession struct {
	Contract     *Mortal           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MortalCallerSession struct {
	Contract *MortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MortalTransactorSession struct {
	Contract     *MortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type MortalRaw struct {
	Contract *Mortal // Generic contract binding to access the raw methods on
}

// MortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MortalCallerRaw struct {
	Contract *MortalCaller // Generic read-only contract binding to access the raw methods on
}

// MortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MortalTransactorRaw struct {
	Contract *MortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMortal creates a new instance of Mortal, bound to a specific deployed contract.
func NewMortal(address common.Address, backend bind.ContractBackend) (*Mortal, error) {
	contract, err := bindMortal(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mortal{MortalCaller: MortalCaller{contract: contract}, MortalTransactor: MortalTransactor{contract: contract}}, nil
}

// NewMortalCaller creates a new read-only instance of Mortal, bound to a specific deployed contract.
func NewMortalCaller(address common.Address, caller bind.ContractCaller) (*MortalCaller, error) {
	contract, err := bindMortal(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MortalCaller{contract: contract}, nil
}

// NewMortalTransactor creates a new write-only instance of Mortal, bound to a specific deployed contract.
func NewMortalTransactor(address common.Address, transactor bind.ContractTransactor) (*MortalTransactor, error) {
	contract, err := bindMortal(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MortalTransactor{contract: contract}, nil
}

// bindMortal binds a generic wrapper to an already deployed contract.
func bindMortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MortalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mortal *MortalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mortal.Contract.MortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mortal *MortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.Contract.MortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mortal *MortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mortal.Contract.MortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mortal *MortalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mortal *MortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mortal *MortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mortal.Contract.contract.Transact(opts, method, params...)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortal.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalSession) Kill() (*types.Transaction, error) {
	return _Mortal.Contract.Kill(&_Mortal.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Mortal *MortalTransactorSession) Kill() (*types.Transaction, error) {
	return _Mortal.Contract.Kill(&_Mortal.TransactOpts)
}
