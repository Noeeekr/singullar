import { useState, useRef } from 'react'

import {
    Box,
    Stack,
    useMediaQuery,
} from '@mui/material'
import {
    useTheme
} from '@mui/material/styles'

import useSignout from '../../hooks/useSignout'
import {
    AppNavbar,
    SideMenu,
    UserProfile,
} from './components'

import { 
    Outlet,
    useNavigate,
    useLocation,
} from 'react-router-dom';

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
    
    // DELETE LATER
    const navigate = useNavigate()
    const url = useLocation().pathname;
    const deleteLaterNavigateHandler = () => {
        if (url === "/home") {
            navigate("/home/pip")
        } else {
            navigate("/home")
        }
    }
       
    return (
        <Box
            display="grid"
            gridTemplateColumns="1fr"
            gridTemplateRows="55px 1fr"
            height="auto"
            minHeight="100vh"
            sx={{
                backgroundColor: theme.palette.primary.paperLight,
                overflow: 'hidden',
            }}
        >
            <AppNavbar 
                showMenu={isOpen}
                menuButtonCallback={toggleIsOpen}
            >
                <UserProfile structure="left"/>
                <SideMenu isMobile={true}/>
            </AppNavbar>

            <Stack direction="row">
                {
                    !isMobile && <SideMenu onHoverOpen={toggleIsOpenRelative} isOpen={isOpen}/>
                }
                <Outlet />
                <Box sx={{
                    position: "absolute",
                    top: 150,
                    right: 0,

                    width: '25px',
                    height: 'auto',
                    borderRadius: '14px 0px 0px 14px',
                    padding: 1,
                    backgroundColor: 'rgba(90,60,180,0.8)',

                    cursor: 'pointer',
                    
                    textWrap: 'wrap',
                    textAlign: 'center',
                    textTransform: 'uppercase',
                    color: "white",
                }} onClick={deleteLaterNavigateHandler}>
                    D e v - C h a n g e - p a g e
                </Box>
            </Stack>
        </Box>
    )
}

export default Layout