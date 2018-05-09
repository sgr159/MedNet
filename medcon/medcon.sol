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

	function medcon() {
		
	}

	enum UserRole {Doctor, Patient, Admin, Pharma}

	struct User {
		string name;
		string username;
		//should be a hash
		uint password;
		UserRole role;
		bool initialized;
	}

	struct MedOrder {
		string med_name;
		string diagnosis;
		uint dose_per_day;
		uint no_of_days;
		string doctor;
		bool initialized;
		// if pharma has dispensed this order, mark as fulfilled
		bool fulfilled;
		string pharma;
	}

	struct Patient {
		mapping(uint=>Prescription) prescriptions;
		mapping(string=>uint) doctor_index_map;
		uint num_of_prescriptions;
		bool initialized;
	}
	
	struct Prescription {
		mapping(uint=>MedOrder) med_orders;
		uint num_of_med_orders;
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

	function getUser(string username, string name, uint password, UserRole role) internal returns (User usr) {
		User memory u;
		u.name = name;
		u.username = username;
		u.password = password;
		u.role = role;
		u.initialized = true;

		return u;
	}

	function getDoctor(string medical_id) internal returns (Doctor doc) {
		Doctor memory d;
		d.medical_id = medical_id;
		d.initialized = true;
		return d;
	}	
	
	function getPharma(string pharma_id) internal returns (Pharma ph) {
		Pharma memory d;
		d.pharma_id = pharma_id;
		d.initialized = true;
		return d;
	}	
	
	function getPatient() internal returns (Patient pt) {
		Patient memory p;
		p.initialized = true;
		p.num_of_prescriptions = 0;
		return p;
	}

	function getMedOrder(string med_name, uint dose_per_day, uint no_of_days, string doctor, string diagnosis) internal returns (MedOrder medo){
		MedOrder memory mo;
		mo.med_name = med_name;
		mo.dose_per_day = dose_per_day;
		mo.no_of_days = no_of_days;
		mo.doctor = doctor;
		mo.diagnosis = diagnosis;
		mo.initialized = true;

		return mo;
	}

	function getPrescription(string doctor) internal returns (Prescription pr) {
		Prescription memory p;
		p.doctor = doctor;
		p.num_of_med_orders = 0;
		p.initialized = true;

		return p;
	}

	function addMedOrderToPrescription(Prescription p, MedOrder m) internal {
		p.med_orders[p.num_of_med_orders++] = m;
	}
	
	function addMedOrderToPatient(string patient, string med_name, uint dose_per_day, uint no_of_days, string doctor) public {
		MedOrder memory md = getMedOrder(med_name, dose_per_day, no_of_days, doctor);
		Patient memory p = patients[patient];
		if (p.doctor_index_map[doctor] == 0) {
			p.doctor_index_map[doctor] = ++p.num_of_prescriptions;
		}

		if (!p.prescriptions[p.doctor_index_map[doctor]].initialized) {
			p.prescriptions[p.doctor_index_map[doctor]] = getPrescription(doctor);
		}

		addMedOrderToPrescription(p.prescriptions[p.doctor_index_map[doctor]], md);
	}

	function addDoctor(string username, string name, uint password, string medical_id) public {
		if (users[username].initialized == true) {
			//user already exists
			revert();
		}
		
		users[username] = getUser(name, username, password, UserRole.Doctor);
		doctors[username] = getDoctor(medical_id);
	}
	
	function addPatient(string username, string name, uint password) public {
		if (users[username].initialized == true) {
			//user already exists
			revert();
		}
		users[username] = getUser(name, username, password, UserRole.Patient);
		patients[username] = getPatient();
	}
	
	function addPharma(string username, string name, uint password, string pharma_id) public {
		if (users[username].initialized == true) {
			//user already exists
			revert();
		}
		users[username] = getUser(name, username, password, UserRole.Pharma);
		pharmas[username] = getPharma(pharma_id);
	}

} 
