import {
    Stack,
    Box,
} from '@mui/material'
import GlowButtonStack from './GlowButtonStack'

import LoginForm from './LoginForm'
import RegisterForm from './RegisterForm'

import {
    useState
} from 'react'

type Forms = "Login" | "Cadastro"

const forms: { [index in Forms]: JSX.Element } = {
    "Login": <LoginForm structure="internal"></LoginForm>,
    "Cadastro": <RegisterForm structure="internal"></RegisterForm>
}

const FormSelection = (): JSX.Element => {
    const [formChoice, setFormChoice] = useState<Forms>("Login")

    return (
        <Box
            sx={{
                display: 'flex',
                flexDirection: 'column',
                backgroundColor: 'rgb(255,255,255)',
                paddingY: 3,
                minHeight: 30 * 10,
                borderRadius: 4,
                overflow: 'auto',
            }}
            component="main"
            maxWidth="md"
        >
            <Stack
                sx={{
                    paddingX: 1.9,
                    overflow: 'hidden'
                }}
            >
                <GlowButtonStack
                    sx={{
                        padding: 0.2,
                    }}
                    active={formChoice}
                    labels={Object.keys(forms)}
                    onClick={setFormChoice}
                />
                {forms[formChoice]}
            </Stack>        
        </Box>
    )
}

export default FormSelection