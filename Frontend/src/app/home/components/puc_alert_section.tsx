import { useReportsData } from "@/hooks/reports/useReportData";
import { CardData } from "@/lib/data";
import { LazyMotion, domAnimation, m, stagger } from "framer-motion";
import AlertCard from "./alert-card";

interface ReportsProps {
	searchQuery: string;
}

export default function Reports({ searchQuery }: ReportsProps) {
	const { reportsData } = useReportsData();

	const filteredData = reportsData.filter((report: CardData) => {
		const searchFields = [
			report.office_name,
			report.pucStatus,
			report.vehicleType,
			report.validUpto,
			report.registrationNo,
			report.vehicleModel,
			report.vehicleDescription,
			report.contact,
			report.pucValidUpto,
		];
		return searchFields.some((field) =>
			field.toLowerCase().includes(searchQuery.toLowerCase())
		);
	});

	return (
		<LazyMotion features={domAnimation}>
			<section className="h-full pb-5">
				{filteredData.map((report: CardData, index: number) => (
					<AlertCard
						key={index}
						office_name=""
						pucStatus={report.pucStatus}
						vehicleType={report.vehicleType}
						validUpto={report.validUpto}
						registrationNo={report.registrationNo}
						vehicleModel={report.vehicleModel}
						vehicleDescription={report.vehicleDescription}
						contact={report.contact}
						pucValidUpto={report.pucValidUpto}
					/>
				))}
			</section>
		</LazyMotion>
	);
}
