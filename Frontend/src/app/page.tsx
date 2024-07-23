"use client";
import { useAuth } from "@/hooks/index";
import Link from "next/link";

export default function Home() {
	const { logout } = useAuth(); 

	return (
		<main className="pl-10">
			Landing Page
			<br />
			<br />
			<Link href={"/auth/login"}>Login</Link>
			<br />
			<br />
			<button onClick={logout}>Logout</button>
		</main>
	);
}
