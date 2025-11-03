import FormControl from "@mui/material/FormControl"
import Grid from "@mui/material/Grid2"
import InputLabel from "@mui/material/InputLabel"
import OutlinedInput from "@mui/material/OutlinedInput"
import Stack from "@mui/material/Stack"
import Typography from "@mui/material/Typography"

import { Controller, useForm } from "react-hook-form"
import { createContext } from "react"
import TextArea from "@components/TextArea"

export interface SecondSectionContextProps {
    questionDescription?: string
    questionQuestion?: string
}

export const SecondSectionContext = createContext<SecondSectionContextProps>({})

export default function (_: { update: (values: SecondSectionContextProps) => void }): JSX.Element {
    const { register, watch, control, formState: { defaultValues } } = useForm<SecondSectionContextProps>({
        defaultValues: {
            questionDescription: undefined,
        }
    })

    return (
        <Grid container spacing={2}>
            <Grid size={12}>
                <Stack gap={1}>
                    <Typography fontWeight="bold" color="grey" component="p">DESCRIÇÃO</Typography>
                    <FormControl>
                        <Controller
                            name="questionDescription"
                            control={control}
                            rules={{ required: true, minLength: 8 }}
                            render={({ field }) => (
                                <TextArea
                                    label="Insira o conteúdo da questão"
                                    {...field}
                                />
                            )}
                        />
                    </FormControl>
                </Stack>
            </Grid>
            <Grid size={12}>
                <Stack gap={1}>
                    <Typography fontWeight="bold" color="grey" component="p">PERGUNTA</Typography>
                    <FormControl>
                        <InputLabel htmlFor="create-question-question-outlined-input">
                            Insira a pergunta da questão
                        </InputLabel>
                        <OutlinedInput
                            {...register("questionQuestion", {
                                minLength: 8,
                                required: true,
                            })}
                            id="create-question-question-outlined-input"
                            label="Insira a pergunta da questão"
                        />
                    </FormControl>
                </Stack>
            </Grid>
        </Grid>
    )
}