import QuestionDisplayer from "@components/displays/question"
import ErrorBubble from "@components/bubbles/ErrorBubble"
import Typography from "@mui/material/Typography"
import LinkButton from "@components/buttons/Link"
import Stack from "@mui/material/Stack"
import Grid from "@mui/material/Grid2"

import useContextAwareFetch, { defaultRequestInit } from "@hooks/useContextAwareFetch"
import { useContext, useEffect, useState } from "react"
import { QuestionFilterContext, QuestionFilters } from "./Filters"
import { SERVER_ADDR } from "../../../../../../configs"

import type { Question } from "@models/server"
import { FormContext } from "../../create/components/CreateForm/Form"

export interface QuestionRequest {
    filters?: QuestionFilters[]
}

export default function (): JSX.Element {
    const { setFormState } = useContext(FormContext)
    const filters = useContext(QuestionFilterContext)
    const [selectedQuestionsIds, setSelectedQuestionsIds] = useState<number[]>([])

    const { response, isLoading, error, send } = useContextAwareFetch<Question[], QuestionRequest>(
        `${SERVER_ADDR}/api/question`,
        { ...defaultRequestInit, method: "POST" }
    )

    useEffect(() => {
        send(filters)
    }, [filters])

    useEffect(() => {
        if (selectedQuestionsIds.length) {
            setFormState("2", { question_list_question_ids: selectedQuestionsIds, complete: true })
            return
        }
        setFormState("2", { question_list_question_ids: [], complete: false })
    }, [selectedQuestionsIds])

    const handleSelection = (targetId: number) => {
        setSelectedQuestionsIds(ids => {
            let filteredIds = ids.filter(id => targetId != id)
            if (filteredIds.length == ids.length) filteredIds.push(targetId)
            return filteredIds
        })
    }

    return (
        <div>
            <Typography fontWeight="bold" component="p" color="grey" textTransform="capitalize">QUESTÕES</Typography>
            <Stack direction="column" alignItems="center" gap={1}>
                {
                    response?.length
                        ? <Grid container width="100%" marginY={2} spacing={2}>
                            {
                                response.map((question, i) => {
                                    return (
                                        <Grid size={{ sm: 6, mobile: 12 }}>
                                            <QuestionDisplayer
                                                key={i}
                                                selectable
                                                onSelected={handleSelection}
                                                isSelected={selectedQuestionsIds.some(id => id == question.id)}
                                                question={question}
                                            />
                                        </Grid>
                                    )
                                })
                            }
                        </Grid>
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