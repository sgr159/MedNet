// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package medcon

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MedconABI is the input ABI used to generate the binding from.
const MedconABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"},{\"name\":\"indx\",\"type\":\"uint64\"}],\"name\":\"showNumOfMedOrdersByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"}],\"name\":\"isExistingUser\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"},{\"name\":\"doctor\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"showMedOrderByDoctor\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"},{\"name\":\"doctor\",\"type\":\"string\"}],\"name\":\"showNumOfMedOrdersByDoc\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"}],\"name\":\"showNumOfPrescriptions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"},{\"name\":\"med_name\",\"type\":\"string\"},{\"name\":\"dose_per_day\",\"type\":\"uint64\"},{\"name\":\"no_of_days\",\"type\":\"uint64\"},{\"name\":\"doctor\",\"type\":\"string\"},{\"name\":\"diagnosis\",\"type\":\"string\"}],\"name\":\"addMedOrderToPatient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint64\"},{\"name\":\"pharma_id\",\"type\":\"string\"}],\"name\":\"addPharma\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"patient\",\"type\":\"string\"},{\"name\":\"pres\",\"type\":\"uint64\"},{\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"showMedOrderByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint64\"}],\"name\":\"addPatient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"password\",\"type\":\"uint64\"},{\"name\":\"medical_id\",\"type\":\"string\"}],\"name\":\"addDoctor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MedconBin is the compiled bytecode used for deploying new contracts.
const MedconBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556120a18061003b6000396000f3006060604052600436106100ae5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663195451c881146100b35780632bbfb0ba1461012d57806341b53c881461019257806341c0e1b5146104095780636196d53b1461041e578063632047c0146104b15780636934821514610502578063774b4a37146106325780637ff0539c1461071a578063d31a1dfc1461077e578063f790812d1461081d575b600080fd5b34156100be57600080fd5b61011060046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505050923567ffffffffffffffff169250610905915050565b60405167ffffffffffffffff909116815260200160405180910390f35b341561013857600080fd5b61017e60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506109bb95505050505050565b604051901515815260200160405180910390f35b341561019d57600080fd5b61023160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505050923567ffffffffffffffff169250610a3d915050565b60405167ffffffffffffffff80871660408301528516606082015282151560a082015260e080825281906020820190608083019060c084019084018c818151815260200191508051906020019080838360005b8381101561029c578082015183820152602001610284565b50505050905090810190601f1680156102c95780820380516001836020036101000a031916815260200191505b5085810384528b818151815260200191508051906020019080838360005b838110156102ff5780820151838201526020016102e7565b50505050905090810190601f16801561032c5780820380516001836020036101000a031916815260200191505b50858103835288818151815260200191508051906020019080838360005b8381101561036257808201518382015260200161034a565b50505050905090810190601f16801561038f5780820380516001836020036101000a031916815260200191505b50858103825286818151815260200191508051906020019080838360005b838110156103c55780820151838201526020016103ad565b50505050905090810190601f1680156103f25780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390f35b341561041457600080fd5b61041c610e59565b005b341561042957600080fd5b61011060046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650610e9a95505050505050565b34156104bc57600080fd5b61011060046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610f9d95505050505050565b341561050d57600080fd5b61041c60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949667ffffffffffffffff87358116976020808201359092169750919550606082019450604091820135860180820194503592508291601f830182900482029091019051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061101b95505050505050565b341561063d57600080fd5b61041c60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803567ffffffffffffffff1690602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061150895505050505050565b341561072557600080fd5b61023160046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505067ffffffffffffffff85358116956020013516935061178092505050565b341561078957600080fd5b61041c60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505050923567ffffffffffffffff16925061189d915050565b341561082857600080fd5b61041c60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803567ffffffffffffffff1690602001909190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650611b1d95505050505050565b60008060006002856040518082805190602001908083835b6020831061093c5780518252601f19909201916020918201910161091d565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020600281015490925067ffffffffffffffff908116908516111561099257600080fd5b5067ffffffffffffffff9283166000908152602091909152604090206001015490911692915050565b60006003826040518082805190602001908083835b602083106109ef5780518252601f1990920191602091820191016109d0565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020600201546901000000000000000000900460ff1692915050565b610a45611e16565b610a4d611e16565b600080610a58611e16565b6000610a62611e16565b600080600060028d6040518082805190602001908083835b60208310610a995780518252601f199092019160209182019101610a7a565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390209250826000600182018e6040518082805190602001908083835b60208310610b065780518252601f199092019160209182019101610ae7565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205467ffffffffffffffff9081168252602082019290925260400160002060018101549093508116908c1610610b6e57600080fd5b5067ffffffffffffffff8a811660009081526020838152604091829020600280820154600483015483549396879660018089019785841697680100000000000000009096049093169560038a01956101009586900460ff169560058c01958c9594831615909102600019019091169190910491601f8301829004820290910190519081016040528092919081815260200182805460018160011615610100020316600290048015610c605780601f10610c3557610100808354040283529160200191610c60565b820191906000526020600020905b815481529060010190602001808311610c4357829003601f168201915b50505050509650858054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cfc5780601f10610cd157610100808354040283529160200191610cfc565b820191906000526020600020905b815481529060010190602001808311610cdf57829003601f168201915b50505050509550828054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d985780601f10610d6d57610100808354040283529160200191610d98565b820191906000526020600020905b815481529060010190602001808311610d7b57829003601f168201915b50505050509250808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e345780601f10610e0957610100808354040283529160200191610e34565b820191906000526020600020905b815481529060010190602001808311610e1757829003601f168201915b5050505050905099509950995099509950995099505050509397509397509397909450565b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161415610e985760005473ffffffffffffffffffffffffffffffffffffffff16ff5b565b60008060006002856040518082805190602001908083835b60208310610ed15780518252601f199092019160209182019101610eb2565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020915081600060018201866040518082805190602001908083835b60208310610f3e5780518252601f199092019160209182019101610f1f565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205467ffffffffffffffff90811682526020820192909252604001600020600101541695945050505050565b6000806002836040518082805190602001908083835b60208310610fd25780518252601f199092019160209182019101610fb3565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390206002015467ffffffffffffffff169392505050565b611023611e28565b6000806110338888888888611d0f565b92506002896040518082805190602001908083835b602083106110675780518252601f199092019160209182019101611048565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020915081600101856040518082805190602001908083835b602083106110d05780518252601f1990920191602091820191016110b1565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205467ffffffffffffffff1615156111c95760028201805467ffffffffffffffff198116600167ffffffffffffffff928316810192831691909117909255908301866040518082805190602001908083835b602083106111745780518252601f199092019160209182019101611155565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020805467ffffffffffffffff191667ffffffffffffffff929092169190911790555b81600060018201876040518082805190602001908083835b602083106112005780518252601f1990920191602091820191016111e1565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205467ffffffffffffffff16815260208101919091526040016000206003015460ff1615156113485761126585611d54565b82600060018201886040518082805190602001908083835b6020831061129c5780518252601f19909201916020918201910161127d565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205467ffffffffffffffff1681526020810191909152604001600020815160018201805467ffffffffffffffff191667ffffffffffffffff9290921691909117905560208201518160020190805161132b929160200190611e85565b506040820151600391909101805460ff1916911515919091179055505b81600060018201876040518082805190602001908083835b6020831061137f5780518252601f199092019160209182019101611360565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390205467ffffffffffffffff9081168252602080830193909352604091820160009081206001808201805467ffffffffffffffff19811690861692830190951694909417909355918152928190529120909150839081518190805161141a929160200190611e85565b50602082015181600101908051611435929160200190611e85565b50604082015160028201805467ffffffffffffffff191667ffffffffffffffff9290921691909117905560608201518160020160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506080820151816003019080516114a8929160200190611e85565b5060a082015160048201805460ff191691151591909117905560c08201516004820180549115156101000261ff001990921691909117905560e0820151816005019080516114fa929160200190611e85565b505050505050505050505050565b6003846040518082805190602001908083835b6020831061153a5780518252601f19909201916020918201910161151b565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020600201546901000000000000000000900460ff1615156001141561159157600080fd5b61159e8385846003611d7b565b6003856040518082805190602001908083835b602083106115d05780518252601f1990920191602091820191016115b1565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181908051611618929160200190611e85565b50602082015181600101908051611633929160200190611e85565b50604082015160028201805467ffffffffffffffff191667ffffffffffffffff92909216919091179055606082015160028201805468ff000000000000000019166801000000000000000083600381111561168a57fe5b021790555060808201516002909101805491151569010000000000000000000269ff00000000000000000019909216919091179055506116c981611dd5565b6004856040518082805190602001908083835b602083106116fb5780518252601f1990920191602091820191016116dc565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181908051611743929160200190611e85565b5060208201518160010190805161175e929160200190611eff565b506040820151600291909101805460ff19169115159190911790555050505050565b611788611e16565b611790611e16565b60008061179b611e16565b60006117a5611e16565b600080600060028d6040518082805190602001908083835b602083106117dc5780518252601f1990920191602091820191016117bd565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020600281015490935068010000000000000000900460ff16151561183257600080fd5b600283015467ffffffffffffffff808e169116111561185057600080fd5b67ffffffffffffffff8c166000908152602084905260409020600381015490925060ff16151561187f57600080fd5b600182015467ffffffffffffffff908116908c1610610b6e57600080fd5b6003836040518082805190602001908083835b602083106118cf5780518252601f1990920191602091820191016118b0565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020600201546901000000000000000000900460ff1615156001141561192657600080fd5b6119338284836001611d7b565b6003846040518082805190602001908083835b602083106119655780518252601f199092019160209182019101611946565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040519081900390208151819080516119ad929160200190611e85565b506020820151816001019080516119c8929160200190611e85565b50604082015160028201805467ffffffffffffffff191667ffffffffffffffff92909216919091179055606082015160028201805468ff0000000000000000191668010000000000000000836003811115611a1f57fe5b021790555060808201516002909101805491151569010000000000000000000269ff0000000000000000001990921691909117905550611a5d611df3565b6002846040518082805190602001908083835b60208310611a8f5780518252601f199092019160209182019101611a70565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815160028201805467ffffffffffffffff191667ffffffffffffffff92909216919091179055602082015160029091018054911515680100000000000000000268ff00000000000000001990921691909117905550505050565b6003846040518082805190602001908083835b60208310611b4f5780518252601f199092019160209182019101611b30565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020600201546901000000000000000000900460ff16151560011415611ba657600080fd5b611bb38385846000611d7b565b6003856040518082805190602001908083835b60208310611be55780518252601f199092019160209182019101611bc6565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051908190039020815181908051611c2d929160200190611e85565b50602082015181600101908051611c48929160200190611e85565b50604082015160028201805467ffffffffffffffff191667ffffffffffffffff92909216919091179055606082015160028201805468ff0000000000000000191668010000000000000000836003811115611c9f57fe5b021790555060808201516002909101805491151569010000000000000000000269ff0000000000000000001990921691909117905550611cde81611dd5565b600185604051808280519060200190808383602083106116fb5780518252601f1990920191602091820191016116dc565b611d17611e28565b611d1f611e28565b9586525067ffffffffffffffff938416604086015291909216606084015260808301919091526020820152600160a082015290565b611d5c611f57565b611d64611f57565b602081019290925250600081526001604082015290565b611d83611f7f565b611d8b611f7f565b8481526020810186905267ffffffffffffffff8416604082015260608101836003811115611db557fe5b90816003811115611dc257fe5b9052506001608082015295945050505050565b611ddd611fb6565b611de5611fb6565b918252506001604082015290565b611dfb611fd7565b611e03611fd7565b60016020820152600081529050805b5090565b60206040519081016040526000815290565b61010060405190810160405280611e3d611e16565b8152602001611e4a611e16565b81526000602082018190526040820152606001611e65611e16565b81526000602082018190526040820152606001611e80611e16565b905290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611ec657805160ff1916838001178555611ef3565b82800160010185558215611ef3579182015b82811115611ef3578251825591602001919060010190611ed8565b50611e12929150611fee565b828054828255906000526020600020908101928215611f4b579160200282015b82811115611f4b57825182908051611f3b929160200190611e85565b5091602001919060010190611f1f565b50611e1292915061200b565b60606040519081016040526000815260208101611f72611e16565b8152600060209091015290565b60a060405190810160405280611f93611e16565b8152602001611fa0611e16565b8152600060208201819052604090910190611f72565b606060405190810160405280611fca611e16565b8152602001611f72611e16565b604080519081016040526000808252602082015290565b61200891905b80821115611e125760008155600101611ff4565b90565b61200891905b80821115611e12576000612025828261202e565b50600101612011565b50805460018160011615610100020316600290046000825580601f106120545750612072565b601f0160209004906000526020600020908101906120729190611fee565b505600a165627a7a7230582079a4706d2c5c3b904a1e51520d07d94732d49f06dbd552b7d81bcda56f617aad0029`

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

// ShowMedOrderByDoctor is a free data retrieval call binding the contract method 0x41b53c88.
//
// Solidity: function showMedOrderByDoctor(patient string, doctor string, index uint64) constant returns(string, string, uint64, uint64, string, bool, string)
func (_Medcon *MedconCaller) ShowMedOrderByDoctor(opts *bind.CallOpts, patient string, doctor string, index uint64) (string, string, uint64, uint64, string, bool, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(uint64)
		ret3 = new(uint64)
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

// ShowMedOrderByDoctor is a free data retrieval call binding the contract method 0x41b53c88.
//
// Solidity: function showMedOrderByDoctor(patient string, doctor string, index uint64) constant returns(string, string, uint64, uint64, string, bool, string)
func (_Medcon *MedconSession) ShowMedOrderByDoctor(patient string, doctor string, index uint64) (string, string, uint64, uint64, string, bool, string, error) {
	return _Medcon.Contract.ShowMedOrderByDoctor(&_Medcon.CallOpts, patient, doctor, index)
}

// ShowMedOrderByDoctor is a free data retrieval call binding the contract method 0x41b53c88.
//
// Solidity: function showMedOrderByDoctor(patient string, doctor string, index uint64) constant returns(string, string, uint64, uint64, string, bool, string)
func (_Medcon *MedconCallerSession) ShowMedOrderByDoctor(patient string, doctor string, index uint64) (string, string, uint64, uint64, string, bool, string, error) {
	return _Medcon.Contract.ShowMedOrderByDoctor(&_Medcon.CallOpts, patient, doctor, index)
}

// ShowMedOrderByIndex is a free data retrieval call binding the contract method 0x7ff0539c.
//
// Solidity: function showMedOrderByIndex(patient string, pres uint64, index uint64) constant returns(string, string, uint64, uint64, string, bool, string)
func (_Medcon *MedconCaller) ShowMedOrderByIndex(opts *bind.CallOpts, patient string, pres uint64, index uint64) (string, string, uint64, uint64, string, bool, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(uint64)
		ret3 = new(uint64)
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

// ShowMedOrderByIndex is a free data retrieval call binding the contract method 0x7ff0539c.
//
// Solidity: function showMedOrderByIndex(patient string, pres uint64, index uint64) constant returns(string, string, uint64, uint64, string, bool, string)
func (_Medcon *MedconSession) ShowMedOrderByIndex(patient string, pres uint64, index uint64) (string, string, uint64, uint64, string, bool, string, error) {
	return _Medcon.Contract.ShowMedOrderByIndex(&_Medcon.CallOpts, patient, pres, index)
}

// ShowMedOrderByIndex is a free data retrieval call binding the contract method 0x7ff0539c.
//
// Solidity: function showMedOrderByIndex(patient string, pres uint64, index uint64) constant returns(string, string, uint64, uint64, string, bool, string)
func (_Medcon *MedconCallerSession) ShowMedOrderByIndex(patient string, pres uint64, index uint64) (string, string, uint64, uint64, string, bool, string, error) {
	return _Medcon.Contract.ShowMedOrderByIndex(&_Medcon.CallOpts, patient, pres, index)
}

// ShowNumOfMedOrdersByDoc is a free data retrieval call binding the contract method 0x6196d53b.
//
// Solidity: function showNumOfMedOrdersByDoc(patient string, doctor string) constant returns(uint64)
func (_Medcon *MedconCaller) ShowNumOfMedOrdersByDoc(opts *bind.CallOpts, patient string, doctor string) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Medcon.contract.Call(opts, out, "showNumOfMedOrdersByDoc", patient, doctor)
	return *ret0, err
}

// ShowNumOfMedOrdersByDoc is a free data retrieval call binding the contract method 0x6196d53b.
//
// Solidity: function showNumOfMedOrdersByDoc(patient string, doctor string) constant returns(uint64)
func (_Medcon *MedconSession) ShowNumOfMedOrdersByDoc(patient string, doctor string) (uint64, error) {
	return _Medcon.Contract.ShowNumOfMedOrdersByDoc(&_Medcon.CallOpts, patient, doctor)
}

// ShowNumOfMedOrdersByDoc is a free data retrieval call binding the contract method 0x6196d53b.
//
// Solidity: function showNumOfMedOrdersByDoc(patient string, doctor string) constant returns(uint64)
func (_Medcon *MedconCallerSession) ShowNumOfMedOrdersByDoc(patient string, doctor string) (uint64, error) {
	return _Medcon.Contract.ShowNumOfMedOrdersByDoc(&_Medcon.CallOpts, patient, doctor)
}

// ShowNumOfMedOrdersByIndex is a free data retrieval call binding the contract method 0x195451c8.
//
// Solidity: function showNumOfMedOrdersByIndex(patient string, indx uint64) constant returns(uint64)
func (_Medcon *MedconCaller) ShowNumOfMedOrdersByIndex(opts *bind.CallOpts, patient string, indx uint64) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Medcon.contract.Call(opts, out, "showNumOfMedOrdersByIndex", patient, indx)
	return *ret0, err
}

// ShowNumOfMedOrdersByIndex is a free data retrieval call binding the contract method 0x195451c8.
//
// Solidity: function showNumOfMedOrdersByIndex(patient string, indx uint64) constant returns(uint64)
func (_Medcon *MedconSession) ShowNumOfMedOrdersByIndex(patient string, indx uint64) (uint64, error) {
	return _Medcon.Contract.ShowNumOfMedOrdersByIndex(&_Medcon.CallOpts, patient, indx)
}

// ShowNumOfMedOrdersByIndex is a free data retrieval call binding the contract method 0x195451c8.
//
// Solidity: function showNumOfMedOrdersByIndex(patient string, indx uint64) constant returns(uint64)
func (_Medcon *MedconCallerSession) ShowNumOfMedOrdersByIndex(patient string, indx uint64) (uint64, error) {
	return _Medcon.Contract.ShowNumOfMedOrdersByIndex(&_Medcon.CallOpts, patient, indx)
}

// ShowNumOfPrescriptions is a free data retrieval call binding the contract method 0x632047c0.
//
// Solidity: function showNumOfPrescriptions(patient string) constant returns(uint64)
func (_Medcon *MedconCaller) ShowNumOfPrescriptions(opts *bind.CallOpts, patient string) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Medcon.contract.Call(opts, out, "showNumOfPrescriptions", patient)
	return *ret0, err
}

// ShowNumOfPrescriptions is a free data retrieval call binding the contract method 0x632047c0.
//
// Solidity: function showNumOfPrescriptions(patient string) constant returns(uint64)
func (_Medcon *MedconSession) ShowNumOfPrescriptions(patient string) (uint64, error) {
	return _Medcon.Contract.ShowNumOfPrescriptions(&_Medcon.CallOpts, patient)
}

// ShowNumOfPrescriptions is a free data retrieval call binding the contract method 0x632047c0.
//
// Solidity: function showNumOfPrescriptions(patient string) constant returns(uint64)
func (_Medcon *MedconCallerSession) ShowNumOfPrescriptions(patient string) (uint64, error) {
	return _Medcon.Contract.ShowNumOfPrescriptions(&_Medcon.CallOpts, patient)
}

// AddDoctor is a paid mutator transaction binding the contract method 0xf790812d.
//
// Solidity: function addDoctor(username string, name string, password uint64, medical_id string) returns()
func (_Medcon *MedconTransactor) AddDoctor(opts *bind.TransactOpts, username string, name string, password uint64, medical_id string) (*types.Transaction, error) {
	return _Medcon.contract.Transact(opts, "addDoctor", username, name, password, medical_id)
}

// AddDoctor is a paid mutator transaction binding the contract method 0xf790812d.
//
// Solidity: function addDoctor(username string, name string, password uint64, medical_id string) returns()
func (_Medcon *MedconSession) AddDoctor(username string, name string, password uint64, medical_id string) (*types.Transaction, error) {
	return _Medcon.Contract.AddDoctor(&_Medcon.TransactOpts, username, name, password, medical_id)
}

// AddDoctor is a paid mutator transaction binding the contract method 0xf790812d.
//
// Solidity: function addDoctor(username string, name string, password uint64, medical_id string) returns()
func (_Medcon *MedconTransactorSession) AddDoctor(username string, name string, password uint64, medical_id string) (*types.Transaction, error) {
	return _Medcon.Contract.AddDoctor(&_Medcon.TransactOpts, username, name, password, medical_id)
}

// AddMedOrderToPatient is a paid mutator transaction binding the contract method 0x69348215.
//
// Solidity: function addMedOrderToPatient(patient string, med_name string, dose_per_day uint64, no_of_days uint64, doctor string, diagnosis string) returns()
func (_Medcon *MedconTransactor) AddMedOrderToPatient(opts *bind.TransactOpts, patient string, med_name string, dose_per_day uint64, no_of_days uint64, doctor string, diagnosis string) (*types.Transaction, error) {
	return _Medcon.contract.Transact(opts, "addMedOrderToPatient", patient, med_name, dose_per_day, no_of_days, doctor, diagnosis)
}

// AddMedOrderToPatient is a paid mutator transaction binding the contract method 0x69348215.
//
// Solidity: function addMedOrderToPatient(patient string, med_name string, dose_per_day uint64, no_of_days uint64, doctor string, diagnosis string) returns()
func (_Medcon *MedconSession) AddMedOrderToPatient(patient string, med_name string, dose_per_day uint64, no_of_days uint64, doctor string, diagnosis string) (*types.Transaction, error) {
	return _Medcon.Contract.AddMedOrderToPatient(&_Medcon.TransactOpts, patient, med_name, dose_per_day, no_of_days, doctor, diagnosis)
}

// AddMedOrderToPatient is a paid mutator transaction binding the contract method 0x69348215.
//
// Solidity: function addMedOrderToPatient(patient string, med_name string, dose_per_day uint64, no_of_days uint64, doctor string, diagnosis string) returns()
func (_Medcon *MedconTransactorSession) AddMedOrderToPatient(patient string, med_name string, dose_per_day uint64, no_of_days uint64, doctor string, diagnosis string) (*types.Transaction, error) {
	return _Medcon.Contract.AddMedOrderToPatient(&_Medcon.TransactOpts, patient, med_name, dose_per_day, no_of_days, doctor, diagnosis)
}

// AddPatient is a paid mutator transaction binding the contract method 0xd31a1dfc.
//
// Solidity: function addPatient(username string, name string, password uint64) returns()
func (_Medcon *MedconTransactor) AddPatient(opts *bind.TransactOpts, username string, name string, password uint64) (*types.Transaction, error) {
	return _Medcon.contract.Transact(opts, "addPatient", username, name, password)
}

// AddPatient is a paid mutator transaction binding the contract method 0xd31a1dfc.
//
// Solidity: function addPatient(username string, name string, password uint64) returns()
func (_Medcon *MedconSession) AddPatient(username string, name string, password uint64) (*types.Transaction, error) {
	return _Medcon.Contract.AddPatient(&_Medcon.TransactOpts, username, name, password)
}

// AddPatient is a paid mutator transaction binding the contract method 0xd31a1dfc.
//
// Solidity: function addPatient(username string, name string, password uint64) returns()
func (_Medcon *MedconTransactorSession) AddPatient(username string, name string, password uint64) (*types.Transaction, error) {
	return _Medcon.Contract.AddPatient(&_Medcon.TransactOpts, username, name, password)
}

// AddPharma is a paid mutator transaction binding the contract method 0x774b4a37.
//
// Solidity: function addPharma(username string, name string, password uint64, pharma_id string) returns()
func (_Medcon *MedconTransactor) AddPharma(opts *bind.TransactOpts, username string, name string, password uint64, pharma_id string) (*types.Transaction, error) {
	return _Medcon.contract.Transact(opts, "addPharma", username, name, password, pharma_id)
}

// AddPharma is a paid mutator transaction binding the contract method 0x774b4a37.
//
// Solidity: function addPharma(username string, name string, password uint64, pharma_id string) returns()
func (_Medcon *MedconSession) AddPharma(username string, name string, password uint64, pharma_id string) (*types.Transaction, error) {
	return _Medcon.Contract.AddPharma(&_Medcon.TransactOpts, username, name, password, pharma_id)
}

// AddPharma is a paid mutator transaction binding the contract method 0x774b4a37.
//
// Solidity: function addPharma(username string, name string, password uint64, pharma_id string) returns()
func (_Medcon *MedconTransactorSession) AddPharma(username string, name string, password uint64, pharma_id string) (*types.Transaction, error) {
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
const MortalBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905560c18061003a6000396000f300606060405260043610603e5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341c0e1b581146043575b600080fd5b3415604d57600080fd5b60536055565b005b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116141560935760005473ffffffffffffffffffffffffffffffffffffffff16ff5b5600a165627a7a7230582030cd8b89ee655abd1cbdda7d3a4dd8a98053d796f8332ece44c1c16320193e190029`

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
