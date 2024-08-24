import React from "react";
import Link from "next/link";
import Image from "next/image";
import { Camera, Car, AlertTriangle, Wifi } from "lucide-react";
import placeholder from "@/public/Placeholder.png";
import individuals from "@/public/icons/individuals.svg";
import authorities from "@/public/icons/authorities.svg";

export default function Landing() {
	return (
		<main className="max-w-md mx-auto bg-white shadow-lg rounded-lg overflow-hidden m-2 p-2">
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

				<h3 className="font-semibold mb-6">Key Features</h3>
				<div className="grid grid-cols-2 gap-4 mb-6">
					{[
						{ icon: <Car size={44} />, text: "Detect vehicle number" },
						{ icon: <Camera size={24} />, text: "Validate PUC certificate" },
						{ icon: <AlertTriangle size={24} />, text: "Monitor pollution" },
						{ icon: <Wifi size={24} />, text: "Alert system" },
					].map((feature, index) => (
						<div
							key={index}
							className="flex items-center p-6 rounded border-4 border-gray_100"
						>
							{feature.icon}
							<span className="ml-2 text-md text-pretty">{feature.text}</span>
						</div>
					))}
				</div>

				<h3 className="font-semibold mb-6">Who is this for?</h3>
				<div className="grid grid-cols-2 gap-4 mb-6">
					<div className="flex p-6 border-4 text-right text-pretty border-gray_100 rounded">
						<Image
							src={individuals}
							width={24}
							height={24}
							alt="individuals-icon"
						></Image>
						<p className="font-regular ml-3">Individuals</p>
					</div>
					<div className="flex p-6 border-4 text-right text-pretty border-gray_100 rounded">
						<Image
							src={authorities}
							width={24}
							height={24}
							alt="authorities-icon"
						></Image>
						<p className="font-regular ml-3">Authorities</p>
					</div>
				</div>

				<h3 className="font-semibold mb-6">About the project</h3>
				<p className="text-sm text-gray_600 mb-6">
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
