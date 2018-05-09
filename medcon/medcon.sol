pragma solidity ^0.4.16;

contract medcon {

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
		uint med_code;
		uint dose_per_day;
		uint no_of_days;
		string doctor;
		bool initialized;
		// if pharma has dispensed this order, mark as fulfilled
		bool fulfilled;
		string pharma;
	}

	struct Patient {
		MedOrder[] med_orders;
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

	mapping(string => Doctor) public doctors;
	mapping(string => Patient) public patients;
	mapping(string => User) public users;
	mapping(string => Pharma) public pharmas;

	function isExistingUser(string username) public returns {
		return users[username].initialized;
	}

	function getUser(string username, string name, uint password, UserRole role) public returns (User u) {
		User memory u;
		u.name = name;
		u.username = username;
		u.password = password;
		u.role = role;
		u.initialized = true;

		return u;
	}

	function getDoctor(string medical_id) public returns (Doctor d) {
		Doctor memory d;
		d.medical_id = medical_id;
		d.initialized = true;
		return d
	}	
	
	function getPharma(string pharma_id) public returns (Pharma d) {
		Pharma memory d;
		d.pharma_id = pharma_id;
		d.initialized = true;
		return d
	}	
	
	function getPatient() public returns (Patient d) {
		Patient memory p;
		p.initialized = true;
		return p
	}	

	function addDoctor(string username, string name, uint password, string medical_id) public {
		if (users[username].initialized == true) {
			//user already exists
			revert();
		}
		
		users[username] = getUser(name, username, password, UserRole.Doctor);
		doctors[username] = Doctor(medical_id, 0, true);
	}
	
	function addPatient(string username, string name, string password) public {
		if (users[username].initialized == true) {
			//user already exists
			revert();
		}
		users[username] = getUser(name, username, password, UserRole.Patient);
		patients[username] = getPatient();
	}
	
	function addPharma(string username, string name, string password, string pharma_id) public {
		if (users[username].initialized == true) {
			//user already exists
			revert();
		}
		users[username] = getUser(name, username, password, UserRole.Pharma);
		pharmas[username] = getPharma(pharma_id);
	}

} 
