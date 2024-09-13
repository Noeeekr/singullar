import {
    Box,
    Stack,
    Typography,
    useMediaQuery,
} from '@mui/material'
import {
    useTheme    
} from '@mui/material/styles'

import FormSelection from './components/FormSelection'
import LoginForm from './components/LoginForm'
import RegisterForm from './components/RegisterForm'
import FooterLinks from './components/FooterLinks'

const AuthPage = (): JSX.Element => {
    const theme = useTheme()

    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))
    const isBigScreen = useMediaQuery(theme.breakpoints.up('md'))
    const isLargerScreen = useMediaQuery(theme.breakpoints.up('lg'))

    return (
        <Stack
            spacing={4}
            component="div"
            sx={{
                backgroundColor: (theme) => theme.palette.primary.darkPurple,
                minHeight: '100vh',
                paddingY: isLargerScreen ? 4 : 2,
                paddingX: isLargerScreen ? 14 : 2,
                overflow: 'hidden',
            }}
        >
            <Typography
                component="h1"
                variant="h5"
                color="primary.light"
                sx={{
                    fontSize: 32,
                    fontWeight: 'bold',
                    textAlign: isMobile ? 'center' : "start",
                    paddingY: isMobile ? 2 : 5,
                }}
            >RandName</Typography>

            {
                isMobile
                    ? <FormSelection />
                    : (
                        <Stack
                            component="div"
                            spacing={5}
                            direction="row"
                            sx={{
                                marginBottom: isLargerScreen ? "30px !important" : 0
                            }}
                        >
                            <Stack
                                spacing={3}
                                component="main"
                                direction={ !isBigScreen ? "column" : "row" }
                                sx={{
                                    alignItems: isBigScreen ? "start" : "initial",
                                    flex: 7,
                                }}
                            >
                                <LoginForm structure="entire" />
                                <RegisterForm structure="entire" />
                            </Stack>
                            <Box sx={{
                                flex: 2.5
                            }}>
                                mascot image if I had one.jpeg
                            </Box>
                        </Stack>
                    )
            }
            <FooterLinks/>
            { /* Popups */}
        </Stack>
    )
}

export default AuthPage