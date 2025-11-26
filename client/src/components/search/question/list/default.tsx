// Components
import QuestionListDisplayer from "@components/displays/question/list";
import ErrorBubble from "@components/bubbles/ErrorBubble";
import Typography from "@mui/material/Typography"
import Grid from "@mui/material/Grid2";

// Utilities
import useFetchQuestionLists from "@hooks/useFetchQuestionLists";
import { useEffect } from "react"

// Models
import type { FilterOptions } from "@hooks/useFetchQuestionLists";
import type { StackProps } from "@mui/material"
import type { JSX } from "react"


export interface SearchQuestionListFilterProps {
    filters?: FilterOptions[]
}

export interface SearchQuestionListStylingProps extends StackProps { }
export type SearchQuestionListProps =
    SearchQuestionListFilterProps
    & SearchQuestionListStylingProps

export default ({ filters }: SearchQuestionListProps): JSX.Element => {
    const { response: lists, isLoading, error, send } = useFetchQuestionLists()

    useEffect(() => {
        send()
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