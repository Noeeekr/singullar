import {
    MouseEventHandler,
} from 'react'

import {
    Box,
    Stack,
} from '@mui/material'
import Divider from './Divider'
import MenuItems from './SideMenuItems'
import SectionTitle from './SectionTitle'
import LinkButton from './LinkButton'
import ButtonGroup from './ButtonGroup'
import SidePopupButton from './SidePopupButton'
import NotificationPopup from './NotificationPopup'

import {
    menuItems_Classroom,
    menuItems_Main,
    menuItems_QuickAccess,
    sideMenuData_MyAccount,
} from '../data'

/*
* onHover must be a toggle type of effect to work.
*/
const SideMenu = (
    { isMobile, isOpen, onHoverOpen }: { isMobile?: boolean, isOpen?: boolean, onHoverOpen?: Function }
): JSX.Element => {
    return (
        <Box
            sx={{
                backgroundColor: 'white',
                borderRight: (theme) => `2px ${theme.palette.primary.whiteHigh} solid`,
                
                overflow: 'hidden',
                
                transition: isMobile ? 'none' : 'width 200ms ease-in-out',
            }}
            height="auto"
            width={isMobile ? '100vw' : isOpen ? 250 : 55}
            padding={isMobile ? 2 : 1}

            onMouseEnter={isOpen ? undefined : onHoverOpen as MouseEventHandler<HTMLDivElement> }
            onMouseLeave={onHoverOpen as (MouseEventHandler<HTMLDivElement> | undefined)}
        >
            <div style={{ overflow: 'hidden' }}>
                <Stack
                    component="nav"
                    gap={isMobile ? 4 : 2}
                >
                    { /* INICIO */}

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
                                            content={<NotificationPopup/>}
                                        />
                                        <SidePopupButton
                                            type="popup"
                                            title={menuItems_Main.items[2].title}
                                            showIcon={true}
                                            icon={menuItems_Main.items[2].icon}
                                            content={<NotificationPopup/>} // THE CONTENT HERE WILL GET A SET HAVE NOTIFICATION CB TO CHANGE HAS NOTIFICATION PARAM
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

                    { /* SALA DE AULA */}

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
                    {/* ACESSO RÁPIDO */}
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
                </Stack>
            </div>

        </Box>
    )
}

export default SideMenu