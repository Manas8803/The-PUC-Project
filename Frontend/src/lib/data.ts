export interface CardData {
	pucStatus: string;
	vehicleType: string;
	validUpto: string;
	registrationNo: string;
	vehicleModel: string;
	vehicleDescription: string;
	contact: string;
	pucValidUpto: string;
	office_name : string;
}

export const cardData: CardData[] = [
	{
		office_name:"Mumbai",
		pucStatus: "Valid",
		vehicleType: "Car",
		validUpto: "31 Nov 2023",
		registrationNo: "ILKPK14703",
		vehicleModel: "Super Splendor",
		vehicleDescription: "Bike",
		contact: "+91 9741053920",
		pucValidUpto: "31 Feb 2025",
	},
	{
		office_name:"Mumbai",
		pucStatus: "Valid",
		vehicleType: "Truck",
		validUpto: "31 Dec 2024",
		registrationNo: "ABC123",
		vehicleModel: "Volvo",
		vehicleDescription: "Heavy Vehicle",
		contact: "+91 9876543210",
		pucValidUpto: "31 Mar 2025",
	},
	{
		office_name:"Mumbai",
		pucStatus: "Expired",
		vehicleType: "Motorcycle",
		validUpto: "31 Jan 2023",
		registrationNo: "XYZ789",
		vehicleModel: "Honda",
		vehicleDescription: "Two-wheeler",
		contact: "+91 9998887776",
		pucValidUpto: "31 Jan 2022",
	},
];
