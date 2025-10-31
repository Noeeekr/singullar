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
    method: "GET",
    headers: {
        "Content-Type": "application/json"
    },
    credentials: "include",
    cache: "default"
}

export default function (): JSX.Element {
    const filters = useContext(QuestionFilterContext)

    const { response, isLoading, error, send } = useContextAwareFetch<Question[]>(
        `${SERVER_ADDR}/api/question/list`,
        requestInit
    )
    // Needs endpoint

    useEffect(() => {
        send()
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
                                ? <ErrorBubble err={error} />
                                : <Stack direction="column" alignItems="center" gap={1}>
                                    <Typography>Nenhuma questão encontrada?</Typography>
                                    <LinkButton title="Criar questões" href="/platform/question/create" variant="solid" />
                                </Stack>

                }
            </Stack>
        </div>
    )
}