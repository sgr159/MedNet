pragma solidity ^0.4.16;

contract medcon {

	enum UserRole {Doctor, Patient, Admin, Pharma}

	struct MedOrder {
		string med_name;
		uint med_code;
		uint dose_per_day;
		uint no_of_days;
		string doctor;
		bool initialized;
	}

	struct Patient {
		string name;
		address addr;
		string username;
		uint password;
		MedOrder[] med_orders;
		bool initialized;
	}

	struct Doctor {
		string name;
		string username;
		string medical_id;
		uint password;
		string[] patients;
		bool initialized;
	}

	mapping(string => Doctor) public doctors;
	mapping(string => Patients) public patients;

	

} 
