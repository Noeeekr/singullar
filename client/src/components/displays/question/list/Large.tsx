// Components
import ErrorBubble from "@components/bubbles/ErrorBubble";
import Stack from "@mui/material/Stack";
import Title from "./components/Title";
import Text from "./components/Text";
import Grid from "@mui/material/Grid2";

// Utilities
import useFetchQuestions from "@hooks/useFetchQuestions";
import { useEffect } from "react";
import { styled } from "@mui/material";

// Models
import type { QuestionListDisplayProps } from ".";
import SmallQuestionDisplayer from "../SmallQuestionDisplayer";

const Emphasis = styled("span")(({
    fontWeight: "bold",
}))

export default function ({ list, ...props }: QuestionListDisplayProps): JSX.Element {
    const { response: questions, error, send } = useFetchQuestions()

    useEffect(() => {
        send({ question_list_ids: [list.id] })
    }, [])

    return (
        <Stack gap={2} {...props}>
            <Stack>
                <Stack direction="row" justifyContent="space-between">
                    <Title fontSize={24}>
                        {list.question_list_title}
                    </Title>
                    <Text textAlign="end" fontSize={16}>
                        <Emphasis>Matéria:</Emphasis> {list.subject_name}
                    </Text>
                </Stack>
                <Text fontSize={16}>
                    <Emphasis>Dificuldade:</Emphasis> {list.question_list_difficulty_level}
                </Text>
            </Stack>
            <ErrorBubble message={error} />
            <Grid container spacing={2} marginY={2}>
                <Grid size={6}>
                    {
                        questions?.map((question, i) => (
                            <SmallQuestionDisplayer question={question} key={i} />
                        ))
                    }
                </Grid>
            </Grid>
        </Stack>
    )
}