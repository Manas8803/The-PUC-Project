import { CardData } from "@/lib/data";
import dropdown_icon from "@/public/home/alert-card/dropdown-icon.webp";
import indicator_green from "@/public/home/alert-card/indicator-green.webp";
import indicator_red from "@/public/home/alert-card/indicator-red.webp";
import { m } from "framer-motion";
import Image from "next/image";
import { useState } from "react";
export default function AlertCard({
	puc_status,
	last_check_date,
	vehicle_class_desc,
	model,
	reg_no,
	vehicle_type,
	office_name,
	owner_name,
	puc_upto,
	mobile,
	reg_upto,
}: CardData) {
	const [isOpen, setIsOpen] = useState(false);

	return (
		<m.div
			className={`mt-3 w-[83.333%] mx-auto ${
				puc_status
					? "bg-gradient-to-r from-white to-main"
					: "bg-gradient-to-r from-gr_white to-gr_red"
			} rounded-3xl p-4 shadow-lg text-black duration-600 transition-display ease-in-out transform`}
			whileTap={{ scale: 0.93 }}
			whileHover={{ scale: 1.03 }}
			whileInView={{
				x: [-20, 0],
				transition: { ease: "easeInOut", duration: 0.2 },
			}}
			onClick={() => setIsOpen(!isOpen)}
		>
			<div className="flex items-center justify-between mb-4">
				<div className="flex items-center gap-2">
					<Image
						src={puc_status ? indicator_green : indicator_red}
						width={10}
						height={10}
						alt="indicator"
						loading="lazy"
					/>
					<p>{puc_status ? "PUC is valid" : "Seems the PUC is outdated!"}</p>
				</div>
				<div
					className={`cursor-pointer transform transition duration-500 ease-in-out rounded-lg ${
						isOpen ? "rotate-180" : ""
					}`}
					onClick={() => setIsOpen(!isOpen)}
				>
					<Image
						src={dropdown_icon}
						alt="dropdown"
						width={20}
						height={20}
						loading="lazy"
					/>
				</div>
			</div>
			<h1 className="text-xl font-bold pl-2">{reg_no}</h1>
			<p className="font-normal pl-2 mb-2">{model}</p>
			<div
				className={`overflow-hidden transition-all ease-in-out delay-0 ${
					isOpen ? "max-h-screen duration-700" : "max-h-0 duration-400"
				}`}
			>
				<div className="pl-2 font-extralight">
					<p>
						Registration No.: &nbsp;<strong>{reg_no}</strong>
					</p>
					<p>
						Vehicle Model: &nbsp;<strong>{model}</strong>
					</p>
					<p>
						Vehicle Description: &nbsp;<strong>{vehicle_class_desc}</strong>
					</p>
					<p>
						Vehicle Type: &nbsp;<strong>{vehicle_type}</strong>
					</p>
					<p>
						Owner Name: &nbsp;<strong>{owner_name}</strong>
					</p>
					<p>
						RTO registered office: &nbsp;<strong>{office_name}</strong>
					</p>
					<p>
						Contact: &nbsp;<strong>{mobile}</strong>
					</p>
					<p>
						Puc valid upto: &nbsp;
						<strong>
							{puc_upto?.day + "/" + puc_upto?.month + "/" + puc_upto?.year}
						</strong>
					</p>
					<p>
						Registration valid upto: &nbsp;
						<strong>
							{reg_upto?.day + "/" + reg_upto?.month + "/" + reg_upto?.year}
						</strong>
					</p>
					<p>
						Last check date: &nbsp;
						<strong>
							{last_check_date?.day +
								"/" +
								last_check_date?.month +
								"/" +
								last_check_date?.year}
						</strong>
					</p>
				</div>
			</div>
		</m.div>
	);
}
