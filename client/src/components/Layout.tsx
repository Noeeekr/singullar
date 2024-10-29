// Features
import { useState } from 'react'
import { useTheme } from '@mui/material/styles'

// Components
import Box from '@mui/material/Box'
import Stack from '@mui/material/Stack'
import useMediaQuery from '@mui/material/useMediaQuery'

import {
    AppNavbar,
    SideMenu,
    UserProfile,
} from './index'

import { Outlet } from 'react-router-dom';

const Layout = (): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))

    const [isOpen, setIsOpen] = useState(false)

    return (
        <Box
            display="grid"
            gridTemplateColumns="1fr"
            gridTemplateRows="55px 1fr"
            height="100vh"
            sx={{
                backgroundColor: theme.palette.primary.paperLight,
                overflow: isMobile ? 'scroll' : 'hidden',
            }}
        >
            <AppNavbar
                showMenu={isOpen}
                menuButtonCallback={() => (setIsOpen(prevState => !prevState))}
            >
                <UserProfile structure="left" />
                <SideMenu isMobile={true} />
            </AppNavbar>

            <Stack direction="row" sx={{ position: 'relative', width: '100vw', overflow: 'scroll' }}>
                { 
                    !isMobile && <SideMenu isOpen={isOpen}/>
                }
                <main
                    style={{
                        width: '100%',
                        padding: isMobile ? '80px 20px 60px 25px' : '80px 20px 30px 120px',
                    }}
                >
                    <Outlet />
                </main>

            </Stack>
        </Box>
    )
}

export default Layout