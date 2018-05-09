pragma solidity ^0.4.16;

contract mortal {
    /* Define variable owner of the type address */
    address owner;

    /* This function is executed at initialization and sets the owner of the contract */
    function mortal() public { owner = msg.sender;  }

    /* Function to recover the funds on the contract */
    function kill() public { if (msg.sender == owner) selfdestruct(owner);  }

}

contract medcon is mortal {

	function medcon() public {
		
	}

	enum UserRole {Doctor, Patient, Admin, Pharma}

	struct User {
		string name;
		string username;
		//should be a hash
		uint64 password;
		UserRole role;
		bool initialized;
	}

	struct MedOrder {
		string med_name;
		string diagnosis;
		uint64 dose_per_day;
		uint64 no_of_days;
		string doctor;
		bool initialized;
		// if pharma has dispensed this order, mark as fulfilled
		bool fulfilled;
		string pharma;
	}

	struct Patient {
		mapping(uint64=>Prescription) prescriptions;
		mapping(string=>uint64) doctor_index_map;
		uint64 num_of_prescriptions;
		bool initialized;
	}
	
	struct Prescription {
		mapping(uint64=>MedOrder) med_orders;
		uint64 num_of_med_orders;
		string doctor;
		bool initialized;
	}

	struct Doctor {
		string medical_id;
		string[] patients;
		bool initialized;
	}
	
	struct Pharma {
		string pharma_id;
		string[] patients;
		bool initialized;
	}

	mapping(string => Doctor) internal doctors;
	mapping(string => Patient) internal patients;
	mapping(string => User) internal users;
	mapping(string => Pharma) internal pharmas;

	function isExistingUser(string username) public view returns (bool res) {
		return users[username].initialized;
	}

	function getUser(string username, string name, uint64 password, UserRole role) internal pure returns (User usr) {
		User memory u;
		u.name = name;
		u.username = username;
		u.password = password;
		u.role = role;
		u.initialized = true;

		return u;
	}

	function getDoctor(string medical_id) internal pure returns (Doctor doc) {
		Doctor memory d;
		d.medical_id = medical_id;
		d.initialized = true;
		return d;
	}	
	
	function getPharma(string pharma_id) internal pure returns (Pharma ph) {
		Pharma memory d;
		d.pharma_id = pharma_id;
		d.initialized = true;
		return d;
	}	
	
	function getPatient() internal pure returns (Patient pt) {
		Patient memory p;
		p.initialized = true;
		p.num_of_prescriptions = 0;
		return p;
	}

	function getMedOrder(string med_name, uint64 dose_per_day, uint64 no_of_days, string doctor, string diagnosis) internal pure returns (MedOrder medo){
		MedOrder memory mo;
		mo.med_name = med_name;
		mo.dose_per_day = dose_per_day;
		mo.no_of_days = no_of_days;
		mo.doctor = doctor;
		mo.diagnosis = diagnosis;
		mo.initialized = true;

		return mo;
	}

	function getPrescription(string doctor) internal pure returns (Prescription pr) {
		Prescription memory p;
		p.doctor = doctor;
		p.num_of_med_orders = 0;
		p.initialized = true;

		return p;
	}

	function addMedOrderToPatient(string patient, string med_name, uint64 dose_per_day, uint64 no_of_days, string doctor, string diagnosis) public {
		MedOrder memory md = getMedOrder(med_name, dose_per_day, no_of_days, doctor, diagnosis);
		Patient storage p = patients[patient];
		if (p.doctor_index_map[doctor] == 0) {
			p.doctor_index_map[doctor] = ++p.num_of_prescriptions;
		}

		if (!p.prescriptions[p.doctor_index_map[doctor]].initialized) {
			p.prescriptions[p.doctor_index_map[doctor]] = getPrescription(doctor);
		}

		Prescription storage pr = p.prescriptions[p.doctor_index_map[doctor]];
		
		pr.med_orders[pr.num_of_med_orders++] = md;

	}

	function addDoctor(string username, string name, uint64 password, string medical_id) public {
		if (users[username].initialized == true) {
			//user already exists
			revert();
		}
		
		users[username] = getUser(name, username, password, UserRole.Doctor);
		doctors[username] = getDoctor(medical_id);
	}
	
	function addPatient(string username, string name, uint64 password) public {
		if (users[username].initialized == true) {
			//user already exists
			revert();
		}
		users[username] = getUser(name, username, password, UserRole.Patient);
		patients[username] = getPatient();
	}
	
	function addPharma(string username, string name, uint64 password, string pharma_id) public {
		if (users[username].initialized == true) {
			//user already exists
			revert();
		}
		users[username] = getUser(name, username, password, UserRole.Pharma);
		pharmas[username] = getPharma(pharma_id);
	}
///////////////////////////////////////////////////////////////////////////////////////////

	function showNumOfMedOrdersByDoc(string patient, string doctor) public view returns (uint64) {
		Patient storage p = patients[patient];
		Prescription storage pr = p.prescriptions[p.doctor_index_map[doctor]];
		return pr.num_of_med_orders;
	}
	
	function showNumOfMedOrdersByIndex(string patient, uint64 indx) public view returns (uint64) {
		Patient storage p = patients[patient];
		if (indx > p.num_of_prescriptions) {
			revert();
		}
		Prescription storage pr = p.prescriptions[indx];
		return pr.num_of_med_orders;
	}
	
	function showNumOfPrescriptions(string patient) public view returns (uint64) {
		Patient storage p = patients[patient];
		return p.num_of_prescriptions;
	}

	function showMedOrderByDoctor(string patient, string doctor, uint64 index) public view returns (string, string, uint64, uint64, string, bool, string) {
		Patient storage p = patients[patient];
		Prescription storage pr = p.prescriptions[p.doctor_index_map[doctor]];
		if (index >= pr.num_of_med_orders){
			revert();
		}
		MedOrder storage m = pr.med_orders[index];

		return (m.med_name, m.diagnosis, m.dose_per_day, m.no_of_days, m.doctor, m.fulfilled, m.pharma); 
	}
	
	function showMedOrderByIndex(string patient, uint64 pres, uint64 index) public view returns (string, string, uint64, uint64, string, bool, string) {
		Patient storage p = patients[patient];
		if (!p.initialized) {
			revert();
		}
		if (p.num_of_prescriptions > pres) {
			revert();
		}
		Prescription storage pr = p.prescriptions[pres];
		if (!pr.initialized) {
			revert();
		}
		if (index >= pr.num_of_med_orders){
			revert();
		}
		MedOrder storage m = pr.med_orders[index];

		return (m.med_name, m.diagnosis, m.dose_per_day, m.no_of_days, m.doctor, m.fulfilled, m.pharma); 
	}

}





