package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/sgr159/MedNet/medcon"
)

func main() {
	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}

	sim := backends.NewSimulatedBackend(alloc)

	// Deploy a token contract on the simulated blockchain
	_, _, token, err := medcon.DeployMedcon(auth, sim)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
//	// Print the current (non existent) and pending name of the contract
//	name, _ := token.Name(nil)
//	fmt.Println("Pre-mining name:", name)
//
//	name, _ = token.Name(&bind.CallOpts{Pending: true})
//	fmt.Println("Pre-mining pending name:", name)
//
//	// Commit all pending transactions in the simulator and print the names again
	token.AddDoctor(auth,"doc1","doc1",123,"id123");
	token.AddDoctor(auth,"doc2","doc2",123,"id123");
	token.AddPatient(auth,"krishna","krishna",123);
	token.AddPharma(auth,"patanjali","patanjali",123,"sad");

	sim.Commit()

	token.AddMedOrderToPatient(auth,"krishna","saridon",2,7,"doc1","headache")
	token.AddMedOrderToPatient(auth,"krishna","pcm",2,7,"doc1","fever")
	token.AddMedOrderToPatient(auth,"krishna","weed",2,7,"doc2","depression")
	sim.Commit()

	token.MarkMedOrderAsMet(auth,"krishna","doc1","patanjali",1)
	sim.Commit()

	num_pres,_ := token.ShowNumOfPrescriptions(nil,"krishna")
	fmt.Println("num of pres:",num_pres)
	for i:=uint64(1);i<=num_pres;i++ {
		num_med,_ := token.ShowNumOfMedOrdersByIndex(nil,"krishna", i)
		for j:=uint64(0);j<num_med;j++ {
			a,b,c,d,e,f,g,h := token.ShowMedOrderByIndex(nil,"krishna",i,j)
			fmt.Println("presc: ",a,b,c,d,e,f,g,h)
		}
	}

//	name, _ = token.Name(nil)
//	fmt.Println("Post-mining name:", name)
//
//	name, _ = token.Name(&bind.CallOpts{Pending: true})
//	fmt.Println("Post-mining pending name:", name)
	exists,_ := token.IsExistingUser(nil,"doc1");
	fmt.Println("doc1 exists? :", exists)
	exists,_ = token.IsExistingUser(nil,"doc2");
	fmt.Println("doc2 exists? :", exists)

}
