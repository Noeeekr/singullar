import { useState, useRef } from 'react'

import {
    Box,
    Stack,
    useMediaQuery,
} from '@mui/material'

import {
    useTheme
} from '@mui/material/styles'

import {
    AppNavbar,
    SideMenu,
    UserProfile,
} from './components'

import { Outlet } from 'react-router-dom';

const Layout = (): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))
    
    const [isOpen, setIsOpen] = useState(false)
    const isAbsoluteOpen = useRef(false)

    /**
     * Nav button has the primary will to toggle menu state
     */
    const toggleIsOpen = () => (setIsOpen(prevState => {
        isAbsoluteOpen.current = !prevState;
        return !prevState
    }))

    /**
     * Only changes menu state if the main toggler is not using it
     */
    const toggleIsOpenRelative = () => (setIsOpen(prevState => {
        if (isAbsoluteOpen.current) return true;
        return !prevState;
    }))

    return (
        <Box
            sx={{
                display: 'grid',
                gridTemplateColumns: "1fr",
                gridTemplateRows: '55px 1fr',

                backgroundColor: theme.palette.primary.paperLight,
                height: '100vh',
                overflow: 'hidden',
            }}
        >
            <AppNavbar 
                showMenu={isOpen}
                menuButtonCallback={toggleIsOpen}
            >
                <UserProfile position="left" bgColor="purple" />
                <SideMenu isMobile={true}/>
            </AppNavbar>

            <Stack direction="row">
                {
                    !isMobile && <SideMenu onHoverOpen={toggleIsOpenRelative} isOpen={isOpen}/>
                }
                <Outlet />
            </Stack>
        </Box>
    )
}

export default Layout