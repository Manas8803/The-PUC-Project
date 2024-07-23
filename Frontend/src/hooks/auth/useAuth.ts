"use client";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function useAuth() {
	const baseUrl = process.env.NEXT_PUBLIC_AUTH_URL;
	const router = useRouter();
	const [isLoading, setIsLoading] = useState(false);

	async function login(email: string, password: string) {
		setIsLoading(true);

		if (email === "") {
			setIsLoading(false);
			throw new Error("Please enter your email address.");
		}

		if (password === "") {
			setIsLoading(false);
			throw new Error("Please enter your password.");
		}

		try {
			const response = await fetch(`${baseUrl}/api/v1/auth/login`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ email, password }),
			});

			switch (response.status) {
				case 200:
					const jsonData = await response.json();
					const token = jsonData.data.token.split("Bearer ")[1];
					const officeName = jsonData.data.office_name;
					localStorage.setItem("officeName", officeName);
					localStorage.setItem("token", token);
					break;
				case 401:
					throw new Error(
						"Invalid credentials. Please check your email and password."
					);
				case 404:
					throw new Error("User not found. Please verify the entered email.");
				default:
					throw new Error(
						"Oops! Something went wrong. Please try again later."
					);
			}
		} catch (error: unknown) {
			throw error;
		} finally {
			setIsLoading(false);
		}
	}
	function logout() {
		localStorage.removeItem("token");
		localStorage.removeItem("officeName");
		router.push("/auth/login");
	}

	return { isLoading, login, logout };
}
