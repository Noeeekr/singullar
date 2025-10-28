import OutlinedInput from "@mui/material/OutlinedInput"
import FormControl from "@mui/material/FormControl"
import InputLabel from "@mui/material/InputLabel"
import Typography from "@mui/material/Typography"
import Stack from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"

import { createContext, useEffect, useState } from "react"
import { useForm } from "react-hook-form"

import type { StackProps } from "@mui/material/Stack"

export interface QuestionFilters {
    name?: string
    subject?: string
    difficultyLevel?: number
}
export interface QuestionFilterProps extends StackProps {
    onParentClick: (cb: () => void) => void
}
export const QuestionFilterContext = createContext<QuestionFilters>({})

export default function ({ children, onParentClick, ...props }: QuestionFilterProps): JSX.Element {
    const { register, getValues } = useForm<QuestionFilters>()
    const [formValue, setFormValue] = useState<QuestionFilters>({})

    useEffect(() => {
        if (onParentClick) onParentClick(() => () => {
            setFormValue(getValues())
        })
    }, [])

    return (
        <Stack gap={2} {...props}>
            <Stack gap={1}>
                <Typography
                    color="grey"
                    component="p"
                    fontWeight="bold"
                    textTransform="capitalize"
                >FILTRAR</Typography>
                <Grid container spacing={2}>
                    <Grid size={{ mobile: 12, xss: 6 }}>
                        <FormControl>
                            <InputLabel htmlFor="questions-search-difficulty-outlined-input">
                                Dificuldade
                            </InputLabel>
                            <OutlinedInput
                                type="number"
                                label="Dificuldade"
                                {...register("difficultyLevel")}
                                id="questions-search-difficulty-outlined-input"
                            />
                        </FormControl>
                    </Grid>
                    <Grid size={{ mobile: 12, xss: 6 }}>
                        <FormControl>
                            <InputLabel htmlFor="questions-search-subject-outlined-input">
                                Matéria
                            </InputLabel>
                            <OutlinedInput
                                label="Matéria"
                                {...register("name")}
                                id="questions-search-subject-outlined-input"
                            />
                        </FormControl>
                    </Grid>
                    <Grid size={{ mobile: 12, xss: 12 }}>
                        <FormControl>
                            <InputLabel htmlFor="questions-search-name-outlined-input">
                                Nome
                            </InputLabel>
                            <OutlinedInput
                                label="Nome"
                                {...register("name")}
                                id="questions-search-name-outlined-input"
                            />
                        </FormControl>
                    </Grid>
                </Grid>
            </Stack>
            <QuestionFilterContext.Provider value={formValue}>
                {children}
            </QuestionFilterContext.Provider>
        </Stack>
    )
}