import {
    Box,
    Typography,
    useMediaQuery
} from '@mui/material'
import {
    useTheme
} from '@mui/material/styles'

import MenuIcon from './NavbarButton'

interface IAppNavBarProps {
    menuButtonCallback: Function,
    children?: JSX.Element[],
    showMenu?: boolean
}

/**
 * Children are shown when navbar menu button is clicked.
 */
const AppNavbar = (props: IAppNavBarProps): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))
    
    const { 
        children = [],
        menuButtonCallback,
        showMenu,
    } = props;

    
    return (
        <Box
            sx={{
                width: "100%",
                height: "100%",
            }}
        >
        <Box
            sx={{
                position: "absolute",
                top: '0',
                left: '0',

                width: '100vw',
                height: 'auto',
                overflow: 'auto',

                zIndex: 2,
            }}
        >
            <Box
                display="flex"
                alignItems="center"
                boxSizing="border-box"
                paddingX={1}
                gap={1}
                width="100%"
                height={55}
                sx={{
                    backgroundColor: (theme) => theme.palette.primary.purpleDark
                }}
            >
                <MenuIcon
                    onClick={() => { menuButtonCallback() }}
                    notifications={true}
                />
                <Typography
                    component="h3"
                    variant="h5"
                    color="primary.light"
                    sx={{
                        fontSize: 22,
                        fontWeight: 'bold',
                    }}
                >Singullar</Typography>
            </Box>
            <Box>
            {
                isMobile && showMenu
                    ? (children.map((child) => {
                        return child
                    }))
                    : <></>
            }
            </Box>
        </Box>
        </Box>
    )
}

export default AppNavbar