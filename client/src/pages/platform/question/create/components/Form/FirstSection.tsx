// Components
import OutlinedInput from "@mui/material/OutlinedInput"
import FormControl from "@mui/material/FormControl"
import Typography from "@mui/material/Typography"
import InputLabel from "@mui/material/InputLabel"
import Stack from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"

// Features
import { createContext, useEffect } from "react"

// Models
import { Controller, useForm } from "react-hook-form"
import useContextAwareFetch from "@hooks/useContextAwareFetch"
import { SERVER_ADDR } from "@components/../configs"
import { QuestionListDifficulty } from "@models/server"
import { MenuItem, Select } from "@mui/material"

export interface FirstSectionContextProps {
    question_title?: string
    question_difficulty_level?: number
}

export const FirstSectionContext = createContext<FirstSectionContextProps>({})

/**
    Updater({ sectionContext, onUpdate }) => {
        const sectionState = = useContex(context)
        update(formState => { ...formState, ...sectionState })
    }
    Context Provider (Section)
        <Updater context={context} onUpdate={update} /> 
        (Generic::Injected) { section.Header }     
        (Generic::Injected) { section.Content }
        (Generic::Injected) { section.Footer } (onClick:Advance)

    *Each section is independent
    *Their result is saved in a shared context
    *OnSubmit sends the shared context value
*/
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

    const questionTitle = watch("question_title")
    const difficultyLevel = watch("question_difficulty_level")

    useEffect(() => {
        send()
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
            <Grid size={6}>
                <Stack gap={1}>
                    <Typography fontWeight="bold" color="grey" component="p">TITULO</Typography>
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
                    <Typography fontWeight="bold" color="grey" component="p">DIFICULDADE</Typography>
                    <FormControl>
                        <InputLabel htmlFor="create-question-difficulty-outlined-input">
                            Dificuldade da questão
                        </InputLabel>
                        <Controller
                            name="question_difficulty_level"
                            rules={{ required: true, validate: (v) => v != undefined}}
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
                </Stack>
            </Grid>
        </Grid>
    )
}