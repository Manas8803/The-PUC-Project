"use client";
import ProtectedRoute from "@/components/auth/ProtectedRoute";
import MainLayout from "@/components/ui/MainLayout";
import { useAuth } from "@/hooks";
import { LogOut } from "lucide-react";
import { useState } from "react";
import Reports from "./components/reports";
import SearchBar from "./components/search_bar";

export default function Home() {
	const [searchQuery, setSearchQuery] = useState("");
	const { logout } = useAuth();

	const handleSearch = (query: string) => {
		setSearchQuery(query);
	};

	function handleLogout() {
		logout();
	}
	return (
		<ProtectedRoute>
			<MainLayout>
				<>
					<div className="fixed top-14 left-0 right-0 bg-bgrnd">
						<div className="fixed flex top-5 left-72 right-0 bg-bgrnd text-main gap-2">
							<LogOut />
							<button className="font-regular" onClick={handleLogout}>
								Sign out
							</button>
						</div>
						<SearchBar onSearch={handleSearch} />
					</div>
					<div className="mt-32 flex-grow overflow-y-auto">
						<Reports searchQuery={searchQuery} />
					</div>
				</>
			</MainLayout>
		</ProtectedRoute>
	);
}
