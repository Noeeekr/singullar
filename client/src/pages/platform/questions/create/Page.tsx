import Stack from "@mui/material/Stack"
import Header from "./components/Header"
import Form from "./components/CreateForm"

export interface FormQuestionList {
    Name: string
    DifficultyName: string
    DifficultyLevel: number,
}

export default function(): JSX.Element {
    return(
        <Stack gap={2}>
            <Header/>
            <Form/>
        </Stack>
    )
}