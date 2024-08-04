import { ReactElement } from "react";
import MenuBar from "./menu_bar";
export default function MainLayout({ children }: { children: ReactElement }) {
	return (
		<main className="flex flex-col relative pb-10">
			{children}
			<MenuBar />
		</main>
	);
}
