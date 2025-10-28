import ErrorBubble from "@components/bubbles/ErrorBubble"
import Typography from "@mui/material/Typography"
import LinkButton from "@components/buttons/Link"
import Stack from "@mui/material/Stack"

import useContextAwareFetch from "@hooks/useContextAwareFetch"
import { useContext, useEffect } from "react"
import { QuestionFilterContext } from "./Filters"
import { SERVER_ADDR } from "../../../../../../configs"

import type { Question } from "@models/server"

const requestInit: RequestInit = {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    credentials: "include",
    cache: "default"
}

export default function (): JSX.Element {
    const filters = useContext(QuestionFilterContext)

    const { response, isLoading, error, send } = useContextAwareFetch<Question[], null>(
        `${SERVER_ADDR}/api/questions`,
        requestInit
    )
    // Needs endpoint

    useEffect(() => {
        send(null)
    }, [filters])
    return (
        <div>
            <Typography fontWeight="bold" component="p" color="grey" textTransform="capitalize">QUESTÕES</Typography>
            <Stack direction="column" alignItems="center" gap={1}>
                {
                    response?.length
                        ? <></>
                        : isLoading
                            ? <Typography>Procurando questões...</Typography>
                            : error
                                ? <ErrorBubble err={"Falha ao encontrar questões"} />
                                : <Stack direction="column" alignItems="center" gap={1}>
                                    <Typography>Nenhuma questão encontrada?</Typography>
                                    <LinkButton title="Criar questões" href="/platform/question/create" variant="solid" />
                                </Stack>

                }
            </Stack>
        </div>
    )
}