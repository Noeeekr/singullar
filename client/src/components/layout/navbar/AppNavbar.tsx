import {
    useMemo,
    useState,
    useEffect,
} from 'react'
import { useTheme } from '@mui/material/styles'
import { NotificationContext } from '../../../context/notificationsContext'

import { useAppSelector } from '../../../slices/store'

// COMPONENTS
import MenuIcon from './Icon'

import {
    Box,
    Stack,
    Typography,
    useMediaQuery
} from '@mui/material'

// DATA
import { admin_navbar_popup_data } from '../sidemenu/data/admin_routes'
import { supervisor_navbar_popup_data } from '../sidemenu/data/supervisor_routes'
import { teacher_navbar_popup_data } from '../sidemenu/data/teacher_routes'
import { students_navbar_popup_data } from '../sidemenu/data/student_routes'
import Popup from '@components/buttons/Popup'

interface IAppNavBarProps {
    menuButtonCallback: () => void,
    children?: JSX.Element[],
    showMenu?: boolean
}

/**
 * Children are shown when navbar menu button is clicked.
 */

const AppNavbar = (props: IAppNavBarProps): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))

    const [hasNotifications, setHasNotifications] = useState(false)
    const [notifications, setNotifications] = useState<{
        "notifications": object[],
        "help": object[],
    }>({
        "notifications": [],
        "help": []
    });

    useEffect(() => {
        setHasNotifications(true)
        setNotifications({
            notifications: [{}],
            help: [{}],
        })
    }, [])

    const role = useAppSelector((state) => state.user.user?.role)

    const datas = useMemo(() => {
        switch (role) {
            case "student":
                return students_navbar_popup_data
            case "admin":
                return admin_navbar_popup_data
            case "supervisor":
                return supervisor_navbar_popup_data
            case "teacher":
                return teacher_navbar_popup_data
        }
    }, [role])

    const {
        children = [],
        menuButtonCallback,
        showMenu,
    } = props;

    return (
        <Box sx={{ zIndex: 10 }}>
            <Box
                top="0"
                left="0"
                width="100vw"
                sx={{
                    overflowX: isMobile ? "hidden" : "initial",
                    height: showMenu && isMobile ? '100vh' : 'auto',
                    overflowY: isMobile && showMenu ? 'scroll' : 'visible',
                    backgroundColor: 'white',
                }}
            >
                <NotificationContext.Provider value={{ notifications }}>
                    <Box
                        display="flex"
                        alignItems="center"
                        padding={1}
                        gap={1}
                        height={55}
                        width="100%"
                        sx={{
                            zIndex: 3,
                            backgroundColor: (theme) => theme.palette.primary.purpleLight
                        }}
                    >
                        <MenuIcon
                            onClick={() => { menuButtonCallback() }}
                            notifications={isMobile ? hasNotifications : false}
                        />
                        <Typography
                            component="h3"
                            variant="h5"
                            color="primary.light"
                            sx={{
                                fontSize: 22,
                                fontWeight: 'bold',
                            }}
                        >
                            Singullar
                        </Typography>

                        { /* Big screen */}

                        <Stack
                            direction="row"
                            alignItems="center"
                            justifyContent="center"
                            gap={1}
                            margin="auto 0 auto auto"
                        >
                            {
                                !isMobile && datas != null && datas.map((data, i) => {
                                    if (i === (datas.length - 1)) {
                                        return (
                                            <Popup
                                                id={data.id}
                                                type="popup"
                                                isCorner={true}
                                                variant={data.variant}
                                                title={data.title}
                                                icon={data.icon}
                                                key={data.title + data.id + data.variant}
                                            >
                                                {data.element}
                                            </Popup>
                                        )
                                    }
                                    return (
                                        <Popup
                                            id={data.id}
                                            type="popup"
                                            variant={data.variant}
                                            title={data.title}
                                            icon={data.icon}
                                            key={data.title + data.id + data.variant}
                                        >
                                            {data.element}
                                        </Popup>
                                    )
                                })
                            }
                        </Stack>
                    </Box>

                    { /* Mobile */}

                    <Box sx={{ zIndex: 1 }}>
                        {
                            isMobile && showMenu && (
                                children.map((child) => {
                                    return child
                                })
                            )
                        }
                    </Box>
                </NotificationContext.Provider>
            </Box>
        </Box>
    )
}

export default AppNavbar