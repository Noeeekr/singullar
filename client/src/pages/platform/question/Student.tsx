import SectionHeader from "@components/headers/sectionHeader"
import ErrorBubble from "@components/bubbles/ErrorBubble"
import Typography from "@mui/material/Typography"
import NotFound from "./NotFound"
import Stack from "@mui/material/Stack"

import useFetchQuestions from "@hooks/useFetchQuestions"
import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import { parseQuestionDifficulty } from "@models/server/question"

import type { Question } from "@models/server"
import type { JSX } from "react"
import Alternatives from "@components/question/Alternatives"

export default (): JSX.Element => {
    const { response: questions, error, send } = useFetchQuestions()
    const [selected, setSelected] = useState("")

    const params = useParams()
    useEffect(() => {
        const question_id = Number(params["question_id"])
        send({
            question_ids: [question_id],
        })
    }, [])

    const question: Question | undefined = questions?.at(0)
    if (!question) return <NotFound />

    return (
        <Stack gap={3}>
            <SectionHeader
                title={error ? "Questão desconhecida" : question.question_title}
                subtitle={"Dificuldade: " + parseQuestionDifficulty(question.question_difficulty_level)}
            />
            <ErrorBubble message={error} />
            <Stack gap={1}>
                <Typography fontSize={14} color="rgb(127,123,125)">
                    {question.question_description}
                </Typography>
                <Typography fontSize={15}>
                    {question.question_short_description}
                </Typography>
                {
                    question.question_alternatives?.map((alternative, i) => {
                        return <Alternatives
                            selectable={selected == ""}
                            isSelected={selected == alternative}
                            title={alternative}
                            decorativeIconIndex={i}
                            onSelection={(s) => setSelected(s)}
                        />
                    })
                }
            </Stack>
        </Stack>
    )
}