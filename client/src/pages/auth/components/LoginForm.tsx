// Components
import Typography from '@mui/material/Typography'
import FormControl from '@mui/material/FormControl'
import FormHelperText from '@mui/material/FormHelperText'
import OutlinedInput from '@mui/material/OutlinedInput'
import InputAdornment from '@mui/material/InputAdornment'
import InputLabel from '@mui/material/InputLabel'
import IconButton from '@mui/material/IconButton'
import ButtonBase from '@mui/material/ButtonBase'
import Button from '@mui/material/Button'
import Box from '@mui/material/Box'
import ErrorBubble from '@components/bubbles/ErrorBubble/ErrorBubble';
import ErrorHelperText from '@components/bubbles/ErrorBubble/ErrorHelperText';

// Icons
import { 
    Visibility, 
    VisibilityOff,
    HighlightOff
} from '@mui/icons-material'

// Features
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { useNavigate } from 'react-router-dom'

import { useTheme } from '@mui/material/styles'
import { useMediaQuery } from '@mui/system'

import useSignin from '@hooks/useSignin'

// Types
import type { SubmitHandler } from 'react-hook-form'


interface Fields {
    password: string
    email: string
}

const Internal = (): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))
    
    const { register, handleSubmit, formState: { errors }, reset, setValue, watch } = useForm<Fields>({
        defaultValues: {
            email: "",
            password: "",
        }
    })
    
    // Input update
    const emailValue = watch("email")
    const passwordValue = watch("password")

    const [showPassword, setShowPassword] = useState(false)
    
    // Fetch
    const navigate = useNavigate()
    
    const { signin, isLoading, signinError } = useSignin()

    const onSubmit: SubmitHandler<Fields> = async (fields) => {
        reset()

        const worked = await signin(fields.email, fields.password);
        if (worked) {
            navigate("/")
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

            <form onSubmit={handleSubmit(onSubmit)}>
                {
                    signinError != null && <ErrorBubble err={signinError} />
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
                        disabled={isLoading}
                        endAdornment={
                            <InputAdornment
                                position="end"
                                onClick={() => { setValue("email", "")}}
                            >
                                <IconButton
                                    edge="end"
                                >
                                    {
                                        emailValue && <HighlightOff sx={{ fontSize: 30, opacity: 0.3 }} />
                                    }
                                </IconButton>
                            </InputAdornment>
                        }
                    />
                    <ErrorHelperText show={Boolean(errors?.email)}>
                        Campo obrigatório. Digite seu email corretamente.
                    </ErrorHelperText>
                </FormControl>
                <FormControl margin="normal">
                    <InputLabel htmlFor="filled-adorment-password">
                        Digite sua senha
                    </InputLabel>
                    <OutlinedInput
                        {...register("password", {
                            required: "Por favor insira uma senha",
                            minLength: { 
                                value: 2, 
                                message: "A senha deve ter no mínimo 2 caracteres"
                            }
                        })}
                        type={showPassword ? 'text' : 'password'}
                        id="filled-adorment-password"
                        endAdornment={
                            <InputAdornment position="end">
                                <IconButton
                                    aria-label="mudar visibilidade da senha"
                                    onClick={() => (setShowPassword((show) => !show))}
                                    edge="end"
                                >
                                    {passwordValue 
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
                        error={Boolean(errors?.password)}
                    />
                    <ErrorHelperText show={Boolean(errors?.password)}>
                       {errors?.password?.message}
                    </ErrorHelperText>
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
                backgroundColor: 'white',
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