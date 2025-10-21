import FormHelperText from "@mui/material/FormHelperText"
import OutlinedInput from "@mui/material/OutlinedInput"
import FormControl from "@mui/material/FormControl"
import Typography from "@mui/material/Typography"
import InputLabel from "@mui/material/InputLabel"
import MenuItem from "@mui/material/MenuItem"
import Select from "@mui/material/Select"
import Stack from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"

import { Controller } from "react-hook-form"
import { styled } from "@mui/material/styles"
import { useForm } from "react-hook-form"
import { useContext, useEffect } from "react"
import { FormContext } from "./Form"

export interface FirstSectionProps {
    questionListName: string
    questionListDifficultyLevel: number
}

const InputTitle = styled(Typography)(() => ({
    textTransform: "uppercase",
    fontWeight: "bold",
    color: "grey",
}))

export default function ({ }: { title: string }): JSX.Element {
    const { register, watch, formState: { isValid }, control } = useForm<FirstSectionProps>()
    const { setFormState } = useContext(FormContext)

    useEffect(() => {
        if (isValid == false) return setFormState("1", undefined);
        setFormState("1", { ...watch(), complete: true })
    }, [isValid])

    return (
        <Grid container spacing={2}>
            <Grid size={6}>
                <Stack gap={1}>
                    <InputTitle>Nome da lista de questões</InputTitle>
                    <FormControl>
                        <InputLabel
                            htmlFor="create-question-list-name-label"
                        >
                            Nome da lista de questões
                        </InputLabel>
                        <OutlinedInput
                            {...register("questionListName", {
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
            <Grid size={6}>
                <Stack gap={1}>
                    <InputTitle>Dificuldade da lista de questões</InputTitle>
                    <Controller
                        name="questionListDifficultyLevel"
                        control={control}
                        rules={{ required: true }}
                        render={({ field }) =>
                            <FormControl>
                                <InputLabel htmlFor="create-question-list-name-label">
                                    Nome da lista de questões
                                </InputLabel>
                                <Select
                                    {...field}
                                    id="create-question-list-name-label"
                                    label="Nome da lista de questões"
                                >
                                    <MenuItem value={"A"}>A</MenuItem>
                                    <MenuItem value={"B"}>B</MenuItem>
                                    <MenuItem value={"C"}>C</MenuItem>
                                </Select>
                            </FormControl>
                        }
                    />
                </Stack>
            </Grid>
        </Grid>
    )
}