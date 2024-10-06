import {
    MouseEventHandler,
} from 'react'

import {
    Box,
    Stack,
} from '@mui/material'

import Divider from './Divider'
import MenuItems from './SideMenuItems'
import { 
    menuItems_Classroom,
    menuItems_Main,
    menuItems_QuickAccess
} from '../data'
import SectionTitle from './SectionTitle'

/*
* onHover must be a toggle type of effect to work.
*/
const SideMenu = (
    { isMobile, isOpen, onHoverOpen }: { isMobile?: boolean, isOpen?: boolean, onHoverOpen?: Function }
): JSX.Element => {

    return (
        <Box
            component="div"
            sx={{
                overflow: 'hidden',
                transition: isMobile ? 'none' : 'width 200ms ease-in-out',
                backgroundColor: 'white',
                borderRight: (theme) => `2px ${theme.palette.primary.whiteHigh} solid`
            }}
            height="auto"
            width={
                isMobile
                    ? '100vw'
                    : isOpen
                        ? 250
                        : 55
            }
            padding={isMobile ? 2 : 1 }
            onMouseEnter={isOpen ? undefined : onHoverOpen as MouseEventHandler<HTMLDivElement>}
            onMouseLeave={onHoverOpen as (MouseEventHandler<HTMLDivElement> | undefined)}            
        >
            <div style={{ overflow: 'hidden' }}>
                <Stack
                    component="nav"

                    gap={2.5}
                >
                    { /* INICIO */}

                    <Box marginTop={isMobile ? 0 : 2}>
                        {
                            <MenuItems
                                showIcon={true}
                                isCompacted={!isOpen}
                                items={
                                    isMobile
                                        ? menuItems_Main.items
                                        : [menuItems_Main.items[0]]
                                }
                            />
                        }
                    </Box>
                    {
                        isOpen && !isMobile 
                        ? <></>
                        : <Divider/>
                    }

                    { /* SALA DE AULA */}

                    <Stack gap={1}>
                        {
                            isOpen && <SectionTitle>
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
                                <Divider/>
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