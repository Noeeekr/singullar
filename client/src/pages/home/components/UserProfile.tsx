import {
    Box,
    Stack,
    Avatar,
    Typography
} from '@mui/material'

import useAuth from '../../../hooks/useAuth'

interface InternalLayout {
    position: "center" | "left"
    bgColor: "purple" | "white"
}

const UserProfile = (props: InternalLayout): JSX.Element => {
    const { user } = useAuth();

    const { position, bgColor } = props;

    return (
        <Stack 
            direction="row"
            paddingX={1.7}
            paddingY={2.5}
            gap={2}
            alignItems="center"
            sx={{
                backgroundColor: (theme) => {
                    return bgColor == "white"
                        ? theme.palette.primary.whiteHigh
                        : theme.palette.primary.purpleDark
                }
            }}
        >
            
            <Box>
                <Avatar
                    alt={user?.name || ""}
                    sx={{
                        bgcolor: 'rgb(240,240,240)',
                        color: (theme) => theme.palette.primary.purpleDark,
                        fontWeight: 'bold',
                        width: 48,
                        height: 48,
                    }}
                >{ user?.name[0] || "</>"}</Avatar>
            </Box>
            <Stack
                gap={0.1}
            >
                    <Typography 
                        component="h3"
                        variant="h6"
                        sx={{ 
                            fontFamily: `Public Sans Web, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol;`,
                            fontWeight: 'semiBold', 
                            color: 'white'}}
                    >
                        Olá, {user?.name || "" } { user?.surname || ""}
                    </Typography>
                    <Typography 
                        variant="subtitle2"
                        sx={{ 
                            fontFamily: `Public Sans Web, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol`,
                            color: 'rgba(255,255,255,0.86)',
                        }}
                    >
                        {user?.email || "email@desconhecido.com" }
                    </Typography>
                    <Typography 
                        sx={{ 
                            fontSize: 14,
                            fontWeight: 'bold',
                            fontFamily: "Arial",
                            color: 'rgba(255,255,255,0.86)',
                        }}
                    >
                        Plataforma ID {user?.id ? Number(user?.id) + 1000000 : "Plataforma ID: Desconhecido" }
                    </Typography>
            </Stack>
        </Stack>
    )
}

export default UserProfile