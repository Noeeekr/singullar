// Components
import Stack from "@mui/material/Stack"

// Models
import type { Question } from "@models/server"
import type { StackProps } from "@mui/material/Stack"
import type { JSX } from "react"

export interface QuestionDisplayerProps extends StackProps {
    data: Question
}
export default ({ data, ...props }: QuestionDisplayerProps): JSX.Element => {
    return (
        <Stack {...props}>
            { data.question_title }
        </Stack>    
    )
}