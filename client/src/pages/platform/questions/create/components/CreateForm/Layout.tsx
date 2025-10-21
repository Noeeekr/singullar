import Stack from "@mui/material/Stack"
import Header from "./Header"
import Form from "./Form"

export default function(): JSX.Element {
    return(
        <Stack gap={2}>
            <Header activeStep={1}/>
            <Form />
        </Stack>
    )
}