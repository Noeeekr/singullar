// Components
import Typography from "@mui/material/Typography"
import Grid from "@mui/material/Grid2";

// Utilities
import { useEffect } from "react"
import { SERVER_ADDR } from "../../../../configs";
import useContextAwareFetch, { defaultRequestInit } from "@hooks/useContextAwareFetch";

// Models
import type { StackProps } from "@mui/material"
import type { QuestionList } from "@models/server/question";
import type { JSX } from "react"
import ErrorBubble from "@components/bubbles/ErrorBubble";
import QuestionListDisplayer from "@components/displays/question/list";

export interface FilterOptions {
    name?: string,
    subject?: string,
    difficultyName?: string,
    difficultyLevel?: number,
}
export interface SearchQuestionListFilterProps {
    filters?: FilterOptions[]
}
export interface SearchQuestionListStylingProps extends StackProps { }
export type SearchQuestionListProps =
    SearchQuestionListFilterProps
    & SearchQuestionListStylingProps

export default ({ filters }: SearchQuestionListProps): JSX.Element => {
    const { response: lists, isLoading, error, send } = useContextAwareFetch<QuestionList[], FilterOptions[]>(
        `${SERVER_ADDR}/api/question/list`,
        { ...defaultRequestInit, method: "POST" }
    )

    useEffect(() => {
        send(filters)
    }, [filters])

    if (lists == null) {
        if (isLoading) return <Typography fontWeight="bold" variant="body1">Carregando listas...</Typography>
        return <ErrorBubble message={error} />
    }
    if (lists.length == 0) return <ErrorBubble message="Nenhuma lista de questão encontrada" />

    return (
        <Grid container spacing={2}>
            {
                lists.map((list, i) => (
                    <Grid size={12} key={i}>
                        <QuestionListDisplayer navigable list={list}/>
                    </Grid>
                ))
            }
        </Grid>
    )
}