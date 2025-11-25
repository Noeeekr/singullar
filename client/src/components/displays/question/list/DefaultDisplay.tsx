// Components
import QuestionContainer from "./Container"
import NumberDisplay from "./DifficultyBox";
import Stack from "@mui/material/Stack";

// Models
import type { QuestionList } from "@models/server/question";
import type { StackProps } from "@mui/material"
import styled from "@emotion/styled";
import Title from "./Title";
import Text from "./Text";

export interface QuestionListDisplayProps extends QuestionListDisplayStylingProps {
    list: QuestionList
}

export interface QuestionListDisplayStylingProps extends StackProps { 
    selectable?: boolean,
    navigable?: boolean,
}

const Emphasis = styled("span")(({
    fontWeight: "bold",
}))

export default function ({ list, ...props }: QuestionListDisplayProps): JSX.Element {
    return (
        <QuestionContainer {...props}>
            <Stack direction="row" justifyContent="space-between">
                <Title>
                    {list.question_list_title}
                </Title>
                <Text textAlign="end">
                    <Emphasis>Matéria:</Emphasis> {list.subject_name}
                </Text>
            </Stack>
            <Text>
                <Emphasis>Dificuldade:</Emphasis> {list.question_list_difficulty_level}
            </Text>
            <Stack direction="row" gap={1} marginY={2}>
                {
                    new Array(list.question_amount).map((_, i) => (
                        <NumberDisplay key={i} level={list.question_list_difficulty_level}>{i + 1}</NumberDisplay>
                    ))
                }
            </Stack>
        </QuestionContainer>
    )
}