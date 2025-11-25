import SectionHeader from "@components/headers/sectionHeader";
import QuestionList from "@components/search/question/list";
import LinkButton from "@components/buttons/Link";
import Divider from "@mui/material/Divider";
import Stack from "@mui/material/Stack";
import Guard from "@components/Guard";

import { ROLE_ADMIN } from "../../../../../models/server";

export default (): JSX.Element => {
    return (
        <Stack direction="column" gap={3} paddingBottom={4}>
            <SectionHeader title="Banco de questões">
                <Guard roles={ROLE_ADMIN}>
                    <LinkButton title="Adicionar questões" href="/platform/question/list/create" variant="solid" />
                </Guard>
            </SectionHeader>
            <Divider />
            <QuestionList />
        </Stack>
    )
}