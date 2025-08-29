import {
    Box,
    Stack,
    Avatar,
    Typography
} from '@mui/material'

import { useAppSelector } from '../slices/store'

interface IUserProfileProps {
    structure?: "center" | "left"
}

const UserProfile = (props: IUserProfileProps): JSX.Element => {
    const user = useAppSelector((state) => state.user.user);

    const { structure } = props;

    console.log(user)
    return (
        <Stack 
            direction={structure === "center" ? 'column' : 'row'}
            paddingX={1.7}
            paddingY={2.5}
            minWidth={300}
            gap={1.5}
            alignItems="center"
            sx={{
                backgroundColor: (theme) => {
                    return structure === "center"
                        ? "white"
                        : theme.palette.primary.purpleLight
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
                gap={0.2}
                sx={{
                    textAlign: structure,
                    color: structure === "center" ? 'rgb(0,0,0)' : 'white'
                }}
            >
                    <Typography 
                        component="h3"
                        variant={structure === "center" ? 'h4' : 'h5'}
                        sx={{ 
                            color: structure === "center" ? 'rgba(0,0,0,0.8)' :'rgba(255,255,255,0.86)',
                            fontFamily: `Public Sans Web, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol;`,
                            fontWeight: 700,
                            marginBottom: structure === "center" ? 0.5 : 0,
                        }}
                    >
                        Olá, {user?.name || "" }
                    </Typography>
                    <Typography 
                        variant="subtitle2"
                        sx={{ 
                            fontWeight: structure === "center" ? 400 : 500,
                            color: structure === "center" ? 'rgba(0,0,0,0.8)' :'rgba(255,255,255,0.86)',
                        }}
                    >
                        {user?.email || "email@desconhecido.com" }
                    </Typography>
                    <Typography 
                        sx={{ 
                            fontSize: 14,
                            fontWeight: structure === 'center' ? 400 : 600,
                            fontFamily: "Arial",
                            color: structure === "center" ? 'rgba(0,0,0,0.8)' :'rgba(255,255,255,0.86)',
                        }}
                    >
                        Plataforma ID {user?.id || "Plataforma ID: Desconhecido" }
                    </Typography>
            </Stack>
        </Stack>
    )
}

export default UserProfile