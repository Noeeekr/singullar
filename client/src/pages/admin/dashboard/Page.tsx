// Components
import SectionHeader from "@components/headers/sectionHeader/SectionHeader";
import Typography from "@mui/material/Typography"
import Stack from "@mui/material/Stack"

// Models
import SubjectInformation from "./components/SubjectInformation";
import InstitutionInformation from "./components/InstitutionInformation";

const Dashboard = (): JSX.Element => {
    return (
        <Stack gap={2}>
            <SectionHeader title="Informações Escolares" subtitle="Ultima atualização: Agora" />
            <Typography variant="h5" component="h4" fontWeight="bold">
                Informações Gerais
            </Typography>
            <InstitutionInformation />
            <Typography variant="h5" component="h4" fontWeight="bold">
                Curriculo Escolar
            </Typography>
            <SubjectInformation />
        </Stack>
    )
}

export default Dashboard;
