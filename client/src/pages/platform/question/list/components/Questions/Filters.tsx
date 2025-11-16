// Components
import OutlinedInput from "@mui/material/OutlinedInput"
import FormControl from "@mui/material/FormControl"
import InputLabel from "@mui/material/InputLabel"
import InputTitle from "@components/titles/BoldTitle"
import MenuItem from "@mui/material/MenuItem"
import Select from "@mui/material/Select"
import Stack from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"
import { Controller } from "react-hook-form"


// Utilities
import { createContext, useEffect, useState } from "react"
import { useForm } from "react-hook-form"
import useContextAwareFetch from "@hooks/useContextAwareFetch"
import { SERVER_ADDR } from "../../../../../../configs"

// Models
import type { StackProps } from "@mui/material/Stack"
import type { QuestionListDifficulty } from "@models/server"

export interface QuestionFilters {
    question_name?: string
    question_subject?: string
    question_difficulty_level?: number
}

export interface QuestionFilterProps extends StackProps {
    onParentClick: (cb: () => void) => void
}

export const QuestionFilterContext = createContext<QuestionFilters>({})

export default function ({ children, onParentClick, ...props }: QuestionFilterProps): JSX.Element {
    const { register, getValues, control } = useForm<QuestionFilters>()
    const [formValue, setFormValue] = useState<QuestionFilters>({})

    useEffect(() => {
        if (onParentClick) onParentClick(() => () => {
            setFormValue(getValues())
        })
    }, [])

    const { response, send } = useContextAwareFetch<QuestionListDifficulty[]>(
        `${SERVER_ADDR}/api/question/list/difficulties`,
        {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
        }
    )

    useEffect(() => {
        send()
    }, [])

    return (
        <Stack gap={2} {...props}>
            <Stack gap={1}>
                <InputTitle>FILTRAR</InputTitle>
                <Grid container spacing={2}>
                    <Grid size={{ mobile: 12, xss: 6 }}>
                        <FormControl>
                            <InputLabel htmlFor="create-question-difficulty-outlined-input">
                                Dificuldade da questão
                            </InputLabel>
                            <Controller
                                name="question_difficulty_level"
                                rules={{ required: true, validate: (v) => v != undefined }}
                                control={control}
                                render={({ field }) => (
                                    <Select
                                        {...field}
                                        id="create-question-difficulty-outlined-input"
                                        label="Dificuldade da questão"
                                        value={field.value ? field.value : ""}
                                    >
                                        <MenuItem value={undefined}>Selecione a dificuldade</MenuItem>
                                        {response?.map((difficulty) => (
                                            <MenuItem key={difficulty.question_difficulty_level} value={difficulty.question_difficulty_level}>{difficulty.question_difficulty_name}</MenuItem>
                                        ))}
                                    </Select>
                                )}
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
                                {...register("question_subject")}
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
                                {...register("question_name")}
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