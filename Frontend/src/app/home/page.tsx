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
			<div className="h-screen">
				<MainLayout>
					<div className="flex flex-col h-full">
						<div className="fixed top-0 left-0 right-0 z-10 bg-bgrnd">
							<SearchBar onSearch={handleSearch} />
						</div>
						<div className="mt-32 flex-grow overflow-y-auto">
							<Reports searchQuery={searchQuery} />
						</div>
					</div>
				</MainLayout>
			</div>
		</ProtectedRoute>
	);
}
