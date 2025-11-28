// Components
import QuestionContainer from "./components/Container"
import NumberDisplay from "./components/DifficultyBox";
import Stack from "@mui/material/Stack";

// Models
import styled from "@emotion/styled";
import Title from "./components/Title";
import Text from "./components/Text";
import { QuestionListDisplayProps } from ".";

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
                <NumberDisplay
                    amount={list.question_amount}
                    level={list.question_list_difficulty_level}
                />
            </Stack>
        </QuestionContainer>
    )
}