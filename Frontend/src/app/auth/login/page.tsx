"use client";

import LoginProtectedRoute from "@/components/auth/LoginProtectedRoute";
import { useToast } from "@/components/ui/use-toast";
import { useAuth } from "@/hooks";
import logo from "@/public/logo.webp";
import Image from "next/image";
import { useRouter } from "next/navigation";
import React, { useState } from "react";

export default function Login() {
	//* States
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");

	//* Router
	const router = useRouter();

	//* Hooks
	const { isLoading, login } = useAuth();

	//* Toast
	const { toast } = useToast();

	const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();
		try {
			await login(email, password);
			router.push("/home");
		} catch (error: unknown) {
			if (error instanceof Error) {
				toast({
					variant: "destructive",
					title: "Action Failed!",
					className: "font-mono tracking-wide",
					description: error.message,
				});
			} else {
				toast({
					variant: "destructive",
					title: "Something went wrong!",
					description: "An unexpected error occurred",
				});
			}
		}
	};

	return (
		<LoginProtectedRoute>
			<div className="flex justify-center items-center min-h-screen bg-gray-200 min-w-[100%] overflow-hidden">
				<div className="mx-auto bg-white p-8 shadow-2xl rounded-2xl border-black border-[0.5px] sm:min-w-[60%] lg:max-w-[60%]">
					<div className="mb-6">
						<Image
							src={logo}
							width={139}
							height={90}
							alt="The PUC Project logo"
							className="mx-auto"
						/>
					</div>
					<form onSubmit={handleSubmit}>
						<div className="mb-4">
							<label
								htmlFor="Email"
								className="block text-gray-700 font-bold mb-2"
							>
								Email
							</label>
							<input
								type="email"
								value={email}
								onChange={(e) => setEmail(e.target.value)}
								placeholder="you@example.com"
								className="shadow appearance-none border rounded-2xl w-full py-3 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
								required
							/>
						</div>
						<div className="mb-6">
							<label
								htmlFor="password"
								className="block text-gray-700 font-bold mb-2"
							>
								Password
							</label>
							<input
								type="password"
								value={password}
								placeholder="•••••••••••"
								onChange={(e) => setPassword(e.target.value)}
								className="shadow appearance-none border rounded-2xl w-full py-3 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
								required
							/>
						</div>
						<div className="flex justify-center">
							<button
								type="submit"
								disabled={isLoading}
								className="bg-main hover:bg-side text-white font-bold border rounded-2xl w-full py-3 px-4 focus:outline-none focus:shadow-outline"
							>
								{isLoading ? "Logging in..." : "Login"}
							</button>
						</div>
					</form>
				</div>
			</div>
		</LoginProtectedRoute>
	);
}
