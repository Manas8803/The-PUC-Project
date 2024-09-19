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
					<div className="flex flex-col min-h-screen bg-bgrnd">
						<header className="sticky top-0 z-10 bg-bgrnd">
							<div className="container mx-auto px-4 py-4 flex justify-end">
								<button
									className="flex items-center gap-2 text-main bg-white px-4 py-2 rounded-[0.75rem] shadow-md"
									onClick={handleLogout}
								>
									<LogOut />
									<span className="font-regular">Sign out</span>
								</button>
							</div>
							<div className="container mx-auto px-4 pb-4">
								<SearchBar onSearch={handleSearch} />
							</div>
						</header>
						<main className="flex-grow overflow-y-auto container mx-auto px-4 py-8">
							<Reports searchQuery={searchQuery} />
						</main>
					</div>
				</>
			</MainLayout>
		</ProtectedRoute>
	);
}
