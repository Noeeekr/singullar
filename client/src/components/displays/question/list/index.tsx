import { QuestionList } from "@models/server/question"
import type { StackProps } from "@mui/material/Stack"

export interface QuestionListDisplayProps extends QuestionListDisplayStylingProps {
    list: QuestionList
}

export interface QuestionListDisplayStylingProps extends StackProps {
    selectable?: boolean,
    navigable?: boolean,
}

export { default } from "./NavigableDisplay"