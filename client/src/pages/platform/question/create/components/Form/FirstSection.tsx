// Components
import OutlinedInput from "@mui/material/OutlinedInput"
import FormControl from "@mui/material/FormControl"
import InputLabel from "@mui/material/InputLabel"
import Stack from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"

// Features
import { createContext, useEffect } from "react"

// Models
import { Controller, useForm } from "react-hook-form"
import useContextAwareFetch, { defaultRequestInit } from "@hooks/useContextAwareFetch"
import { SERVER_ADDR } from "@components/../configs"
import { QuestionListDifficulty, Subject } from "@models/server"
import { MenuItem, Select } from "@mui/material"
import BoldTitle from "@components/titles/BoldTitle"

export interface FirstSectionContextProps {
    question_title?: string
    question_subject_id?: number
    question_difficulty_level?: number
}

export const FirstSectionContext = createContext<FirstSectionContextProps>({})

export default function ({
    update
}: {
    update: (values: FirstSectionContextProps | null) => void
}): JSX.Element {
    const { register, getValues, watch, control, formState: { isValid } } = useForm<FirstSectionContextProps>({
        defaultValues: {
            question_difficulty_level: undefined,
            question_title: undefined,
        }
    })

    const { response: difficulties, send: sendDifficultiesRequest } = useContextAwareFetch<QuestionListDifficulty[]>(
        `${SERVER_ADDR}/api/question/list/difficulties`,
        defaultRequestInit,
    )

    const questionTitle = watch("question_title")
    const difficultyLevel = watch("question_difficulty_level")

    const { response: subjects, send: sendSubjectsRequest } = useContextAwareFetch<Subject[]>(
        `${SERVER_ADDR}/api/subjects`,
        {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
        }
    )
    
    console.log(subjects)
    useEffect(() => {
        sendDifficultiesRequest()
        sendSubjectsRequest()
    }, [])

    useEffect(() => {
        if (isValid) {
            update(getValues())
            return
        }
        update(null)
    }, [difficultyLevel, questionTitle, isValid])

    return (
        <Grid container spacing={2}>
            <Grid size={12}>
                <Stack gap={1}>
                    <BoldTitle>Titulo</BoldTitle>
                    <FormControl>
                        <InputLabel htmlFor="create-question-title-outlined-input">
                            Insira o titulo da questão
                        </InputLabel>
                        <OutlinedInput
                            {...register("question_title", {
                                minLength: 8,
                                required: true,
                            })}
                            id="create-question-title-outlined-input"
                            label="Insira o título da questão"
                        />
                    </FormControl>
                </Stack>
            </Grid>
            <Grid size={6}>
                <Stack gap={1}>
                    <BoldTitle>Matéria</BoldTitle>
                    <FormControl>
                        <InputLabel htmlFor="questions-search-subject-outlined-input">
                            Matéria
                        </InputLabel>
                        <Controller
                            name="question_subject_id"
                            control={control}
                            rules={{ required: true }}
                            render={({ field }) => (
                                <Select
                                    required={true}
                                    {...field}
                                    label="Matéria"
                                    id="questions-search-subject-outlined-input"
                                >
                                    <MenuItem value={undefined}>Escolha uma matéria</MenuItem>
                                    {
                                        subjects?.map((subject) => <MenuItem value={subject.id}>{subject.subject_name} </MenuItem>)
                                    }
                                </Select>
                            )}
                        />
                    </FormControl>
                </Stack>
            </Grid>
            <Grid size={6}>
                <Stack gap={1}>
                    <BoldTitle>Dificuldade</BoldTitle>
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
                                    {difficulties?.map((difficulty) => (
                                        <MenuItem key={difficulty.question_difficulty_level} value={difficulty.question_difficulty_level}>{difficulty.question_difficulty_name}</MenuItem>
                                    ))}
                                </Select>
                            )}
                        />
                    </FormControl>
                </Stack>
            </Grid>
        </Grid>
    )
}