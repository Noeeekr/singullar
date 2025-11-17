import ErrorBubble from "@components/bubbles/ErrorBubble"
import Typography from "@mui/material/Typography"
import LinkButton from "@components/buttons/Link"
import Stack from "@mui/material/Stack"

import useContextAwareFetch, { defaultRequestInit } from "@hooks/useContextAwareFetch"
import { useContext, useEffect } from "react"
import { QuestionFilterContext, QuestionFilters } from "./Filters"
import { SERVER_ADDR } from "../../../../../../configs"

import type { Question } from "@models/server"
import QuestionDisplay from "../Question/Display"

export interface QuestionRequest {
    filters?: QuestionFilters[]
}

export default function (): JSX.Element {
    const filters = useContext(QuestionFilterContext)

    const { response, isLoading, error, send } = useContextAwareFetch<Question[], QuestionRequest>(
        `${SERVER_ADDR}/api/question`,
        { ...defaultRequestInit, method: "POST" }
    )

    useEffect(() => {
        send(filters)
    }, [filters])
    
    return (
        <div>
            <Typography fontWeight="bold" component="p" color="grey" textTransform="capitalize">QUESTÕES</Typography>
            <Stack direction="column" alignItems="center" gap={1}>
                {
                    response?.length
                        ? response.map((question, i) => <QuestionDisplay key={i} data={question} />)
                        : isLoading
                            ? <Typography>Procurando questões...</Typography>
                            : error
                                ? <ErrorBubble message={error} />
                                : <Stack direction="column" alignItems="center" gap={1}>
                                    <Typography>Nenhuma questão encontrada?</Typography>
                                    <LinkButton title="Criar questões" href="/platform/question/create" variant="solid" />
                                </Stack>

                }
            </Stack>
        </div>
    )
}