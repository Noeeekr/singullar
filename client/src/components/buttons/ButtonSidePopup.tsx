import { cloneElement } from 'react'
import { useCallback } from 'react'
import { useContext } from 'react'
import { useState } from 'react'
import { useMemo } from 'react'

import Box from "@mui/material/Box"
import Typography from "@mui/material/Typography"
import Stack from "@mui/material/Stack"

import SidePopup from '../popups/SidePopup'
import { DefaultButtonLayout } from '@components/buttons/Button/Default'

import { useTheme } from '@mui/material/styles'
import { NotificationContext } from '../../context/notificationsContext'

import type { ButtonProps } from './Button/Button'
import type { AdminNavbarPopupIds } from '@components/layout/sidemenu/data/admin_routes';
import type { StudentNavbarPopupIds } from '@components/layout/sidemenu/data/student_routes';
import { TeacherNavbarPopupId } from '@components/layout/sidemenu/data/teacher_routes'
import { SupervisorNavbarPopupId } from '@components/layout/sidemenu/data/supervisor_routes'

export interface PopupButtonProps extends ButtonProps { 
    element: JSX.Element
    id: AdminNavbarPopupIds | StudentNavbarPopupIds | TeacherNavbarPopupId | SupervisorNavbarPopupId
    variant: "side" | "bubble"
    type: "popup"
}

const PopupButton = ({
    icon: { component: IconComponent, size: iconSize, display: displayIcon } = { component: <></> },
    isMobile,
    id,
    fontWeight,
    fontSize,
    title,
    element,
    ...props
}: PopupButtonProps) => {
    const theme = useTheme();

    const [isOpen, setIsOpen] = useState(false)
    const toggleIsOpen = useCallback(() => (setIsOpen(prevState => !prevState)), [])

    const { notifications } = useContext(NotificationContext)
    const hasNotifications = useMemo(() => Object.keys(notifications).includes(id), [id])

    return (
        <>
            <DefaultButtonLayout
                component="li"
                onClick={toggleIsOpen}
                {...props}
            >
                <Box sx={{ position: 'relative' }}>
                    {
                        displayIcon &&
                        <Box display='flex' sx={{ opacity: 0.7 }}>
                            {
                                cloneElement(IconComponent,
                                    {
                                        fontSize: iconSize || 21,
                                        color: theme.palette.primary.purpleDark,
                                    }
                                )
                            }
                        </Box>
                    }
                    {
                        hasNotifications &&
                        <Box sx={{
                            position: 'absolute',
                            top: '0',
                            right: '0',

                            content: '""',
                            backgroundColor: (theme) => theme.palette.primary.contrast,
                            width: 7.6,
                            height: 7.6,
                            borderRadius: 20,
                        }} />
                    }
                </Box>
                <Stack gap={0.5}>
                    <Typography variant="body2" sx={{ paddingX: 1 }} fontSize={fontSize} fontWeight={fontWeight || 500}>
                        {title}
                    </Typography>
                </Stack>
            </DefaultButtonLayout>
            <SidePopup onClickCb={toggleIsOpen} title={title} isOpen={isOpen}>
                { element }
            </SidePopup>
        </>
    )
}

export default PopupButton;