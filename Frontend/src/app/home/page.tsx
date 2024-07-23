"use client";
import ProtectedRoute from "@/components/auth/ProtectedRoute";
import MainLayout from "@/components/ui/MainLayout";
import { useState } from "react";
import Reports from "./components/puc_alert_section";
import SearchBar from "./components/search_bar";

export default function Home() {
	const [searchQuery, setSearchQuery] = useState("");

	const handleSearch = (query: string) => {
		setSearchQuery(query);
	};
	return (
		<ProtectedRoute>
			<MainLayout>
				<>
					<SearchBar onSearch={handleSearch} />
					<Reports searchQuery={searchQuery} />
				</>
			</MainLayout>
		</ProtectedRoute>
	);
}
