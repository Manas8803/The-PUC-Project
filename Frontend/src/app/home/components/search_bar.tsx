import search_icon from "@/public/home/searchbar/search-icon.webp";
import Image from "next/image";
import React from "react";

interface SearchBarProps {
	onSearch: (query: string) => void;
}

export default function SearchBar({ onSearch }: SearchBarProps) {
	const handleSearch = (event: React.ChangeEvent<HTMLInputElement>) => {
		onSearch(event.target.value.trim());
	};

	return (
		<div className="p-2 flex items-center justify-center mt-[10%]">
			<div className={"relative min-w-[83.33%]"}>
				<input
					className="bg-white py-3 px-5 w-full focus:outline-none focus:border-[0.2px] focus:border-black border border-gray-300 rounded-2xl shadow-lg"
					placeholder="Search for reports..."
					type="search"
					onChange={handleSearch}
				/>
				<div className="absolute inset-y-5 right-3 flex items-center">
					<Image
						className="w-8 text-gray-400"
						src={search_icon}
						width={50}
						alt="search-icon.svg"
					/>
				</div>
			</div>
		</div>
	);
}
