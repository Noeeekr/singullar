// Features
import { styled } from '@mui/material'
import { useAppSelector } from '../../../slices/store';
import {
    useMemo,
    Fragment,
} from 'react'

// Types
import type { BoxProps } from '@mui/material'

// Components
import { Stack } from '@mui/material'
import MenuItems from './SideMenuButtons'

// Data 
import { admin_sidemenu_data } from './data/admin_routes'
import { supervisor_sidemenu_data } from './data/supervisor_routes'
import { teacher_sidemenu_data } from './data/teacher_routes'
import { students_sidemenu_data } from './data/student_routes'

interface SideMenuProps extends BoxProps {
    isMobile?: boolean,
    isOpen?: boolean,
    onHoverOpen?: () => void,
}

const Layout = styled('nav')<SideMenuProps>(({ theme, isMobile, isOpen }) => ({
    position: isMobile ? "initial" : "fixed",

    backgroundColor: 'white',
    width: isMobile ? '100vw' : isOpen ? 250 : 55,
    height: "100%",
    borderRight: `2px ${theme.palette.primary.whiteHigh} solid`,

    transition: isMobile ? 'none' : 'width 200ms ease-in-out',

    overflow: 'hidden',
    zIndex: 9,
}))
const NavigationBar = styled(({ children, ...props }: SideMenuProps) => (
    <Stack component="nav" {...props}>
        {children}
    </Stack>
))(({ isMobile }) => ({
    gap: isMobile ? "2rem" : "1rem",

    padding: isMobile ? "0.7rem" : "0.5rem",
    paddingBottom: "3rem",
    height: "100%",
    zIndex: "8",

    overflowY: 'scroll',
    scrollbarWidth: 'none',
    '&::WebkitScrollbar': {
        display: 'none',
    }
}))

const SideMenu = (
    { isMobile, isOpen }: SideMenuProps
): JSX.Element => {
    const user = useAppSelector((store) => store.user.user)

    const data = useMemo(() => {
        switch (user?.role) {
            case "student":
                return students_sidemenu_data
            case "admin":
                return admin_sidemenu_data
            case "supervisor":
                return supervisor_sidemenu_data
            case "teacher":
                return teacher_sidemenu_data
            default:
                return []
        }
    }, [user?.role])

    return (
        <Layout isOpen={isOpen} isMobile={isMobile}>
            <NavigationBar isMobile={isMobile}>
                {
                    data.map((section, i) => {
                        return (
                            <Fragment key={section.title + i}>
                                <MenuItems
                                    title={i == 0 ? "" : section.title}
                                    showIcon={true}
                                    isOpen={isOpen}
                                    items={!isMobile ? i == 0 ? [section.items[0]] : section.items : section.items}
                                />
                            </Fragment>
                        )
                    })
                }
            </NavigationBar>
        </Layout>
    )
}

export default SideMenu
