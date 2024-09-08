import {
    Typography,
    FormControl,
    FormHelperText,
    OutlinedInput,
    InputAdornment,
    InputLabel,
    IconButton,
    ButtonBase,
    Button,
    Box
} from '@mui/material'
import {
    useTheme
} from '@mui/material/styles'
import {
    useMediaQuery
} from '@mui/system'
import {
    Visibility,
    VisibilityOff,
    HighlightOff
} from '@mui/icons-material'

import {
    useState
} from 'react'

import {
    useNavigate
} from 'react-router-dom'

const Internal = (): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))

    const [showPassword, setShowPassword] = useState(false)
    const [emailFieldValue, setEmailFieldValue] = useState("")

    const handleClickShowPassword = () => setShowPassword((show) => !show)
    const handleClickCleanInput = () => setEmailFieldValue("")

    const navigate = useNavigate()

    return (
        <Box component="section">
            <Typography
                component="p"
                variant="body2"
                marginBottom={2}
                color={isMobile ? "primary.dark" : "primary.semiLight"}
            >
                Digite seus dados de acesso para entrar
            </Typography>

            <form>
                <FormControl>
                    <InputLabel
                        htmlFor="text"
                    >
                        Insira seu e-mail
                    </InputLabel>
                    <OutlinedInput
                        id="text"
                        type="text"
                        label="Insira seu e-mail"
                        value={emailFieldValue}
                        onChange={(e) => setEmailFieldValue(e.target.value)}
                        endAdornment={
                            <InputAdornment
                                position="end"
                                onClick={handleClickCleanInput}
                            >
                                <IconButton
                                    edge="end"
                                >
                                    {emailFieldValue
                                        ? <HighlightOff sx={{ opacity: 0.5 }} />
                                        : <></>
                                    }
                                </IconButton>
                            </InputAdornment>
                        }
                    />
                </FormControl>
                <FormControl
                    margin="normal"
                >
                    <InputLabel
                        htmlFor="filled-adorment-password"
                    >
                        Digite sua senha
                    </InputLabel>
                    <OutlinedInput
                        id="filled-adorment-password"
                        type={showPassword ? 'text' : 'password'}
                        endAdornment={
                            <InputAdornment position="end">
                                <IconButton
                                    aria-label="mudar visibilidade da senha"
                                    onClick={handleClickShowPassword}
                                    edge="end"
                                >
                                    {showPassword ? <Visibility sx={{ opacity: '0.5' }} /> : <VisibilityOff sx={{ opacity: '0.5' }} />}
                                </IconButton>
                            </InputAdornment>
                        }
                        label="Digite sua senha"
                    />
                    <FormHelperText
                        variant="outlined"
                        sx={{
                            textAlign: 'end'
                        }}
                    >
                        <ButtonBase
                            disableRipple={true}
                            sx={{
                                color: 'primary.darkPurple',
                                textDecoration: 'underline',
                                fontFamily: 'Verdana',
                            }}
                            onClick={() => navigate("/recover")}
                        >
                            Esqueci minha senha
                        </ButtonBase>
                    </FormHelperText>
                </FormControl>
                <Box sx={{
                    display: 'flex',
                    alignItems: 'center',
                    marginTop: 2,
                    height: 44
                }}>
                    <Button
                        type="submit"
                        sx={{
                            scale: 1,
                            height: 38,
                            width: '100%',
                            color: 'white',
                            fontWeight: 'bold',
                            borderRadius: '40px',
                            fontFamily: 'Verdana',
                            textTransform: 'Capitalize',
                            transition: 'all 180ms ease-in-out',
                            backgroundColor: (theme) => `${theme.palette.primary.contrast}`,
                            '&:hover': {
                                scale: '1.005 1.1'
                            },
                        }} // the background color gives an err if not in ```probably because it is string | undefined
                    >
                        Entrar
                    </Button>
                </Box>
            </form>
        </Box>
    )
}

const Entire = (): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))

    return (
        <Box
            sx={{
                display: 'flex',
                flexDirection: 'column',
                backgroundColor: 'rgb(255,255,255)',
                padding: isMobile ? 2 : 3,
                borderRadius: 4,
                flex: 1
            }}
            component="section"
        >
            <Typography
                variant="h4"
                marginBottom={0.2}
            >
                Login
            </Typography>
            <Internal />
        </Box>
    )
}

const LoginForm = ({ structure }: { structure: "internal" | "entire" }): JSX.Element => {

    return structure === "internal"
        ? <Internal />
        : <Entire />
}

export default LoginForm