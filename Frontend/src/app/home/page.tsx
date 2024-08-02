"use client";
import ProtectedRoute from "@/components/auth/ProtectedRoute";
import MainLayout from "@/components/ui/MainLayout";
import { useState } from "react";
import Reports from "./components/reports";
import SearchBar from "./components/search_bar";

export default function Home() {
	const [searchQuery, setSearchQuery] = useState("");

	const handleSearch = (query: string) => {
		setSearchQuery(query);
	};
	return (
		<ProtectedRoute>
			<div className="bg-bgrnd min-h-[100vh]">
				<MainLayout>
					<>
						<SearchBar onSearch={handleSearch} />
						<Reports searchQuery={searchQuery} />
					</>
				</MainLayout>
			</div>
		</ProtectedRoute>
	);
}
