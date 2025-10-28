import SectionHeader from "@components/headers/sectionHeader"
import Stack from "@mui/material/Stack"

export default function (): JSX.Element {
    return (
        <Stack gap={4}>
            <SectionHeader title="Criar questão" subtitle="Informe as informações necessárias para criar uma nova questão" />
        </Stack>
    )
}