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
const MedconABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint256\"}],\"name\":\"addPatient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"}],\"name\":\"isExistingUser\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"},{\"name\":\"doctor\",\"type\":\"string\"}],\"name\":\"showNumOfMedOrdersByDoc\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"}],\"name\":\"showNumOfPrescriptions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"},{\"name\":\"pres\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"showMedOrderByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint256\"},{\"name\":\"pharma_id\",\"type\":\"string\"}],\"name\":\"addPharma\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint256\"},{\"name\":\"medical_id\",\"type\":\"string\"}],\"name\":\"addDoctor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"},{\"name\":\"med_name\",\"type\":\"string\"},{\"name\":\"dose_per_day\",\"type\":\"uint256\"},{\"name\":\"no_of_days\",\"type\":\"uint256\"},{\"name\":\"doctor\",\"type\":\"string\"},{\"name\":\"diagnosis\",\"type\":\"string\"}],\"name\":\"addMedOrderToPatient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"},{\"name\":\"doctor\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"showMedOrderByDoctor\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"},{\"name\":\"indx\",\"type\":\"uint256\"}],\"name\":\"showNumOfMedOrdersByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MedconBin is the compiled bytecode used for deploying new contracts.
const MedconBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a0319909116179055611db88061003b6000396000f3006060604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630a72fa4381146100b35780632bbfb0ba1461014a57806341c0e1b5146101af5780636196d53b146101c2578063632047c014610267578063ac300fc4146102b8578063c3d122da146104dc578063cc7ad2ff146105ba578063d0a57add14610698578063eaa90eeb146107b8578063fcbc37261461084d575b600080fd5b34156100be57600080fd5b61014860046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965050933593506108a092505050565b005b341561015557600080fd5b61019b60046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610aae95505050505050565b604051901515815260200160405180910390f35b34156101ba57600080fd5b610148610b28565b34156101cd57600080fd5b61025560046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650610b6995505050505050565b60405190815260200160405180910390f35b341561027257600080fd5b61025560046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610c5f95505050505050565b34156102c357600080fd5b61031060046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050843594602001359350610cd392505050565b60405180806020018060200188815260200187815260200180602001861515151581526020018060200185810385528c818151815260200191508051906020019080838360005b8381101561036f578082015183820152602001610357565b50505050905090810190601f16801561039c5780820380516001836020036101000a031916815260200191505b5085810384528b818151815260200191508051906020019080838360005b838110156103d25780820151838201526020016103ba565b50505050905090810190601f1680156103ff5780820380516001836020036101000a031916815260200191505b50858103835288818151815260200191508051906020019080838360005b8381101561043557808201518382015260200161041d565b50505050905090810190601f1680156104625780820380516001836020036101000a031916815260200191505b50858103825286818151815260200191508051906020019080838360005b83811015610498578082015183820152602001610480565b50505050905090810190601f1680156104c55780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390f35b34156104e757600080fd5b61014860046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061108f95505050505050565b34156105c557600080fd5b61014860046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506112c895505050505050565b34156106a357600080fd5b61014860046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094968635966020808201359750919550606081019450604090810135860180830194503592508291601f83018190048102019051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061147b95505050505050565b34156107c357600080fd5b61031060046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061188692505050565b341561085857600080fd5b61025560046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050933593506119a592505050565b6003836040518082805190602001908083835b602083106108d25780518252601f1990920191602091820191016108b3565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390206003015460ff6101009091041615156001141561092257600080fd5b61092f8284836001611a39565b6003846040518082805190602001908083835b602083106109615780518252601f199092019160209182019101610942565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390208151819080516109a9929160200190611b2a565b506020820151816001019080516109c4929160200190611b2a565b506040820151816002015560608201518160030160006101000a81548160ff021916908360038111156109f357fe5b02179055506080820151600390910180549115156101000261ff001990921691909117905550610a21611a8a565b6002846040518082805190602001908083835b60208310610a535780518252601f199092019160209182019101610a34565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181600201556020820151600391909101805460ff191691151591909117905550505050565b60006003826040518082805190602001908083835b60208310610ae25780518252601f199092019160209182019101610ac3565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902060030154610100900460ff1692915050565b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161415610b675760005473ffffffffffffffffffffffffffffffffffffffff16ff5b565b60008060006002856040518082805190602001908083835b60208310610ba05780518252601f199092019160209182019101610b81565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020915081600060018201866040518082805190602001908083835b60208310610c0d5780518252601f199092019160209182019101610bee565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902054815260208101919091526040016000206001015495945050505050565b6000806002836040518082805190602001908083835b60208310610c945780518252601f199092019160209182019101610c75565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020600201549392505050565b610cdb611ba4565b610ce3611ba4565b600080610cee611ba4565b6000610cf8611ba4565b600080600060028d6040518082805190602001908083835b60208310610d2f5780518252601f199092019160209182019101610d10565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020600381015490935060ff161515610d7957600080fd5b8b83600201541115610d8a57600080fd5b60008c8152602084905260409020600381015490925060ff161515610dae57600080fd5b60018201548b10610dbe57600080fd5b8160000160008c81526020019081526020016000209050806000018160010182600201548360030154846004018560050160019054906101000a900460ff1686600601868054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e965780601f10610e6b57610100808354040283529160200191610e96565b820191906000526020600020905b815481529060010190602001808311610e7957829003601f168201915b50505050509650858054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f325780601f10610f0757610100808354040283529160200191610f32565b820191906000526020600020905b815481529060010190602001808311610f1557829003601f168201915b50505050509550828054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fce5780601f10610fa357610100808354040283529160200191610fce565b820191906000526020600020905b815481529060010190602001808311610fb157829003601f168201915b50505050509250808054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561106a5780601f1061103f5761010080835404028352916020019161106a565b820191906000526020600020905b81548152906001019060200180831161104d57829003601f168201915b5050505050905099509950995099509950995099505050509397509397509397909450565b6003846040518082805190602001908083835b602083106110c15780518252601f1990920191602091820191016110a2565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390206003015460ff6101009091041615156001141561111157600080fd5b61111e8385846003611a39565b6003856040518082805190602001908083835b602083106111505780518252601f199092019160209182019101611131565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181908051611198929160200190611b2a565b506020820151816001019080516111b3929160200190611b2a565b506040820151816002015560608201518160030160006101000a81548160ff021916908360038111156111e257fe5b02179055506080820151600390910180549115156101000261ff00199092169190911790555061121181611aad565b6004856040518082805190602001908083835b602083106112435780518252601f199092019160209182019101611224565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902081518190805161128b929160200190611b2a565b506020820151816001019080516112a6929160200190611bb6565b506040820151600291909101805460ff19169115159190911790555050505050565b6003846040518082805190602001908083835b602083106112fa5780518252601f1990920191602091820191016112db565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390206003015460ff6101009091041615156001141561134a57600080fd5b6113578385846000611a39565b6003856040518082805190602001908083835b602083106113895780518252601f19909201916020918201910161136a565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390208151819080516113d1929160200190611b2a565b506020820151816001019080516113ec929160200190611b2a565b506040820151816002015560608201518160030160006101000a81548160ff0219169083600381111561141b57fe5b02179055506080820151600390910180549115156101000261ff00199092169190911790555061144a81611aad565b600185604051808280519060200190808383602083106112435780518252601f199092019160209182019101611224565b611483611c0e565b6000806114938888888888611acb565b92506002896040518082805190602001908083835b602083106114c75780518252601f1990920191602091820191016114a8565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020915081600101856040518082805190602001908083835b602083106115305780518252601f199092019160209182019101611511565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054600014156115e3576002820180546001908101918290558301866040518082805190602001908083835b602083106115ad5780518252601f19909201916020918201910161158e565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020555b81600060018201876040518082805190602001908083835b6020831061161a5780518252601f1990920191602091820191016115fb565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405190819003902054815260208101919091526040016000206003015460ff16151561172f5761167585611b03565b82600060018201886040518082805190602001908083835b602083106116ac5780518252601f19909201916020918201910161168d565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054815260200190815260200160002060008201518160010155602082015181600201908051611712929160200190611b2a565b506040820151600391909101805460ff1916911515919091179055505b81600060018201876040518082805190602001908083835b602083106117665780518252601f199092019160209182019101611747565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020548152602080820192909252604090810160009081206001808201805491820190558252928390522090915083908151819080516117dc929160200190611b2a565b506020820151816001019080516117f7929160200190611b2a565b506040820151816002015560608201518160030155608082015181600401908051611826929160200190611b2a565b5060a082015160058201805460ff191691151591909117905560c08201516005820180549115156101000261ff001990921691909117905560e082015181600601908051611878929160200190611b2a565b505050505050505050505050565b61188e611ba4565b611896611ba4565b6000806118a1611ba4565b60006118ab611ba4565b600080600060028d6040518082805190602001908083835b602083106118e25780518252601f1990920191602091820191016118c3565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390209250826000600182018e6040518082805190602001908083835b6020831061194f5780518252601f199092019160209182019101611930565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020548152602001908152602001600020915081600101548b101515610dbe57600080fd5b60008060006002856040518082805190602001908083835b602083106119dc5780518252601f1990920191602091820191016119bd565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902091508160020154841115611a2257600080fd5b506000928352602052506040902060010154919050565b611a41611c6b565b611a49611c6b565b848152602081018690526040810184905260608101836003811115611a6a57fe5b90816003811115611a7757fe5b9052506001608082015295945050505050565b611a92611cb2565b611a9a611cb2565b60016020820152600081529050805b5090565b611ab5611cc9565b611abd611cc9565b918252506001604082015290565b611ad3611c0e565b611adb611c0e565b958652506040850193909352606084019190915260808301526020820152600160a082015290565b611b0b611cea565b611b13611cea565b602081019290925250600081526001604082015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611b6b57805160ff1916838001178555611b98565b82800160010185558215611b98579182015b82811115611b98578251825591602001919060010190611b7d565b50611aa9929150611d05565b60206040519081016040526000815290565b828054828255906000526020600020908101928215611c02579160200282015b82811115611c0257825182908051611bf2929160200190611b2a565b5091602001919060010190611bd6565b50611aa9929150611d22565b61010060405190810160405280611c23611ba4565b8152602001611c30611ba4565b81526020016000815260200160008152602001611c4b611ba4565b81526000602082018190526040820152606001611c66611ba4565b905290565b60a060405190810160405280611c7f611ba4565b8152602001611c8c611ba4565b81526020016000815260200160006003811115611ca557fe5b8152600060209091015290565b604080519081016040526000808252602082015290565b606060405190810160405280611cdd611ba4565b8152602001611ca5611ba4565b60606040519081016040528060008152602001611ca5611ba4565b611d1f91905b80821115611aa95760008155600101611d0b565b90565b611d1f91905b80821115611aa9576000611d3c8282611d45565b50600101611d28565b50805460018160011615610100020316600290046000825580601f10611d6b5750611d89565b601f016020900490600052602060002090810190611d899190611d05565b505600a165627a7a72305820640e1dbe98fec89dc7bccd4cbd588f4f6d1c5123cc02dabf7aa7a1e517de415d0029`

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

// ShowMedOrderByDoctor is a free data retrieval call binding the contract method 0xeaa90eeb.
//
// Solidity: function showMedOrderByDoctor(patient string, doctor string, index uint256) constant returns(string, string, uint256, uint256, string, bool, string)
func (_Medcon *MedconCaller) ShowMedOrderByDoctor(opts *bind.CallOpts, patient string, doctor string, index *big.Int) (string, string, *big.Int, *big.Int, string, bool, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(string)
		ret5 = new(bool)
		ret6 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _Medcon.contract.Call(opts, out, "showMedOrderByDoctor", patient, doctor, index)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// ShowMedOrderByDoctor is a free data retrieval call binding the contract method 0xeaa90eeb.
//
// Solidity: function showMedOrderByDoctor(patient string, doctor string, index uint256) constant returns(string, string, uint256, uint256, string, bool, string)
func (_Medcon *MedconSession) ShowMedOrderByDoctor(patient string, doctor string, index *big.Int) (string, string, *big.Int, *big.Int, string, bool, string, error) {
	return _Medcon.Contract.ShowMedOrderByDoctor(&_Medcon.CallOpts, patient, doctor, index)
}

// ShowMedOrderByDoctor is a free data retrieval call binding the contract method 0xeaa90eeb.
//
// Solidity: function showMedOrderByDoctor(patient string, doctor string, index uint256) constant returns(string, string, uint256, uint256, string, bool, string)
func (_Medcon *MedconCallerSession) ShowMedOrderByDoctor(patient string, doctor string, index *big.Int) (string, string, *big.Int, *big.Int, string, bool, string, error) {
	return _Medcon.Contract.ShowMedOrderByDoctor(&_Medcon.CallOpts, patient, doctor, index)
}

// ShowMedOrderByIndex is a free data retrieval call binding the contract method 0xac300fc4.
//
// Solidity: function showMedOrderByIndex(patient string, pres uint256, index uint256) constant returns(string, string, uint256, uint256, string, bool, string)
func (_Medcon *MedconCaller) ShowMedOrderByIndex(opts *bind.CallOpts, patient string, pres *big.Int, index *big.Int) (string, string, *big.Int, *big.Int, string, bool, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(string)
		ret5 = new(bool)
		ret6 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _Medcon.contract.Call(opts, out, "showMedOrderByIndex", patient, pres, index)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// ShowMedOrderByIndex is a free data retrieval call binding the contract method 0xac300fc4.
//
// Solidity: function showMedOrderByIndex(patient string, pres uint256, index uint256) constant returns(string, string, uint256, uint256, string, bool, string)
func (_Medcon *MedconSession) ShowMedOrderByIndex(patient string, pres *big.Int, index *big.Int) (string, string, *big.Int, *big.Int, string, bool, string, error) {
	return _Medcon.Contract.ShowMedOrderByIndex(&_Medcon.CallOpts, patient, pres, index)
}

// ShowMedOrderByIndex is a free data retrieval call binding the contract method 0xac300fc4.
//
// Solidity: function showMedOrderByIndex(patient string, pres uint256, index uint256) constant returns(string, string, uint256, uint256, string, bool, string)
func (_Medcon *MedconCallerSession) ShowMedOrderByIndex(patient string, pres *big.Int, index *big.Int) (string, string, *big.Int, *big.Int, string, bool, string, error) {
	return _Medcon.Contract.ShowMedOrderByIndex(&_Medcon.CallOpts, patient, pres, index)
}

// ShowNumOfMedOrdersByDoc is a free data retrieval call binding the contract method 0x6196d53b.
//
// Solidity: function showNumOfMedOrdersByDoc(patient string, doctor string) constant returns(uint256)
func (_Medcon *MedconCaller) ShowNumOfMedOrdersByDoc(opts *bind.CallOpts, patient string, doctor string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Medcon.contract.Call(opts, out, "showNumOfMedOrdersByDoc", patient, doctor)
	return *ret0, err
}

// ShowNumOfMedOrdersByDoc is a free data retrieval call binding the contract method 0x6196d53b.
//
// Solidity: function showNumOfMedOrdersByDoc(patient string, doctor string) constant returns(uint256)
func (_Medcon *MedconSession) ShowNumOfMedOrdersByDoc(patient string, doctor string) (*big.Int, error) {
	return _Medcon.Contract.ShowNumOfMedOrdersByDoc(&_Medcon.CallOpts, patient, doctor)
}

// ShowNumOfMedOrdersByDoc is a free data retrieval call binding the contract method 0x6196d53b.
//
// Solidity: function showNumOfMedOrdersByDoc(patient string, doctor string) constant returns(uint256)
func (_Medcon *MedconCallerSession) ShowNumOfMedOrdersByDoc(patient string, doctor string) (*big.Int, error) {
	return _Medcon.Contract.ShowNumOfMedOrdersByDoc(&_Medcon.CallOpts, patient, doctor)
}

// ShowNumOfMedOrdersByIndex is a free data retrieval call binding the contract method 0xfcbc3726.
//
// Solidity: function showNumOfMedOrdersByIndex(patient string, indx uint256) constant returns(uint256)
func (_Medcon *MedconCaller) ShowNumOfMedOrdersByIndex(opts *bind.CallOpts, patient string, indx *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Medcon.contract.Call(opts, out, "showNumOfMedOrdersByIndex", patient, indx)
	return *ret0, err
}

// ShowNumOfMedOrdersByIndex is a free data retrieval call binding the contract method 0xfcbc3726.
//
// Solidity: function showNumOfMedOrdersByIndex(patient string, indx uint256) constant returns(uint256)
func (_Medcon *MedconSession) ShowNumOfMedOrdersByIndex(patient string, indx *big.Int) (*big.Int, error) {
	return _Medcon.Contract.ShowNumOfMedOrdersByIndex(&_Medcon.CallOpts, patient, indx)
}

// ShowNumOfMedOrdersByIndex is a free data retrieval call binding the contract method 0xfcbc3726.
//
// Solidity: function showNumOfMedOrdersByIndex(patient string, indx uint256) constant returns(uint256)
func (_Medcon *MedconCallerSession) ShowNumOfMedOrdersByIndex(patient string, indx *big.Int) (*big.Int, error) {
	return _Medcon.Contract.ShowNumOfMedOrdersByIndex(&_Medcon.CallOpts, patient, indx)
}

// ShowNumOfPrescriptions is a free data retrieval call binding the contract method 0x632047c0.
//
// Solidity: function showNumOfPrescriptions(patient string) constant returns(uint256)
func (_Medcon *MedconCaller) ShowNumOfPrescriptions(opts *bind.CallOpts, patient string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Medcon.contract.Call(opts, out, "showNumOfPrescriptions", patient)
	return *ret0, err
}

// ShowNumOfPrescriptions is a free data retrieval call binding the contract method 0x632047c0.
//
// Solidity: function showNumOfPrescriptions(patient string) constant returns(uint256)
func (_Medcon *MedconSession) ShowNumOfPrescriptions(patient string) (*big.Int, error) {
	return _Medcon.Contract.ShowNumOfPrescriptions(&_Medcon.CallOpts, patient)
}

// ShowNumOfPrescriptions is a free data retrieval call binding the contract method 0x632047c0.
//
// Solidity: function showNumOfPrescriptions(patient string) constant returns(uint256)
func (_Medcon *MedconCallerSession) ShowNumOfPrescriptions(patient string) (*big.Int, error) {
	return _Medcon.Contract.ShowNumOfPrescriptions(&_Medcon.CallOpts, patient)
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
const MortalBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905560c18061003a6000396000f300606060405260043610603e5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341c0e1b581146043575b600080fd5b3415604d57600080fd5b60536055565b005b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116141560935760005473ffffffffffffffffffffffffffffffffffffffff16ff5b5600a165627a7a72305820573853eae8ec847c0165ab4c603cac69a60e72ed22f18ee9c487af8bd9a37cf80029`

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
