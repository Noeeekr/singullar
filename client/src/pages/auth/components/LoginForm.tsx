import { useState, ReactNode } from 'react'
import { useNavigate } from 'react-router-dom'
import { useForm, SubmitHandler } from 'react-hook-form'

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
    Stack,
    Box,
    Fade
} from '@mui/material'
import { 
    Visibility, 
    VisibilityOff,
    HighlightOff
} from '@mui/icons-material'
import { useTheme } from '@mui/material/styles'
import { useMediaQuery } from '@mui/system'

import useSignIn from '../../../hooks/useSignin'

interface Fields {
    password: string
    email: string
}

const ErrorHelperText = (
    { children }: { children: ReactNode }
): JSX.Element => {

    return (
        <Box
            sx={{
                position: 'absolute',
                bottom: -25,
                backgroundColor: (theme) => (theme.palette.error.whiteHigh),
                paddingX: 2,
                paddingTop: 2,
                paddingBottom: 0.5,
                borderEndEndRadius: 10,
                borderEndStartRadius: 10,
                zIndex: 1,
                width: '100%',
            }}
        >
            <Typography
                variant="body1"
                color="error.dark"
            >
                {children}
            </Typography>
        </Box>
    )
}
const ErrorBubble = (
    { err }: { err: string }
): JSX.Element => {

    return (
        <Box
            sx={{
                marginY: 2,
                padding: 2,
                borderRadius: 3.5,
                backgroundColor: (theme) => `${theme.palette.error.light}`,
            }}>
            <Stack
                direction="row"
                spacing={1}
                sx={{
                    alignItems: 'center',
                }}
            >
                <Box
                    sx={{
                        position: 'relative',
                        display: 'flex',
                        '&:before': {
                            position: 'absolute',
                            top: -16,
                            content: '""',
                            backgroundColor: (theme) => theme.palette.error.main,
                            width: '80%',
                            height: 5,
                            marginLeft: '2px',
                            borderRadius: 2,
                        }
                    }}
                >

                    <HighlightOff
                        color="error"
                    />
                </Box>
                <Typography
                    color="error.dark"
                    variant="body1"
                    sx={{
                        fontWeight: 'bold',
                    }}
                >
                    {err}
                </Typography>
            </Stack>
        </Box>
    )
}

const Internal = (): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))
    
    const { register, handleSubmit, formState: { errors } } = useForm<Fields>({
        defaultValues: {
            email: "",
            password: "",
        }
    })

    // Input update
    const [showPassword, setShowPassword] = useState(false)
    const [formFields, setFormFields] = useState<Fields>({
        email: "",
        password: "",
    })

    const handleClickCleanInput = () => setFormFields((prevState: Fields) => ({ ...prevState, email: "" }))
    
    // Fetch
    const navigate = useNavigate()

    const [signin, isLoading, loginError] = useSignIn()

    const onSubmit: SubmitHandler<Fields> = async (fields) => {
        setFormFields({
            password: "",
            email: ""
        })

        await signin(fields.email, fields.password)

        if(!loginError) {
            navigate("/home")
        }
    }

    return (
        <Box component="section">
            <Typography
                component="p"
                variant="body2"
                marginBottom={2}
                color={isMobile ? "primary.whiteNone" : "primary.whiteMedium"}
            >
                Digite seus dados de acesso para entrar
            </Typography>

            <form noValidate onSubmit={handleSubmit(onSubmit)}>
                {
                    loginError != null && <ErrorBubble err={loginError} />
                }
                <FormControl>
                    <InputLabel htmlFor="text">
                        Insira seu e-mail
                    </InputLabel>
                    <OutlinedInput
                        {...register("email", {
                            required: "Por favor insira um e-mail.",
                            pattern: {
                                value: /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
                                message: 'Por favor insira um email válido.',
                            },
                        })}
                        error={Boolean(errors?.email)}
                        id="text"
                        type="text"
                        label="Insira seu e-mail"
                        value={formFields.email}
                        disabled={isLoading}
                        onChange={(e) => setFormFields((prevState: Fields) => (
                            {
                                email: e.target.value,
                                password: prevState.password
                            }
                        ))}
                        endAdornment={
                            <InputAdornment
                                position="end"
                                onClick={handleClickCleanInput}
                            >
                                <IconButton
                                    edge="end"
                                >
                                    {
                                        formFields.email && <HighlightOff sx={{ fontSize: 30, opacity: 0.3 }} />
                                    }
                                </IconButton>
                            </InputAdornment>
                        }
                    />
                    <Fade in={Boolean(errors?.email)}>
                        <div style={{ position: 'relative', marginBottom: errors?.email ? 20 : 1 }}>
                            <ErrorHelperText>
                                Campo obrigatório. Digite seu email corretamente.
                            </ErrorHelperText>
                        </div>
                    </Fade>
                </FormControl>
                <FormControl margin="normal">
                    <InputLabel htmlFor="filled-adorment-password">
                        Digite sua senha
                    </InputLabel>
                    <OutlinedInput
                        {...register("password", {
                            required: "por favor insira uma senha",
                        })}
                        error={Boolean(errors?.password)}
                        id="filled-adorment-password"
                        type={showPassword ? 'text' : 'password'}
                        endAdornment={
                            <InputAdornment position="end">
                                <IconButton
                                    aria-label="mudar visibilidade da senha"
                                    onClick={() => (setShowPassword((show) => !show))}
                                    edge="end"
                                >
                                    {formFields.password 
                                        ? showPassword 
                                            ? <Visibility sx={{ fontSize: 30, opacity: '0.3' }} />
                                            : <VisibilityOff sx={{ fontSize: 30, opacity: '0.3' }} />
                                        : <></>
                                    } 
                                </IconButton>
                            </InputAdornment>
                        }
                        label="Digite sua senha"
                        disabled={isLoading}
                        onChange={(e) => setFormFields((prevState: Fields) => (
                            {
                                password: e.target.value,
                                email: prevState.email
                            }
                        ))}
                        value={formFields.password}
                    />
                    <Fade in={Boolean(errors?.password)}>
                        <div style={{ position: 'relative', marginBottom: errors?.password ? 20 : 1 }}>
                            <ErrorHelperText>
                                Campo obrigatório. Digite sua senha.
                            </ErrorHelperText>
                        </div>
                    </Fade>
                    <FormHelperText sx={{ textAlign: 'end' }}>
                        <ButtonBase
                            disableRipple={true}
                            sx={{
                                color: 'primary.purpleDark',
                                textDecoration: 'underline',
                                fontFamily: 'Verdana',
                            }}
                            onClick={() => navigate("/recover")}
                        >
                            Esqueci minha senha
                        </ButtonBase>
                    </FormHelperText>
                </FormControl>
                <Box 
                    marginTop={2}
                    display="flex"
                    alignItems="center"
                    height={44}
                >
                    <Button
                        type="submit"
                        disabled={isLoading}
                        fullWidth={true}
                        color="primary"
                        sx={{
                            fontWeight: "bold",
                            fontFamily: "Verdana",
                            borderRadius: 20,
                            textTransform: 'Capitalize',
                            transition: 'all 180ms ease-in-out',
                            backgroundColor: `primary.contrast`,
                            '&:hover': {
                                scale: '1.005 1.1'
                            },
                            '&.Mui-disabled': {
                                backgroundColor: 'rgb(212,220,214)'
                            }
                        }} 
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
            padding={isMobile ? 2 : 3}
            display="flex"
            flexDirection="column"
            borderRadius={4}
            flex={1}
            sx={{
                backgroundColor: 'primary.whiteHigh',
            }}
            component="section"
        >
            <Typography
                variant="h6"
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