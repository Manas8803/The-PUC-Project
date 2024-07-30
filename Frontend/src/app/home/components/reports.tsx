import React from "react";
import { useReportsData } from "@/hooks/reports/useReportData";
import { CardData } from "@/lib/data";
import { LazyMotion, domAnimation } from "framer-motion";
import AlertCard from "./alert-card";
import "./loader.css";

interface ReportsProps {
	searchQuery: string;
}

export default function Reports({ searchQuery }: ReportsProps) {
	const { reportsData, isLoading, error } = useReportsData();

	if (isLoading) {
		return (
			<div className="split h-[75vh] flex justify-center items-center">
				<div></div>
			</div>
		);
	}

	if (error) {
		return <div>Error: {error}</div>;
	}

	if (!reportsData || reportsData.length === 0) {
		return (
			<div className="flex justify-center items-center h-[80vh] text-2xl text-side">
				No reports available...
			</div>
		);
	}

	const filteredData = reportsData.filter((report: Partial<CardData>) => {
		const searchFields = [
			report.office_name,
			report.puc_status,
			report.vehicle_type,
			report.puc_upto,
			report.reg_no,
			report.model,
			report.vehicle_class_desc,
			report.mobile,
			report.reg_upto,
			report.last_check_date,
			report.vehicle_class_desc,
		];
		const searchTerm = searchQuery.toLowerCase();
		return searchFields.some((field) => {
			const fieldString = String(field).toLowerCase();
			return fieldString.includes(searchTerm);
		});
	});

	return (
		<LazyMotion features={domAnimation}>
			<section className="h-full pb-5">
				{filteredData.map((report: Partial<CardData>, index: number) => (
					<AlertCard
						key={index}
						office_name={report.office_name || ""}
						owner_name={report.owner_name || ""}
						puc_status={report.puc_status || false}
						vehicle_type={report.vehicle_type || ""}
						reg_no={report.reg_no || ""}
						model={report.model || ""}
						vehicle_class_desc={report.vehicle_class_desc || ""}
						mobile={report.mobile || 0}
						puc_upto={report.puc_upto || { year: 0, month: 0, day: 0 }}
						reg_upto={report.reg_upto || { year: 0, month: 0, day: 0 }}
						last_check_date={report.puc_upto || { year: 0, month: 0, day: 0 }}
					/>
				))}
			</section>
		</LazyMotion>
	);
}
