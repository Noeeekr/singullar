import FormHelperText from "@mui/material/FormHelperText"
import OutlinedInput from "@mui/material/OutlinedInput"
import FormControl from "@mui/material/FormControl"
import InputLabel from "@mui/material/InputLabel"
import MenuItem from "@mui/material/MenuItem"
import Select from "@mui/material/Select"
import Stack from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"

import { Controller } from "react-hook-form"
import { useForm } from "react-hook-form"
import { useContext, useEffect } from "react"
import { FormContext } from "./Form"
import useContextAwareFetch, { defaultRequestInit } from "@hooks/useContextAwareFetch"
import { SERVER_ADDR } from "../../../../../../../configs"
import { QuestionListDifficulty, Subject } from "@models/server"
import InputTitle from "@components/titles/BoldTitle"

export interface FirstSectionFormState {
    question_list_title: string,
    question_list_subject_id: number,
    question_list_difficulty_level: number,
}

export default function ({ }: { title: string }): JSX.Element {
    const { register, getValues, formState: { isValid }, control } = useForm<FirstSectionFormState>()
    const { setFormState } = useContext(FormContext)

    useEffect(() => {
        if (isValid == false) return setFormState("1", undefined);
        setFormState("1", { ...getValues(), complete: true })
    }, [isValid])

    const { response: subjects, send: sendSubjectRequest } = useContextAwareFetch<Subject[]>(
        `${SERVER_ADDR}/api/subjects`,
        defaultRequestInit,
    )

    const { response: difficulties, send: sendDifficultyRequest } = useContextAwareFetch<QuestionListDifficulty[]>(
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
        sendDifficultyRequest()
        sendSubjectRequest()
    }, [])

    return (
        <Grid container spacing={2}>
            <Grid size={{ mobile: 12, xss: 6 }}>
                <Stack gap={1}>
                    <InputTitle>Nome da lista de questões</InputTitle>
                    <FormControl>
                        <InputLabel
                            htmlFor="create-question-list-name-label"
                        >
                            Nome da lista de questões
                        </InputLabel>
                        <OutlinedInput
                            {...register("question_list_title", {
                                required: true,
                                minLength: 8,
                                maxLength: 225,
                            })}
                            id="create-question-list-name-label"
                            label="Nome da lista de questões"
                        />
                        {
                            !isValid
                                ? <FormHelperText sx={(theme) => ({ color: theme.palette.error.dark })}>*Por favor, insira um mínimo de 8 letras</FormHelperText>
                                : <></>
                        }
                    </FormControl>
                </Stack>
            </Grid>
            <Grid size={{ mobile: 12, xss: 6 }}>
                <Stack gap={1}>
                    <InputTitle>Dificuldade da lista de questões</InputTitle>
                    <FormControl>
                        <InputLabel htmlFor="create-question-difficulty-outlined-input">
                            Dificuldade da questão
                        </InputLabel>
                        <Controller
                            name="question_list_difficulty_level"
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
            <Grid size={{ mobile: 12, xss: 6 }}>
                <Stack gap={1}>
                    <InputTitle>Matéria da lista de questões</InputTitle>
                    <FormControl>
                        <InputLabel htmlFor="create-question-subject-outlined-input">
                            Matéria da questão
                        </InputLabel>
                        <Controller
                            name="question_list_subject_id"
                            rules={{ required: true, validate: (v) => v != undefined }}
                            control={control}
                            render={({ field }) => (
                                <Select
                                    {...field}
                                    id="create-question-subject-outlined-input"
                                    label="Matéria da questão"
                                    value={field.value ? field.value : ""}
                                >
                                    <MenuItem value={undefined}>Selecione a matéria</MenuItem>
                                    {subjects?.map((subject) => (
                                        <MenuItem key={subject.id} value={subject.id}>{subject.subject_name}</MenuItem>
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