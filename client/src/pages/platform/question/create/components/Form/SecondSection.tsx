import OutlinedInput from "@mui/material/OutlinedInput"
import SolidButton from "@components/buttons/Default/Solid"
import FormControl from "@mui/material/FormControl"
import InputLabel from "@mui/material/InputLabel"
import Typography from "@mui/material/Typography"
import TextArea from "@components/TextArea"
import Stack, { StackProps } from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"

import { JSX, createContext, useEffect, useState } from "react"
import { useForm } from "react-hook-form"
import CircularButton from "@components/buttons/Circular"

export interface SecondSectionContextProps {
    questionDescription?: string
    questionQuestion?: string
    alternatives?: string[]
    correctAlternative?: string
}
export interface SecondSectionProps {
    update: (values: SecondSectionContextProps | null) => void
}

export const SecondSectionContext = createContext<SecondSectionContextProps>({})

export interface AlternativeDisplayerProps extends StackProps {
    onDelete: (title: string) => void
    title: string
    decorativeIconIndex: number
}

const AlternativeCreatorDisplayer = ({
    title,
    onDelete,
    decorativeIconIndex,
    ...props
}: AlternativeDisplayerProps): JSX.Element => {
    return (
        <Stack {...props} direction="row" alignItems="center" justifyContent="space-between">
            <Stack direction="row" alignItems="center" justifyContent="space-between" gap={2}>
                <CircularButton 
                    sx={(theme) => ({
                        width: 40,
                        height: 40,
                        paddingTop: 0.2,
                        borderColor: "rgb(170,170,170)",
                        backgroundColor: "rgb(249,249,249)",
                        color: "rgb(90,90,90)",
                        "&:hover": {
                            borderColor: theme.palette.primary.purpleLight,
                            backgroundColor: theme.palette.primary.purpleLight,
                            color: "white",
                        },
                    })}
                >
                    { String.fromCharCode(('a'.charCodeAt(0) + decorativeIconIndex)).toUpperCase() }
                </CircularButton>
                <Typography
                    variant="body2"
                    component="p"
                >
                    {title}
                </Typography>
            </Stack>
            <SolidButton title="Remover alternativa" onClick={() => onDelete(title)} />
        </Stack>
    )
}

export interface AlternativeCreatorProps extends StackProps {
    onUpdate: (title: string) => void
}

const AlternativeCreator = ({
    onUpdate,
    ...props
}: AlternativeCreatorProps) => {
    const [title, setTitle] = useState("")

    return (
        <Stack {...props} direction="row" justifyContent="space-between" alignItems="center" gap={2}>
            <FormControl>
                <InputLabel htmlFor={`outlined-input-alternative-creator-${props.key}`}>
                    Titulo da alternativa
                </InputLabel>
                <OutlinedInput
                    id={`outlined-input-alternative-creator-${props.key}`}
                    label="Titulo da alternativa"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                />
            </FormControl>
            <SolidButton
                title="Finalizar"
                disabled={!title}
                onClick={() => onUpdate(title)}
            />
        </Stack>
    )
}
export default function ({ update }: SecondSectionProps): JSX.Element {
    const { register, watch, getValues, setValue, formState: { isValid } } = useForm<SecondSectionContextProps>()
    const [alternativeAmount, setAlternativeAmount] = useState(0)
    const questionDescription = watch("questionDescription")
    const questionQuestion = watch("questionQuestion")
    const alternatives = watch("alternatives")

    useEffect(() => {
        if (isValid) {
            update(getValues())
        } else {
            update(null)
        }
    }, [questionDescription, questionQuestion, isValid])

    const handleAlternativeAddition = (title: string) => {
        let s = getValues("alternatives")
        if (s == undefined) s = [];
        s.push(title)
        setValue("alternatives", s)
        setAlternativeAmount(prev => --prev)
    }
    const handleAlternativeDelete = (title: string) => {
        let s = getValues("alternatives")
        if (s == undefined) return
        s = s.filter(alternativeTitle => alternativeTitle != title)
        if (s.length == 0) s = undefined
        setValue("alternatives", s)
    }
    return (
        <Stack direction="column" gap={2}>
            <Grid container spacing={2}>
                <Grid size={12}>
                    <Stack gap={1}>
                        <Typography fontWeight="bold" color="grey" component="p">DESCRIÇÃO</Typography>
                        <FormControl>
                            <TextArea
                                label="Insira o conteúdo da questão"
                                {...register("questionDescription", {
                                    minLength: 8,
                                    required: true,
                                })}
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
            <Stack direction="row" gap={2} justifyContent="space-between" alignItems="center">
                <Typography fontWeight="bold" color="grey" component="p">ALTERNATIVAS</Typography>
                <SolidButton
                    maxWidth={200}
                    title="Adicionar alternativa"
                    onClick={() => setAlternativeAmount(prev => ++prev)}
                />
            </Stack>
            {
                (() => {
                    const alternatives: JSX.Element[] = new Array(alternativeAmount)
                    for (let i = 0; i < alternativeAmount; i++) {
                        alternatives[i] = <AlternativeCreator key={i} onUpdate={handleAlternativeAddition} />
                    }
                    return alternatives
                })()
            }
            {
                alternatives?.map((alternative, i) => {
                    return <AlternativeCreatorDisplayer key={i} decorativeIconIndex={i} title={alternative} onDelete={handleAlternativeDelete} />
                })
            }
        </Stack>
    )
}