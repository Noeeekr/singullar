// features

import { useAppSelector } from '../slices/store'; 
import {
    useMemo,
    Fragment,
    MouseEventHandler,
} from 'react'

import type {
    SxProps,
    Theme,
} from '@mui/material'
// components

import {
    Box,
    Stack,
} from '@mui/material'
import Divider from './Divider'
import MenuItems from './SideMenuItems'
import SectionTitle from './SectionTitle'


// data 

import { admin_sidemenu_data } from '../pages/admin/data'
import { supervisor_sidemenu_data } from '../pages/supervisor/data'
import { teacher_sidemenu_data } from '../pages/teacher/data'
import { students_sidemenu_data } from '../pages/home/data'

/*
* onHover must be a toggle type of effect to work.
*/
const SideMenu = (
    { isMobile, isOpen, onHoverOpen, sx }: { sx?: SxProps<Theme>, isMobile?: boolean, isOpen?: boolean, onHoverOpen?: Function }
): JSX.Element => {

    const user = useAppSelector((store) => store.auth.user)

    const data = useMemo(() => {
        switch(user?.role) {
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
    },[user?.role])

    return (
        <Box
            sx={{
                backgroundColor: 'white',
                borderRight: (theme) => `2px ${theme.palette.primary.whiteHigh} solid`,

                overflow: 'hidden',

                transition: isMobile ? 'none' : 'width 200ms ease-in-out',
                ...sx
            }}
            height="auto"
            width={isMobile ? '100vw' : isOpen ? 250 : 55}
            padding={isMobile ? 2 : 1}

            onMouseEnter={isOpen ? undefined : onHoverOpen as MouseEventHandler<HTMLDivElement>}
            onMouseLeave={onHoverOpen as (MouseEventHandler<HTMLDivElement> | undefined)}
        >
            <div style={{ overflow: 'hidden' }}>
                <Stack
                    component="nav"
                    gap={isMobile ? 4 : 2}
                >
                    { /* INICIO */}

                    {
                        data.map((section, i) => {
                            if (i === 0) {
                                return (
                                    <Stack gap={1} marginTop={1} key={section.title}>
                                        <MenuItems
                                            showIcon={true}
                                            isCompacted={!isOpen}
                                            items={isMobile ? section.items : [section.items[0]]}
                                        />
                                    </Stack>
                                )
                            }
                            if ((data.length - 1) === i) {
                                return (
                                    <Fragment key={section.title + i}>
                                        {
                                            !isOpen && <Divider/>
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
                                                isCompacted={!isOpen}
                                                items={section.items}
                                            />
                                        </Stack>
                                    </Fragment>
                                )
                            }
                            return (
                                <Fragment key={section.title + i}>
                                    {
                                        isOpen && !isMobile
                                            ? <></>
                                            : <Divider/>
                                    }
                                    <Stack gap={1}>
                                        {
                                            (isOpen || isMobile) && <SectionTitle>
                                                {section.title}
                                            </SectionTitle>
                                        }
                                        <MenuItems
                                            showIcon={true}
                                            isCompacted={!isOpen}
                                            items={section.items}
                                        />
                                    </Stack>
                                </Fragment>
                            )
                        })
                    }
                </Stack>
            </div>

        </Box>
    )
}

export default SideMenu

/**
<Box marginTop={isMobile ? 0 : 2}>
                        {
                            <LinkButton
                                showIcon={true}
                                type="link"
                                href="/home"
                                title={menuItems_Main.items[0].title}
                                icon={menuItems_Main.items[0].icon} />
                        }
                        {
                            isMobile
                                ? (
                                    <>
                                        <SidePopupButton
                                            type="popup"
                                            title={menuItems_Main.items[1].title}
                                            showIcon={true}
                                            icon={menuItems_Main.items[1].icon}
                                            content={<NotificationPopup />}
                                        />
                                        <SidePopupButton
                                            type="popup"
                                            title={menuItems_Main.items[2].title}
                                            showIcon={true}
                                            icon={menuItems_Main.items[2].icon}
                                            content={<NotificationPopup />} // THE CONTENT HERE WILL GET A SET HAVE NOTIFICATION CB TO CHANGE HAS NOTIFICATION PARAM
                                        />
                                        <ButtonGroup
                                            title={menuItems_Main.items[3].title}
                                            icon={menuItems_Main.items[3].icon}
                                            items={sideMenuData_MyAccount}
                                        />
                                    </>
                                )
                                : <></>
                        }
                    </Box>
                    {
                        isOpen && !isMobile
                            ? <></>
                            : <Divider />
                    }

                    { /* SALA DE AULA }

                    <Stack gap={1}>
                        {
                            (isOpen || isMobile) && <SectionTitle>
                                {menuItems_Classroom.title}
                            </SectionTitle>
                        }
                        <MenuItems
                            showIcon={true}
                            isCompacted={!isOpen}
                            items={menuItems_Classroom.items}
                        />
                    </Stack>
                    {/* ACESSO RÁPIDO }
                    {
                        isMobile && (
                            <>
                                <Divider />
                                <Stack gap={1}>
                                    <SectionTitle>
                                        {menuItems_QuickAccess.title}
                                    </SectionTitle>
                                    <MenuItems
                                        showIcon={true}
                                        isCompacted={!isOpen}
                                        items={menuItems_QuickAccess.items}
                                    />
                                </Stack>
                            </>
                        )
                    }
 */