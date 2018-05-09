package main

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"html/template"
	"math/big"
//	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/sgr159/MedNet/medcon"
)


type UserSignupReq struct {
	username string
	name string
	password string
	role string
	id string
}

type MedOrder struct {
	MedName string
	Diagnosis string
	DosePerDay uint64
	NoOfDays uint64
	Doctor string
	Fulfilled bool
	Pharma string
}

type Prescription struct {
	Doctor string
	MedOrders []MedOrder
}

type UserData struct {
	UserName string
	Name string
	Prescriptions []Prescription
}

type PageVariables struct {
	PageTitle string
}

type UserPageVariables struct {
	PageTitle string
	UserMedData UserData
}

var token * medcon.Medcon
var auth * bind.TransactOpts
var sim * backends.SimulatedBackend

func main() {
	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	auth = bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}

	sim = backends.NewSimulatedBackend(alloc)

	// Deploy a token contract on the simulated blockchain
	_, _, token, _ = medcon.DeployMedcon(auth, sim)
	fmt.Println(token)
	token.AddDoctor(auth,"doc1","doc1",123,"id123");
	token.AddDoctor(auth,"doc2","doc2",123,"id123");
	token.AddPatient(auth,"krishna","krishna",123);
	token.AddPharma(auth,"patanjali","patanjali",123,"sad");

	sim.Commit()

	token.AddMedOrderToPatient(auth,"krishna","saridon",2,7,"doc1","headache")
	token.AddMedOrderToPatient(auth,"krishna","pcm",2,7,"doc1","fever")
	token.AddMedOrderToPatient(auth,"krishna","weed",2,7,"doc2","depression")
	sim.Commit()

	getUserData("krishna")

	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/login", Login).Methods("GET")
	router.HandleFunc("/patient", Patient).Methods("GET")
//	router.HandleFunc("/loginReq", LoginReq).Methods("POST")
	router.HandleFunc("/signup", Signup).Methods("GET")
	router.HandleFunc("/signupReq", SignupReq).Methods("POST")

	fmt.Println("Server deployed")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {

	Title := "SAMPLE PAGE"

	MyPageVariables := PageVariables{Title}

	parsePage("index.html",MyPageVariables,w)

	fmt.Println("Index: Get request Recieved")
}

func Patient (w http.ResponseWriter, r *http.Request) {

	Title := "Login"

	MyPageVariables := UserPageVariables{Title,getUserData("krishna")}

	parsePatientPage("patient_home.html",MyPageVariables,w)

	fmt.Println("Patient: Get request Recieved")
}

func Login(w http.ResponseWriter, r *http.Request) {

	Title := "Login"

	MyPageVariables := PageVariables{Title}

	parsePage("Login.html",MyPageVariables,w)

	fmt.Println("Login: Get request Recieved")
}

func LoginReq(w http.ResponseWriter, r *http.Request) {
}

func Signup(w http.ResponseWriter, r *http.Request) {

	Title := "Signup"

	MyPageVariables := PageVariables{Title}

	parsePage("Signup.html",MyPageVariables,w)

	fmt.Println("Signup: Get request Recieved")
}

func SignupReq(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signupreq: POST request Recieved")
	if r.ContentLength == 0 {
		fmt.Println("Empty request")
		return
	}
	r.ParseForm()
	fmt.Println(r)
	signupReq := UserSignupReq{r.Form.Get("userName"), r.Form.Get("name"), r.Form.Get("password"), r.Form.Get("role"), r.Form.Get("medId")}
	fmt.Println("harcode:",r.Form.Get("userName"), r.Form.Get("name"), r.Form.Get("password"), r.Form.Get("role"), r.Form.Get("medId"))
	AddUser(signupReq)
	fmt.Println("request: ", signupReq)
}

func AddUser(req UserSignupReq) {
	fmt.Println("Add User request type: ", req.role)
	pwint,_ := strconv.Atoi(req.password)
	pw := uint64(pwint)
	fmt.Println("Addu: token",token)
	switch req.role {
	case "1":
		_,err := token.AddPatient(auth,req.username,req.name,pw);
		if err != nil {
			fmt.Println("AddUser:",err)
		}
	case "2":
		_,err := token.AddDoctor(auth,req.username,req.name,pw,req.id);
		if err != nil {
			fmt.Println("AddUser:",err)
		}
	case "3":
		_,err := token.AddPharma(auth,req.username,req.name,pw,req.id);
		if err != nil {
			fmt.Println("AddUser:",err)
		}
	}


	sim.Commit()

	exists,_ := token.IsExistingUser(nil,req.username);
	fmt.Println("user exists on blockchain? :", exists)
}

func parsePage(htmlFile string,MyPageVariables PageVariables, w http.ResponseWriter) {
	t, err := template.ParseFiles(htmlFile) //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func parsePatientPage(htmlFile string,MyPageVariables UserPageVariables, w http.ResponseWriter) {
	t, err := template.ParseFiles(htmlFile) //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func getUserData (user string) (UserData) {
	num_pres,_ := token.ShowNumOfPrescriptions(nil,user)
	fmt.Println("num of prescriptions:",num_pres)
	var usrData UserData;
	usrData.UserName, usrData.Name, _ = token.ShowPatientDetails(nil,user)
	for i:=uint64(1);i<=num_pres;i++ {
		num_med,_ := token.ShowNumOfMedOrdersByIndex(nil,user, i)
		var mos []MedOrder
		var doctor string
		for j:=uint64(0);j<num_med;j++ {
			a,b,c,d,e,f,g,err := token.ShowMedOrderByIndex(nil,user,i,j)
			mos = append(mos, MedOrder{a,b,c,d,e,f,g})
			fmt.Println("presc: ",a,b,c,d,e,f,g,err)
			doctor = e
		}
		usrData.Prescriptions = append(usrData.Prescriptions,Prescription{doctor,mos})
	}
	fmt.Println("GetUserData: ", usrData)
	return usrData
}
