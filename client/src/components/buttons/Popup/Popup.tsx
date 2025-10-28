import { useContext } from 'react'
import { useState } from 'react'
import { useMemo } from 'react'

import { NotificationContext } from '../../../context/notificationsContext'

import type { ButtonProps } from '../Default'
import type { AdminNavbarPopupIds } from '@components/layout/sidemenu/data/admin_routes';
import type { StudentNavbarPopupIds } from '@components/layout/sidemenu/data/student_routes';
import { TeacherNavbarPopupId } from '@components/layout/sidemenu/data/teacher_routes'
import { SupervisorNavbarPopupId } from '@components/layout/sidemenu/data/supervisor_routes'
import Popup, { PopupsProps } from '@components/popups'
import Icon from './Icon'
import { Stack, StackProps, Typography } from '@mui/material'
import styled from '@emotion/styled'

export interface PopupButtonProps extends ButtonProps, PopupsProps {
    id: AdminNavbarPopupIds | StudentNavbarPopupIds | TeacherNavbarPopupId | SupervisorNavbarPopupId
    type: "popup"
}

const PopupContainer = styled(({ children, ...props}: StackProps & { variant?: "large" | "small" }) => (
    <Stack direction="row" alignItems="center" {...props}>
        {children}
    </Stack>
))(({ theme, variant }) => ({
    width: variant == "large" ? "100%" : "initial",
    "&:hover": variant == "large" ? {
        backgroundColor: theme.palette.primary.purpleLightInv,
        borderRadius: 4,
    } : {},
    cursor: "pointer",
}))
const PopupButton = ({
    isMobile,
    id,
    children,
    fontWeight,
    fontSize,
    iconVariant,
    ...props
}: PopupButtonProps) => {
    const [isOpen, setIsOpen] = useState(false)

    const { notifications } = useContext(NotificationContext)
    const hasNotifications = useMemo(() => Object.keys(notifications).includes(id), [id])

    return (
        <PopupContainer variant={iconVariant}>
            <Icon iconVariant={iconVariant}
                hasNotifications={hasNotifications}
                onClick={() => setIsOpen(prevState => !prevState)}
                isOpen={isOpen}
                {...props}
            />
            {
                iconVariant == "large" 
                ? <Typography variant="body2" sx={{ paddingX: 1 }} fontSize={fontSize} fontWeight={fontWeight || 500}>
                    {props.title}
                </Typography>
                : <></>
            }
            <Popup
                title={props.title}
                variant={props.variant}
                onClose={() => setIsOpen(prev => !prev)}
                isOpen={isOpen}
            >
                {children}
            </Popup>
        </PopupContainer>
    )
}

export default PopupButton;