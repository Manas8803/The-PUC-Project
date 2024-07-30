import { useState, useEffect } from "react";
import { CardData } from "@/lib/data";

const baseUrl = process.env.NEXT_PUBLIC_FETCH_VEHICLES_URL;

export const useReportsData = () => {
	const [reportsData, setReportsData] = useState<CardData[]>([]);
	const [isLoading, setIsLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		let socket: WebSocket | null = null;

		const fetchData = async () => {
			const office_name = localStorage.getItem("officeName");
			setIsLoading(true);
			try {
				const response = await fetch(`${baseUrl}`, {
					method: "POST",
					headers: {
						"Content-Type": "application/json",
					},
					body: JSON.stringify({ office_name: office_name }),
				});
				if (response.ok) {
					const data = await response.json();
					setReportsData(data.vehicles);
				} else {
					setError("Failed to fetch data");
				}
			} catch (error) {
				setError("Error fetching data: " + (error as Error).message);
			} finally {
				setIsLoading(false);
			}
		};

		const connectWebSocket = () => {
			try {
				const socketUrl = process.env.NEXT_PUBLIC_SOCKET_URL;
				const officeName = localStorage.getItem("officeName");
				if (officeName) {
					socket = new WebSocket(`${socketUrl}${officeName}`);
					socket.addEventListener("open", () => {
						console.log("WebSocket connection established");
					});
					socket.addEventListener("message", (event) => {
						console.log("A MESSAGE IS RECEIVED: ", JSON.parse(event.data));
						setReportsData((prevData) => [...prevData, JSON.parse(event.data)]);
					});
					socket.addEventListener("error", (error) => {});
				}
			} catch (error) {}
		};

		fetchData();
		connectWebSocket();

		return () => {
			if (socket) {
				socket.close();
			}
		};
	}, []);

	return { reportsData, isLoading, error };
};
