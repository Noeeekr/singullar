import {
    useMemo,
    useState,
    useEffect,
    useCallback,
} from 'react'
import { useTheme } from '@mui/material/styles'
import { NotificationContext } from '../context/notificationsContext'

import { useAppSelector } from '../slices/store'

// COMPONENTS
import NavbarItemPopup from './PopupIconButton'
import MenuIcon from './NavbarIconButton'

import {
    Box,
    Stack,
    Typography,
    useMediaQuery
} from '@mui/material'

// DATA
import { admin_navbar_popup_data, AdminNavbarPopupId } from '../pages/admin/data'
import { supervisor_navbar_popup_data, SupervisorNavbarPopupId } from '../pages/supervisor/data'
import { teacher_navbar_popup_data, TeacherNavbarPopupId } from '../pages/teacher/data'
import { students_navbar_popup_data, StudentNavbarPopupId } from '../pages/home/data'

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
    },[])

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
            default:
                return []
        }
    }, [role])

    type PopupKeys =
        typeof role extends "student" ? StudentNavbarPopupId
        : typeof role extends "admin" ? AdminNavbarPopupId
        : typeof role extends "supervisor" ? SupervisorNavbarPopupId
        : typeof role extends "teacher" ? TeacherNavbarPopupId
        : "";

    const [isOpen, setIsOpen] = useState<PopupKeys>("")

    // could become a switch
    /**
    *   Closes other popups when opening a new one
    */
    const toggleIsOpen = useCallback((key: PopupKeys) => {
        setIsOpen((isOpen) => key === isOpen ? "" : key)
    }, [])

    const {
        children = [],
        menuButtonCallback,
        showMenu,
    } = props;

    return (
        <Box sx={{ zIndex: 4 }}>
            <Box
                top="0"
                left="0"
                width="100vw"
                sx={{
                    height: showMenu && isMobile ? '100vh' : 'auto',
                    overflow: isMobile ? 'scroll' : 'visible',
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
                        sx={{
                            backgroundColor: (theme) => theme.palette.primary.purpleDark
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
                                !isMobile && datas.map((data, i) => {

                                    if (i === (datas.length - 1)) {
                                        return (
                                            <NavbarItemPopup
                                                isCorner={true}
                                                structure={data.structure}
                                                title={data.title}
                                                id={data.id}
                                                isOpen={isOpen}
                                                onClickCb={toggleIsOpen}
                                                icon={data.icon}
                                                key={data.title + data.id + data.structure}
                                            >
                                                {data.content}
                                            </NavbarItemPopup>
                                        )

                                    }
                                    return (
                                        <NavbarItemPopup
                                            structure={data.structure}
                                            title={data.title}
                                            id={data.id}
                                            isOpen={isOpen}
                                            onClickCb={toggleIsOpen}
                                            icon={data.icon}
                                            key={data.title + data.id + data.structure}
                                        >
                                            {data.content}
                                        </NavbarItemPopup>
                                    )
                                })
                            }
                        </Stack>
                    </Box>

                    { /* Mobile */}

                    <Box>
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