import { Fragment } from "react"
import Typography from "@mui/material/Typography"
import ErrorBubble from "@components/bubbles/ErrorBubble"

// Utilities
import QuestionListDisplayer from "@components/displays/question/list"
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
        <Fragment>
            {
                list?.length
                    ? <QuestionListDisplayer list={list[0]} />
                    : error
                        ? <ErrorBubble message={error} />
                        : isLoading
                            ? <Typography> Carregando lista de questões... </Typography>
                            : <Typography> Lista de questões indisponível. </Typography>
            }
        </Fragment>
    )
}