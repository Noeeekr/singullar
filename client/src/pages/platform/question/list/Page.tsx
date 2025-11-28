import QuestionListDisplayer from "@components/displays/question/list"
import SectionHeader from "@components/headers/sectionHeader"
import ErrorBubble from "@components/bubbles/ErrorBubble"
import Typography from "@mui/material/Typography"
import Divider from "@mui/material/Divider"
import Stack from "@mui/material/Stack"

// Utilities
import useFetchQuestionLists from "@hooks/useFetchQuestionLists"
import { useEffect } from "react"
import { useParams } from "react-router-dom"

export default (): JSX.Element => {
    const { response: list, error, send, isLoading } = useFetchQuestionLists()

    const { list_id } = useParams()


    useEffect(() => {
        // "List_id" could be "abcd", "" or undefined, which isn't parseable to number. Hence the validation.
        if (Number.isNaN(Number(list_id))) return
        send({
            filters: {
                fields: [],
                ids: [Number(list_id)],
            },
            offset: 0,
        })
    }, [])

    return (
        <Stack component="section" gap={2}>
            <SectionHeader title="Lista de questões" subtitle="Veja a descrição da lista de questões" />
            <Divider/>
            {
                list?.length
                    ? <QuestionListDisplayer variant="lg" list={list[0]} />
                    : error
                        ? <ErrorBubble message={error} />
                        : isLoading
                            ? <Typography> Carregando lista de questões... </Typography>
                            : <Typography> Lista de questões indisponível. </Typography>
            }
        </Stack>
    )
}