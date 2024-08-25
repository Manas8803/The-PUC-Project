import React from "react";
import Link from "next/link";
import Image from "next/image";
import {
	Camera,
	Car,
	AlertTriangle,
	Wifi,
	FileCheck,
	Activity,
	Bell,
} from "lucide-react";
import placeholder from "@/public/Placeholder.png";
import individuals from "@/public/icons/individuals.svg";
import authorities from "@/public/icons/authorities.svg";
import Tile from "./components/Tile";

export default function Landing() {
	return (
		<main className="max-w-md mx-auto bg-bgrnd shadow-lg rounded-lg overflow-hidden p-2">
			<div className="p-4">
				<div className="flex justify-end items-center mb-4">
					<Link href="/home">
						<button className="text-main font-regular">Sign in</button>
					</Link>
				</div>

				<div className="mb-4">
					<Image
						src={placeholder}
						alt="PUC Project"
						width={400}
						height={200}
						className="w-full h-48 object-cover"
					/>
				</div>

				<h1 className="text-gray_600 mb-4 text-center text-2xl text-pretty py-2 px-4 font-regular">
					A digital solution to vehicle pollution
				</h1>

				<h3 className="font-regular text-center mb-6">What we do?</h3>
				<div className="grid grid-cols-2 gap-2 mb-6">
					<Tile
						icon={<Car size={20} />}
						title="Detection"
						subtitle="Detect vehicle number"
					/>
					<Tile
						icon={<FileCheck size={20} />}
						title="Validation"
						subtitle="PUC Certificates"
					/>
					<Tile
						icon={<Activity size={20} />}
						title="Monitoring"
						subtitle="Pollution monitor"
					/>
					<Tile
						icon={<Bell size={20} />}
						title="Alerts"
						subtitle="Real time alerts"
					/>
				</div>

				<h3 className="font-regular text-center mb-6">Who do we serve?</h3>
				<div className="grid grid-cols-2 gap-4 mb-6">
					<Tile
						icon={
							<Image
								src={individuals}
								width={20}
								height={20}
								alt="individuals-icon"
							/>
						}
						title="Individuals"
						subtitle="People can avail PUC
certificates."
					/>
					<Tile
						icon={
							<Image
								src={authorities}
								width={20}
								height={20}
								alt="authorities-icon"
							/>
						}
						title="Authorities"
						subtitle="Authorities can validate certificates."
					/>
				</div>

				<h3 className="font-regular text-center mb-6">About the project</h3>
				<p className="text-sm text-gray_600 mb-6 bg-white p-6 rounded-[1.25rem]">
					The PUC Project is an initiative to reduce air pollution by ensuring
					that vehicles are regularly checked for emissions and have valid
					Pollution Under Control (PUC) certificates.
				</p>
			</div>

			<footer className="text-center py-2 text-sm text-main">
				© 2024 The PUC Project
			</footer>
		</main>
	);
}
