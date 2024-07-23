"use client";

import { LazyMotion, domAnimation, m } from "framer-motion";
export default function Transition({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<LazyMotion features={domAnimation} strict>
			<m.div
				initial={{ y: 20, opacity: 0 }}
				animate={{ y: 0, opacity: 1 }}
				transition={{ ease: "easeInOut", duration: 0.4 }}
			>
				{children}
			</m.div>
		</LazyMotion>
	);
}
