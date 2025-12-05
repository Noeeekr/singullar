import SearchQuestionLists from "./default";
import OutlinedInput from "@mui/material/OutlinedInput";
import SolidButton from "@components/buttons/Default/Solid";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Stack from "@mui/material/Stack";
import Grid from "@mui/material/Grid2";

import { useForm } from "react-hook-form";
import { useState } from "react";

import type { JSX } from "react"
import type { FilterOptions } from "@hooks/useFetchQuestionLists";

export default (): JSX.Element => {
    const [isHidden, setIsHidden] = useState(false);

    const { watch } = useForm<FilterOptions>({ defaultValues: {} })

    return (
        <Stack gap={3}>
            <Stack direction="row" justifyContent="space-between">
                <Stack direction="row" alignItems="center" gap={1}>
                    <Typography
                        variant="h4"
                        fontWeight="bold"
                        component="h4"
                    >
                        Filtrar
                    </Typography>
                    <Typography
                        variant="body1"
                        component="p"
                        sx={(theme) => ({
                            cursor: "pointer",
                            color: theme.palette.primary.purpleDark,
                            textDecoration: "underline"
                        })}
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
            <Divider />
            <Typography variant="h4" fontWeight="bold" component="h4">Listas de questões</Typography>
            <SearchQuestionLists filters={[watch()]} />
        </Stack>
    )
}