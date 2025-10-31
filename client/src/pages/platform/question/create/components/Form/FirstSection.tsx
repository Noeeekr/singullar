// Components
import OutlinedInput from "@mui/material/OutlinedInput"
import FormControl from "@mui/material/FormControl"
import Typography from "@mui/material/Typography"
import InputLabel from "@mui/material/InputLabel"
import Stack from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"

// Features
import { createContext, useContext, useEffect } from "react"

// Models
import { useForm } from "react-hook-form"

export interface FirstSectionContextProps {
    questionTitle?: string
    difficultyLevel?: number
}

export const FirstSectionContext = createContext<FirstSectionContextProps>({
    questionTitle: "",
    difficultyLevel: 0,
})

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
    update: (values: FirstSectionContextProps) => void
}): JSX.Element {
    const { register, getValues, watch, formState: { isValid, defaultValues } } = useForm<FirstSectionContextProps>({
        defaultValues: {
            difficultyLevel: undefined,
            questionTitle: undefined,
        }
    })

    const questionTitle = watch("questionTitle")
    const difficultyLevel = watch("difficultyLevel")
    
    useEffect(() => {
        if (isValid) {
            update(getValues())
        } else {
            update(defaultValues as FirstSectionContextProps)
        }
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
                            {...register("questionTitle", {
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
                        <OutlinedInput
                            {...register("difficultyLevel", {
                                required: true,
                                validate: (value) => value !== undefined && (Boolean(value) || value >= 0),
                            })}
                            type="number"
                            id="create-question-difficulty-outlined-input"
                            label="Dificuldade da questão"
                        />
                    </FormControl>
                </Stack>
            </Grid>
        </Grid>
    )
}