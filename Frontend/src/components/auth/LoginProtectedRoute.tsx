"use client";
import { ReactElement, useEffect } from "react";

import { useRouter } from "next/navigation";
import { toast } from "../ui/use-toast";

export default function LoginProtectedRoute({
	children,
}: {
	children: ReactElement;
}) {
	const router = useRouter();
	useEffect(() => {
		const token = localStorage.getItem("token");

		if (token) {
			toast({
				variant: "normal",
				title: "Already logged in",
			});
			router.back();
			return;
		}
	}, [router]);

	return children;
}
