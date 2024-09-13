import {
    Stack,
    Divider,
} from '@mui/material'
import {
    useMediaQuery
} from '@mui/system'

import {
    Link
} from 'react-router-dom'

interface FooterLink {
    label: string
    url: string
}

const links: FooterLink[] = [
    {
        label: 'Conheça a equipe',
        url: '/dummyRoute',
    },
    {
        label: 'Termos de Uso',
        url: '/dummyRoute',
    },
    {
        label: 'Portal de Privacidade',
        url: '/dummyRoute',
    },
    {
        label: 'Status da Plataforma',
        url: '/dummyRoute',
    },
    {
        label: 'Perguntas frequentes',
        url: '/dummyRoute',
    },
    {
        label: 'Ajuda',
        url: '/dummyRoute',
    }
]

const FooterLinks = (): JSX.Element => {
    const isMobile = useMediaQuery(`(max-width: 760px)`)

    return (
        <Stack
            direction="row"
            spacing={0.75}
            color="white"
            maxWidth="850px"
            paddingY={3}
            sx={{
                justifyContent: 'space-around',
                alignItems: isMobile ? 'center' : 'end',
                minWidth: 525,
            }}
        >
            {
                links.map((link) => (
                        <Link
                            to={link.url}
                            key={link.label}
                            style={{
                                fontSize: 15,
                                color: 'white',
                                paddingBottom: '3px',
                                textDecoration: 'underline',
                                borderBottom: isMobile ? 'none' : 'solid 1px white',
                                textAlign: isMobile ? 'center' : 'start',
                            }}
                        >
                            {link.label}
                        </Link>
                ))
            }
        </Stack>
    )
}

export default FooterLinks