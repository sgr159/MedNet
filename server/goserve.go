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
	Index uint64
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
	Patients []string
	Doctors []string
	Pharmas []string
	SelectedDoctor string
}

var token * medcon.Medcon
var auth * bind.TransactOpts
var sim * backends.SimulatedBackend

var Patients []string
var Doctors []string
var Pharmas []string

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
	token.AddDoctor(auth,"Dr. Rajesh","Dr. Rajesh",123,"id123");
	Doctors = append(Doctors,"Dr. Rajesh")
	token.AddDoctor(auth,"Dr. Kripa","Dr. Kripa",123,"id123");
	Doctors = append(Doctors,"Dr. Kripa")
	token.AddPatient(auth,"krishna","krishna",123);
	Patients = append(Patients,"krishna")
	token.AddPatient(auth,"Shubham","Shubham",123);
	Patients = append(Patients,"Shubham")
	token.AddPharma(auth,"MercyCorps","MercyCorps",123,"sad");
	Pharmas = append(Pharmas,"MercyCorps")

	sim.Commit()

	token.AddMedOrderToPatient(auth,"krishna","saridon",2,7,"Dr. Rajesh","headache")
	token.AddMedOrderToPatient(auth,"krishna","pcm",2,7,"Dr. Rajesh","fever")
	token.AddMedOrderToPatient(auth,"krishna","ativan",2,7,"Dr. Kripa","depression")
	token.AddMedOrderToPatient(auth,"krishna","vasicool",2,7,"Dr. Kripa","hear constriction")
	token.AddMedOrderToPatient(auth,"Shubham","vomistop",2,7,"Dr. Kripa","vommiting")
	sim.Commit()

	getUserData("krishna")

	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/login", Login).Methods("GET")
	router.HandleFunc("/patient", Patient).Methods("GET")
	router.HandleFunc("/patientReq", Patient).Methods("GET")
	router.HandleFunc("/patientReq", PatientReq).Methods("POST")
	router.HandleFunc("/doctors", Doctor).Methods("GET")
//	router.HandleFunc("/loginReq", LoginReq).Methods("POST")
	router.HandleFunc("/signup", Signup).Methods("GET")
	router.HandleFunc("/signupReq", SignupReq).Methods("POST")
	router.HandleFunc("/pharma", Pharma).Methods("GET")
	router.HandleFunc("/pharmaReq", Pharma).Methods("GET")
	router.HandleFunc("/pharmaReq", PharmaReq).Methods("POST")
	router.HandleFunc("/pharmafill", PharmaFill).Methods("POST")
	router.HandleFunc("/pharmafill", Pharma).Methods("GET")
//	router.HandleFunc("/docAddReq", DocAddReq).Methods("POST")

	fmt.Println("Server deployed")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {

	Title := "SAMPLE PAGE"

	MyPageVariables := PageVariables{Title}

	parsePage("index.html",MyPageVariables,w)

	fmt.Println("Index: Get request Recieved")
}

func Doctor (w http.ResponseWriter, r *http.Request) {

	Title := "Doctor"

	MyPageVariables := UserPageVariables{Title,getUserData("krishna"),Patients,Doctors,Pharmas,"Dr. Rajesh"}

	parsePatientPage("doctor.html",MyPageVariables,w)

	fmt.Println("Doctor: Get request Recieved")

}

func Pharma (w http.ResponseWriter, r *http.Request) {

	Title := "Pharma"

	MyPageVariables := UserPageVariables{Title,getUserData("krishna"),Patients,Doctors,Pharmas,""}

	parsePatientPage("pharma.html",MyPageVariables,w)

	fmt.Println("Pharma: Get request Recieved")

}

func Patient (w http.ResponseWriter, r *http.Request) {

	Title := "Login"

	MyPageVariables := UserPageVariables{Title,getUserData("krishna"),Patients,Doctors,Pharmas,""}

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

func PatientReq(w http.ResponseWriter, r *http.Request) {
	fmt.Println("patient req: POST request Recieved")
	if r.ContentLength == 0 {
		fmt.Println("Empty request")
		return
	}
	r.ParseForm()
	fmt.Println(r)
	user := r.Form.Get("user")
	Title := "Patient"
	MyPageVariables := UserPageVariables{Title,getUserData(user),Patients,Doctors,Pharmas,""}

	parsePatientPage("patient_home.html",MyPageVariables,w)

	fmt.Println("request: ", user)
}

func PharmaFill(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pharma fill: POST request Recieved")
	if r.ContentLength == 0 {
		fmt.Println("Empty request")
		return
	}
	r.ParseForm()
	fmt.Println(r)
	user := r.Form.Get("user")
	Title := "Pharma"
	fulfillUser(user,"MercyCorps")
	MyPageVariables := UserPageVariables{Title,getUserData(user),Patients,Doctors,Pharmas,""}

	parsePatientPage("pharma.html",MyPageVariables,w)

	fmt.Println("request: ", user)
}

func PharmaReq(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pharma req: POST request Recieved")
	if r.ContentLength == 0 {
		fmt.Println("Empty request")
		return
	}
	r.ParseForm()
	fmt.Println(r)
	user := r.Form.Get("user")
	Title := "Pharma"
	MyPageVariables := UserPageVariables{Title,getUserData(user),Patients,Doctors,Pharmas,""}

	parsePatientPage("pharma.html",MyPageVariables,w)

	fmt.Println("request: ", user)
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
			mos = append(mos, MedOrder{a,b,c,d,e,f,g,num_med})
			fmt.Println("presc: ",a,b,c,d,e,f,g,err)
			doctor = e
		}
		usrData.Prescriptions = append(usrData.Prescriptions,Prescription{doctor,mos})
	}
	fmt.Println("GetUserData: ", usrData)
	return usrData
}

func fulfillUser(user string, pharma string) {
	num_pres,_ := token.ShowNumOfPrescriptions(nil,user)
	fmt.Println("num of prescriptions:",num_pres)
	for i:=uint64(1);i<=num_pres;i++ {
		num_med,_ := token.ShowNumOfMedOrdersByIndex(nil,user, i)
		for j:=uint64(0);j<num_med;j++ {
			a,b,c,d,e,f,g,err := token.ShowMedOrderByIndex(nil,user,i,j)
			fmt.Println("presc: ",a,b,c,d,e,f,g,err)
			token.MarkMedOrderAsMet(auth,user,e,pharma,j)
			sim.Commit()
		}
	}
}
