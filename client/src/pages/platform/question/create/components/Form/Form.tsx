// Components
import MultiStepForm from "@components/forms/MultiStepForm"
import FirstSection, { FirstSectionContext, FirstSectionContextProps } from "./FirstSection"
import SecondSection, { SecondSectionContext, SecondSectionContextProps } from "./SecondSection"
import Typography from "@mui/material/Typography"
import Stack from "@mui/material/Stack"
import Box from "@mui/material/Box"

// Utilities
import useContextAwareFetch, { defaultRequestInit } from "@hooks/useContextAwareFetch"
import { useEffect, useState } from "react"
import { SERVER_ADDR } from "../../../../../../configs"

// Models
import type { SectionHeader } from "@components/forms/MultiStepForm/Header"
import ErrorBubble from "@components/bubbles/ErrorBubble"

const sections = [
    {
        content: FirstSection,

        title: "Informações Básicas",

        context: FirstSectionContext,
        defaultValues: {
            question_title: undefined as string | undefined,
            question_difficulty_level: undefined as number | undefined,
        }
    },
    {
        content: SecondSection,

        title: "Conteúdo da questão",

        context: SecondSectionContext,
        defaultValues: {
            question_short_description: undefined as undefined | string,
            question_description: undefined as undefined | string,
            alternatives: [] as string[] | undefined
        }
    },
]

type FormData = FirstSectionContextProps & SecondSectionContextProps
interface FormOperationHistoryCell {
    title: string
    ok: boolean
}

export default function (): JSX.Element {
    const headers: SectionHeader[] = [
        { label: "Informações básicas" },
        { label: "Adicionar questões" },
    ]

    const { send, error, registerOnResponseCallback } = useContextAwareFetch<null, FormData>(
        `${SERVER_ADDR}/api/question/create`,
        {
            ...defaultRequestInit,
            method: "POST",
        },
    )

    const [operationHistory, setOperationHistory] = useState<FormOperationHistoryCell[]>([])

    const onSubmit = (formData: FirstSectionContextProps & SecondSectionContextProps) => {
        send(formData)
    }

    useEffect(() => {
        registerOnResponseCallback(({ body }) => {
            setOperationHistory((history) => {
                history.push(error
                    ? { title: error, ok: false }
                    : { title: `Criou a seguinte questão: ${body?.question_title}`, ok: true }
                )
                return history
            })
        })
    }, [])

    return (
        <div>
            <MultiStepForm sections={sections} onSubmit={onSubmit} headers={headers} />
            <ErrorBubble err={error} />
            {
                operationHistory.length == 0
                    ? <></>
                    : <Stack gap={2} component="section">
                        <Typography variant="body2" fontWeight="bold" textTransform="capitalize">
                            Ultimas ações
                        </Typography>
                        <Box sx={{
                            width: "min(560px, 100%)",
                            backgroundColor: "rgb(245,245,245)",
                            border: "solid 1px rgb(32,30,35)",
                            borderTopLeftRadius: 9,
                            borderTopRightRadius: 9,
                            borderBottomLeftRadius: 5,
                            borderBottomRightRadius: 5,

                            overflow: "hidden",
                        }}>
                            {operationHistory.map((action) => {
                                return (
                                    <Typography variant="body2" sx={() => ({
                                        backgroundColor: action.ok ? "rgba(130, 210, 120, 0.8)" : "rgba(250,170,160,0.7)",
                                        padding: 1,

                                        userSelect: "none",

                                        "&:hover": {
                                            backgroundColor: action.ok ? "rgba(100, 180, 110, 0.9)" : "rgba(250,120,120,0.7)"
                                        }
                                    })}>
                                        {action.title}
                                    </Typography>
                                )
                            })}
                        </Box>
                    </Stack>
            }
        </div>
    )
}