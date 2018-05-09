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
const MedconABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint256\"}],\"name\":\"addPatient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"}],\"name\":\"isExistingUser\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint256\"},{\"name\":\"pharma_id\",\"type\":\"string\"}],\"name\":\"addPharma\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint256\"},{\"name\":\"medical_id\",\"type\":\"string\"}],\"name\":\"addDoctor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"},{\"name\":\"med_name\",\"type\":\"string\"},{\"name\":\"dose_per_day\",\"type\":\"uint256\"},{\"name\":\"no_of_days\",\"type\":\"uint256\"},{\"name\":\"doctor\",\"type\":\"string\"},{\"name\":\"diagnosis\",\"type\":\"string\"}],\"name\":\"addMedOrderToPatient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MedconBin is the compiled bytecode used for deploying new contracts.
const MedconBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556112a68061003b6000396000f3006060604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630a72fa43811461007c5780632bbfb0ba1461011357806341c0e1b514610178578063c3d122da1461018b578063cc7ad2ff14610269578063d0a57add14610347575b600080fd5b341561008757600080fd5b61011160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061046792505050565b005b341561011e57600080fd5b61016460046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061067595505050505050565b604051901515815260200160405180910390f35b341561018357600080fd5b6101116106ef565b341561019657600080fd5b61011160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061073095505050505050565b341561027457600080fd5b61011160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061096995505050505050565b341561035257600080fd5b61011160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094968635966020808201359750919550606081019450604090810135860180830194503592508291601f83018190048102019051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650610b1c95505050505050565b6003836040518082805190602001908083835b602083106104995780518252601f19909201916020918201910161047a565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390206003015460ff610100909104161515600114156104e957600080fd5b6104f68284836001610f27565b6003846040518082805190602001908083835b602083106105285780518252601f199092019160209182019101610509565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181908051610570929160200190611018565b5060208201518160010190805161058b929160200190611018565b506040820151816002015560608201518160030160006101000a81548160ff021916908360038111156105ba57fe5b02179055506080820151600390910180549115156101000261ff0019909216919091179055506105e8610f78565b6002846040518082805190602001908083835b6020831061061a5780518252601f1990920191602091820191016105fb565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181600201556020820151600391909101805460ff191691151591909117905550505050565b60006003826040518082805190602001908083835b602083106106a95780518252601f19909201916020918201910161068a565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902060030154610100900460ff1692915050565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116141561072e5760005473ffffffffffffffffffffffffffffffffffffffff16ff5b565b6003846040518082805190602001908083835b602083106107625780518252601f199092019160209182019101610743565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390206003015460ff610100909104161515600114156107b257600080fd5b6107bf8385846003610f27565b6003856040518082805190602001908083835b602083106107f15780518252601f1990920191602091820191016107d2565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181908051610839929160200190611018565b50602082015181600101908051610854929160200190611018565b506040820151816002015560608201518160030160006101000a81548160ff0219169083600381111561088357fe5b02179055506080820151600390910180549115156101000261ff0019909216919091179055506108b281610f9b565b6004856040518082805190602001908083835b602083106108e45780518252601f1990920191602091820191016108c5565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902081518190805161092c929160200190611018565b50602082015181600101908051610947929160200190611092565b506040820151600291909101805460ff19169115159190911790555050505050565b6003846040518082805190602001908083835b6020831061099b5780518252601f19909201916020918201910161097c565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390206003015460ff610100909104161515600114156109eb57600080fd5b6109f88385846000610f27565b6003856040518082805190602001908083835b60208310610a2a5780518252601f199092019160209182019101610a0b565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181908051610a72929160200190611018565b50602082015181600101908051610a8d929160200190611018565b506040820151816002015560608201518160030160006101000a81548160ff02191690836003811115610abc57fe5b02179055506080820151600390910180549115156101000261ff001990921691909117905550610aeb81610f9b565b600185604051808280519060200190808383602083106108e45780518252601f1990920191602091820191016108c5565b610b246110ea565b600080610b348888888888610fb9565b92506002896040518082805190602001908083835b60208310610b685780518252601f199092019160209182019101610b49565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020915081600101856040518082805190602001908083835b60208310610bd15780518252601f199092019160209182019101610bb2565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390205460001415610c84576002820180546001908101918290558301866040518082805190602001908083835b60208310610c4e5780518252601f199092019160209182019101610c2f565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020555b81600060018201876040518082805190602001908083835b60208310610cbb5780518252601f199092019160209182019101610c9c565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902054815260208101919091526040016000206003015460ff161515610dd057610d1685610ff1565b82600060018201886040518082805190602001908083835b60208310610d4d5780518252601f199092019160209182019101610d2e565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054815260200190815260200160002060008201518160010155602082015181600201908051610db3929160200190611018565b506040820151600391909101805460ff1916911515919091179055505b81600060018201876040518082805190602001908083835b60208310610e075780518252601f199092019160209182019101610de8565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902054815260208082019290925260409081016000908120600180820180549182019055825292839052209091508390815181908051610e7d929160200190611018565b50602082015181600101908051610e98929160200190611018565b506040820151816002015560608201518160030155608082015181600401908051610ec7929160200190611018565b5060a082015160058201805460ff191691151591909117905560c08201516005820180549115156101000261ff001990921691909117905560e082015181600601908051610f19929160200190611018565b505050505050505050505050565b610f2f611147565b610f37611147565b848152602081018690526040810184905260608101836003811115610f5857fe5b90816003811115610f6557fe5b9052506001608082015295945050505050565b610f8061118e565b610f8861118e565b60016020820152600081529050805b5090565b610fa36111a5565b610fab6111a5565b918252506001604082015290565b610fc16110ea565b610fc96110ea565b958652506040850193909352606084019190915260808301526020820152600160a082015290565b610ff96111c6565b6110016111c6565b602081019290925250600081526001604082015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061105957805160ff1916838001178555611086565b82800160010185558215611086579182015b8281111561108657825182559160200191906001019061106b565b50610f979291506111e1565b8280548282559060005260206000209081019282156110de579160200282015b828111156110de578251829080516110ce929160200190611018565b50916020019190600101906110b2565b50610f979291506111fe565b610100604051908101604052806110ff611221565b815260200161110c611221565b81526020016000815260200160008152602001611127611221565b81526000602082018190526040820152606001611142611221565b905290565b60a06040519081016040528061115b611221565b8152602001611168611221565b8152602001600081526020016000600381111561118157fe5b8152600060209091015290565b604080519081016040526000808252602082015290565b6060604051908101604052806111b9611221565b8152602001611181611221565b60606040519081016040528060008152602001611181611221565b6111fb91905b80821115610f9757600081556001016111e7565b90565b6111fb91905b80821115610f975760006112188282611233565b50600101611204565b60206040519081016040526000815290565b50805460018160011615610100020316600290046000825580601f106112595750611277565b601f01602090049060005260206000209081019061127791906111e1565b505600a165627a7a72305820516b7418426365b731ff72766d7a4b92485dc329d577c5046fa54f85263ec4ea0029`

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

// AddMedOrderToPatient is a paid mutator transaction binding the contract method 0xd0a57add.
//
// Solidity: function addMedOrderToPatient(patient string, med_name string, dose_per_day uint256, no_of_days uint256, doctor string, diagnosis string) returns()
func (_Medcon *MedconTransactor) AddMedOrderToPatient(opts *bind.TransactOpts, patient string, med_name string, dose_per_day *big.Int, no_of_days *big.Int, doctor string, diagnosis string) (*types.Transaction, error) {
	return _Medcon.contract.Transact(opts, "addMedOrderToPatient", patient, med_name, dose_per_day, no_of_days, doctor, diagnosis)
}

// AddMedOrderToPatient is a paid mutator transaction binding the contract method 0xd0a57add.
//
// Solidity: function addMedOrderToPatient(patient string, med_name string, dose_per_day uint256, no_of_days uint256, doctor string, diagnosis string) returns()
func (_Medcon *MedconSession) AddMedOrderToPatient(patient string, med_name string, dose_per_day *big.Int, no_of_days *big.Int, doctor string, diagnosis string) (*types.Transaction, error) {
	return _Medcon.Contract.AddMedOrderToPatient(&_Medcon.TransactOpts, patient, med_name, dose_per_day, no_of_days, doctor, diagnosis)
}

// AddMedOrderToPatient is a paid mutator transaction binding the contract method 0xd0a57add.
//
// Solidity: function addMedOrderToPatient(patient string, med_name string, dose_per_day uint256, no_of_days uint256, doctor string, diagnosis string) returns()
func (_Medcon *MedconTransactorSession) AddMedOrderToPatient(patient string, med_name string, dose_per_day *big.Int, no_of_days *big.Int, doctor string, diagnosis string) (*types.Transaction, error) {
	return _Medcon.Contract.AddMedOrderToPatient(&_Medcon.TransactOpts, patient, med_name, dose_per_day, no_of_days, doctor, diagnosis)
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
const MortalBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905560c18061003a6000396000f300606060405260043610603e5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341c0e1b581146043575b600080fd5b3415604d57600080fd5b60536055565b005b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116141560935760005473ffffffffffffffffffffffffffffffffffffffff16ff5b5600a165627a7a723058207ef0dc0d8cb7545b16841daf4abf331f5b852568a744d14ef600662176da896d0029`

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
