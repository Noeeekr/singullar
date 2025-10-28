import SectionHeader from "@components/headers/sectionHeader";
import QuestionLists from "./components/QuestionList";
import OutlinedInput from "@mui/material/OutlinedInput";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import Typography from "@mui/material/Typography";
import LinkButton from "@components/buttons/Link";
import Divider from "@mui/material/Divider";
import Stack from "@mui/material/Stack";
import Grid from "@mui/material/Grid2";

import { useForm } from "react-hook-form";
import { useTheme } from "@mui/material";

import type { FilterOptions } from "./components/QuestionList";
import SolidButton from "@components/buttons/Default/Solid";
import { useState } from "react";
import { useAppSelector } from "@slices/store";
import { ROLE_ADMIN } from "../../../../models/server";

const QuestionsPage = (): JSX.Element => {
    const user = useAppSelector((store) => store.user.user)
    const theme = useTheme();

    const [isHidden, setIsHidden] = useState(false);

    const { watch } = useForm<FilterOptions>({ defaultValues: {} })
    return (
        <Stack direction="column" gap={3} paddingBottom={4}>
            <SectionHeader title="Banco de questões">
                {
                    user?.role == ROLE_ADMIN 
                        ? <LinkButton title="Adicionar questões" href="/platform/question/list/create" variant="solid" />
                        : <></>
                }
            </SectionHeader>
            <Divider/>
            <Stack direction="row" justifyContent="space-between">
                <Stack direction="row" alignItems="center" gap={2}>
                    <Typography
                        variant="h4"
                        fontWeight="bold"
                        component="h4"
                    >
                        Filtrar
                    </Typography>
                    <Typography
                        variant="body1"
                        color={theme.palette.primary.purpleDark}
                        component="p"
                        sx={{
                            cursor: "pointer",
                            textDecoration: "underline"
                        }}
                        onClick={() => setIsHidden(prev => !prev)}
                    >
                        {isHidden ? "Mostrar" : "Esconder"} filtros
                    </Typography>
                </Stack>
                <SolidButton title="Filtrar" />
            </Stack>
            <form style={{ 
                height: "auto",
                maxHeight: isHidden ? "0px" : "300px",
                transition: "max-height 250ms ease-in-out",
                overflow: "hidden",
            }}>
                <Grid container spacing={2}>
                    <Grid size={{ xs: 4, mobile: 12 }}>
                        <Stack direction="column" gap={1}>
                            <Typography variant="subtitle1" color="gray">FILTRAR POR NOME</Typography>
                            <FormControl>
                                <InputLabel htmlFor="question-list-filter-name-input">Nome da lista de questões</InputLabel>
                                <OutlinedInput id="question-list-filter-name-input" label="Nome da lista de questões" />
                            </FormControl>
                        </Stack>
                    </Grid>
                    { /* Deveria fazer uma query para matérias */}
                    <Grid size={{ xs: 4, mobile: 12 }}>
                        <Stack direction="column" gap={1}>
                            <Typography variant="subtitle1" color="gray">FILTRAR POR MATÉRIA</Typography>
                            <FormControl>
                                <InputLabel htmlFor="question-list-filter-subject-input">Nome da matéria</InputLabel>
                                <OutlinedInput id="question-list-filter-subject-input" label="Nome da matéria" />
                            </FormControl>
                        </Stack>
                    </Grid>
                    <Grid size={{ xs: 4, mobile: 12 }}>
                        <Stack direction="column" gap={1}>
                            <Typography variant="subtitle1" color="gray">FILTRAR POR DIFICULDADE</Typography>
                            <FormControl>
                                <InputLabel htmlFor="question-list-filter-subject-input">Especifique a dificuldade da lista</InputLabel>
                                <OutlinedInput id="question-list-filter-subject-input" label="Especifique a dificuldade da lista" />
                            </FormControl>
                        </Stack>
                    </Grid>
                </Grid>
            </form>
            <Divider/>
            <Typography variant="h4" fontWeight="bold" component="h4">Listas de questões</Typography>
            <QuestionLists filters={watch()} />
        </Stack>
    )
}
export default QuestionsPage;