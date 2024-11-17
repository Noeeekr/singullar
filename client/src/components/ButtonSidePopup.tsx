import {
    cloneElement,
    useCallback,
    useContext,
    useState,
    useMemo,
} from 'react'
import {
    NotificationContext
} from '../context/notificationsContext'

import {
    Box,
    Typography,
    Stack,
    useMediaQuery
} from '@mui/material'
import { useTheme } from '@mui/material/styles'

import { MenuItemStack } from './SideMenuButtons'

import { ISideMenuPopupButton } from '../types/buttonProps';
import SidePopup from './SidePopup'

interface IPopupButtonProps extends ISideMenuPopupButton {
    hasHoverEffect?: boolean,
    showIcon?: boolean,
    iconSize?: number,
    fontSize?: number,
    fontWeight?: number,
}

const PopupButton = (props: IPopupButtonProps) => {
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))

    const [isOpen, setIsOpen] = useState(false)
    const toggleIsOpen = useCallback(
        () => (setIsOpen(prevState => !prevState))
    ,[])

    const {
        title,
        content,

        id,
        icon,
        showIcon,

        iconSize,
        fontSize,
        fontWeight,
        hasHoverEffect,
    } = props;

    const Icon = useMemo(() => {
        return Boolean(showIcon)
            ? <Box display='flex' sx={{ opacity: 0.7 }}>
                {
                    cloneElement(icon,
                        {
                            fontSize: iconSize || 21,
                            color: theme.palette.primary.purpleDark,
                        }
                    )
                }
            </Box>
            : <></>
    }, [showIcon]);

    const { notifications } = useContext(NotificationContext)
    const hasNotifications = useMemo(
        () => (Object.keys(notifications).includes(id))
    ,[id])

    return (
        <>
            <MenuItemStack
                component="li"

                paddingX={1}
                paddingY={isMobile ? 1.5 : 1}

                sx={{
                    "&:hover": {
                        backgroundColor: hasHoverEffect ? theme.palette.primary.purpleLightInv : "none",
                    },
                }}

                onClick={toggleIsOpen}
            >
                <Box sx={{ position: 'relative' }}>
                    {Icon}
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
            </MenuItemStack>
            <SidePopup onClickCb={toggleIsOpen} title={title} isOpen={isOpen}>
                {content}
            </SidePopup>
        </>
    )
}

export default PopupButton;