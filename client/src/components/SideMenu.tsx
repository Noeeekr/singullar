// Features
import { styled } from '@mui/material'
import { useAppSelector } from '../slices/store';
import {
    useMemo,
    Fragment,
} from 'react'

// Types
import type { BoxProps } from '@mui/material'

// Components
import { Stack } from '@mui/material'
import Divider from './Divider'
import MenuItems from './SideMenuButtons'
import SectionTitle from './SectionTitle'

// Data 
import { admin_sidemenu_data } from '../pages/admin/data'
import { supervisor_sidemenu_data } from '../pages/supervisor/data'
import { teacher_sidemenu_data } from '../pages/teacher/data'
import { students_sidemenu_data } from '../pages/student/data'

interface ISideMenuLayout extends BoxProps {
    isMobile?: boolean,
    isOpen?: boolean,
}

const SideMenuLayout = styled('nav')<ISideMenuLayout>(
    ({ theme, isMobile, isOpen }) => (
        {
            position: isMobile ? "initial" : "fixed",

            backgroundColor: 'white',
            width: isMobile ? '100vw' : isOpen ? 250 : 55,
            height: "100%",
            borderRight: `2px ${theme.palette.primary.whiteHigh} solid`,

            transition: isMobile ? 'none' : 'width 200ms ease-in-out',

            overflow: 'hidden',
            zIndex: 6,
        }
    )
)

/**
* onHover must be a toggle type of effect to work.
*/
const SideMenu = (
    { isMobile, isOpen }: { isMobile?: boolean, isOpen?: boolean, onHoverOpen?: () => void }
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

/*
    *Memoize???
    
    const roles = useAppSelector(() => user.roles)

    const TeacherOnlyMenus = lazy(...)
    const StudentOnlyMenus = lazy(...)

    // All the components that should be loaded
    const LinkButtonArray = []

    for role in roles {
        switch(role) {
        case Student:
            LinkButtonArray.append(a, b, c, d, e, f, g, h)
        case Teacher:
            LinkButtonArray.append(a, b, e, f, i, j, k)
        ...
        }
    }

    return LinkButtonArray
*/

    return (
        <SideMenuLayout isOpen={isOpen} isMobile={isMobile}>
                <Stack
                    component="nav"
                    gap={isMobile ? 4 : 2}
                    
                    padding={isMobile ? 1.2 : 1}
                    paddingBottom={3}
                    height="100%"
                    zIndex="8"

                    sx={{
                        overflowY: 'scroll',
                        scrollbarWidth: 'none',
                        '&::WebkitScrollbar': {
                            display: 'none',
                        }
                    }}
                >
                    { /* INICIO */}

                    {
                        data.map((section, i) => {
                            if (i === 0) {
                                return (
                                    <Stack gap={1} marginTop={1} key={section.title}>
                                        <MenuItems
                                            showIcon={true}
                                            isOpen={isOpen}
                                            items={isMobile ? section.items : [section.items[0]]}
                                        />
                                    </Stack>
                                )
                            }
                            if ((data.length - 1) === i) {
                                return (
                                    <Fragment key={section.title + i}>
                                        {
                                            !isOpen && <Divider />
                                        }
                                        <Stack gap={1}>
                                            {
                                                isOpen &&
                                                <SectionTitle>
                                                    {section.title}
                                                </SectionTitle>
                                            }
                                            <MenuItems
                                                showIcon={true}
                                                isOpen={isOpen}
                                                items={section.items}
                                            />
                                        </Stack>
                                    </Fragment>
                                )
                            }
                            return (
                                <Fragment key={section.title + i}>
                                    {
                                        !isOpen && <Divider />
                                    }
                                    <Stack gap={1}>
                                        {
                                            isOpen && <SectionTitle>
                                                {section.title}
                                            </SectionTitle>
                                        }
                                        <MenuItems
                                            showIcon={true}
                                            isOpen={isOpen}
                                            items={section.items}
                                        />
                                    </Stack>
                                </Fragment>
                            )
                        })
                    }
                </Stack>
        </SideMenuLayout>
    )
}

export default SideMenu
