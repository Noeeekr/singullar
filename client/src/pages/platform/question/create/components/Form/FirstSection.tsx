// Components
import OutlinedInput from "@mui/material/OutlinedInput"
import FormControl from "@mui/material/FormControl"
import Typography from "@mui/material/Typography"
import InputLabel from "@mui/material/InputLabel"
import Stack from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"

// Features
import { createContext } from "react"
import { staticValues } from "@components/forms/MultiStepForm/Form"

// Models
import { type FormContextSectionValues } from "@components/forms/MultiStepForm/Form"
import type { FormContextProps } from "./Form"
import { useForm } from "react-hook-form"

export interface FirstSectionContextProps {
    title: string
    difficultyLevel: number
}

export const FirstSectionContext = createContext<FormContextSectionValues<FormContextProps, FirstSectionContextProps>>({
    ...staticValues,
    title: "",
})

export default function (): JSX.Element {
    const { register } = useForm<FirstSectionContextProps>()

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
                            {...register("title")}
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
                        <OutlinedInput
                            {...register("difficultyLevel")}
                            id="create-question-difficulty-outlined-input"
                            label="Dificuldade da questão"
                        />
                    </FormControl>
                </Stack>
            </Grid>
        </Grid>
    )
}