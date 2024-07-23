import { ReactElement } from "react";
import MenuBar from "./menu_bar";
import { motion } from "framer-motion";
export default function MainLayout({ children }: { children: ReactElement }) {
	return (
		<>
			<main className="flex flex-col bg-bgrnd relative pb-10 mb-[16svh] h-[100%]">
				{children}
				<MenuBar />
			</main>
		</>
	);
}
