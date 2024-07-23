"use client";
import { CardData, cardData } from "@/lib/data";
import { useEffect, useState } from "react";

export const useReportsData = () => {
	const [reportsData, setReportsData] = useState<CardData[]>(cardData);
	const isBrowser = typeof window !== "undefined";
	useEffect(() => {
		if (!isBrowser) {
			return;
		}
		let socket: WebSocket | null = null;

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
					socket.addEventListener("error", (error) => {
						// console.error("WebSocket error:", error.currentTarget);
					});
				}
			} catch (error) {}
		};

		connectWebSocket();

		return () => {
			if (socket) {
				socket.close();
			}
		};
	}, [isBrowser]);

	return { reportsData };
};
