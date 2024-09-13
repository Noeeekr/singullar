import {
    Box,
    Button,
    Divider,
    Typography,
    ButtonGroup,
} from '@mui/material'
import {
    ArrowRight,
    Brightness4,
    Brightness5,
    Brightness6,
    Brightness7,
} from '@mui/icons-material'
import {
    styled,
    useMediaQuery
} from '@mui/system'
import {
    Theme,
    useTheme,
} from '@mui/material/styles'

const HoverBox = styled(Box)({
    display: 'flex',
    flexDirection: "column",
    gap: 1,
    paddingTop: 6.5,
    paddingBottom: 6.5,
    '&:hover .MuiButton-root': {
        backgroundColor: 'rgba(230,230,230,0.4)'
    },
    '& .MuiButtonGroup-grouped': {
        border: 'none',
        boxShadow: 'none',
        borderRadius: 2,
    }
})

const ListButton = styled(Button)(({ theme }) => ({
    color: theme.palette.primary.lightGray,
    justifyContent: 'start',
    height: '38px',
    margin: '10px 0px',
    fontFamily: 'Helvetica',
    fontSize: 14,
    textTransform: 'capitalize',
    '& .MuiTouchRipple-root .MuiTouchRipple-rippleVisible': {
        color: 'rgba(150,150,150,0.2)', // Ripple color on click
    },
    '&:hover': {
        backgroundColor: 'rgba(200,200,200,0.2)', // Button hover color
        '& .MuiTouchRipple-root .MuiTouchRipple-rippleVisible': {
            color: 'rgba(200,0,150,0.2)', // Ripple color on hover
        },
    },
}))


interface BaseButtonGroupItems {
    label: string
    icon: JSX.Element
}
interface ButtonGroupRedirect extends BaseButtonGroupItems {
    type: "url"
    url: string
}
interface ButtonGroupPopup extends BaseButtonGroupItems {
    type: "component"
    component: JSX.Element
}


const buttonGroupItems: (ButtonGroupRedirect | ButtonGroupPopup)[] = [
    {
        label: "Alunos",
        type: "url",
        url: "/changeThisLater39393939",
        icon: <Brightness4/>
    },
    {
        label: "Responsavel",
        type: "component",
        component: <div>Change This Later</div>,
        icon: <Brightness5 />
    },
    {
        label: "Professor e Equipe escolar",
        type: "component",
        component: <div>Change This Later</div>,
        icon: <Brightness6 />
    },
    {
        label: "Código de material",
        type: "component",
        component: <div>Change This Later</div>,
        icon: <Brightness7 />
    },
]

const Internal = (): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))

    return (
        <>
            <Typography
                component="p"
                variant="body1"
                marginBottom={2}
                color={ isMobile ? "primary.dark": "primary.semiLight" }
            >
                Selecione uma opção para se cadastrar
            </Typography>

            <Box padding={1} >
                <ButtonGroup
                    variant="text"
                    orientation="vertical"
                    fullWidth={true}
                    sx={{
                        '& .MuiButtonGroup-grouped': {
                            border: 'none',
                            margin: '0',
                        },
                    }}
                >
                    {buttonGroupItems.map((item, i) => {
                        const Button: () => JSX.Element = () => (
                            <HoverBox
                                sx={(theme: Theme) => ({
                                '& .MuiSvgIcon-root': {
                                    color: theme.palette.primary.contrast
                                },
                                '&:hover .MuiSvgIcon-root': {
                                    color: theme.palette.primary.lightGray,
                                    opacity: 0.6,
                                }
                                })}
                            >
                                <ListButton 
                                    sx={{ justifyContent: 'space-between', }}
                                    endIcon={<ArrowRight sx={{
                                        fontSize: "25px !important" 
                                    }}/>}
                                >
                                    <Box sx={{ display: 'flex', gap: 2, }} >
                                        {item.icon}
                                        {item.label}
                                    </Box>
                                </ListButton>
                            </HoverBox>
                        )

                        return i == 0
                            ? <Button key={item.label}/>
                            : <div key={item.label}>
                                <Divider />
                                <Button />
                            </div>
                    })}
                </ButtonGroup>
            </Box>
        </>

    )
}

const Entire = (): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))

    return (
        <Box
            padding={isMobile ? 2 : 3}
            sx={{
                display: 'flex',
                flexDirection: 'column',
                backgroundColor: 'primary.main',
                borderRadius: 4,
                flex: 1
            }}
            component="section"
        >
            <Typography
                variant="h6"
                marginBottom={0.2}
            >
                Cadastro
            </Typography>
            <Internal />
        </Box>
    )
}

const RegisterForm = ({ structure }: { structure: "internal" | "entire" }): JSX.Element => {

    return structure === "internal"
        ? <Internal />
        : <Entire />

}

export default RegisterForm