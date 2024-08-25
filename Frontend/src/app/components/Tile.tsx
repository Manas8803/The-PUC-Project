import { LucideIcon } from "lucide-react";

interface TileProps {
	icon: React.ReactElement<LucideIcon>;
	title: string;
	subtitle: string;
}

export default function Tile({ icon, title, subtitle }: TileProps) {
	return (
		<div className="bg-white px-[1.25rem] py-[1rem] rounded-[1.25rem] shadow-sm flex flex-col items-start">
			<div className="flex flex-row">
				<div className="text-main">{icon}</div>
				<h3 className="text-[0.875rem] font-regular mb-1 ml-2">{title}</h3>
			</div>
			<p className="text-[0.75rem] font-light text-gray_500 text-wrap">
				{subtitle}
			</p>
		</div>
	);
}
