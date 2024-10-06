import {
    cloneElement,
    useState,
} from 'react'
import {
    Box,
    Typography,
    Stack, 
    useMediaQuery
} from '@mui/material'
import { useTheme } from '@mui/material/styles'

import { MenuItemStack } from './SideMenuItems'

import { ISideMenuPopupButton } from '../../../types/sideMenu';
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
    const toggleIsOpen = () => (setIsOpen(prevState => !prevState))

    const { 
        title,
        content,

        icon,
        showIcon,
    
        iconSize,
        fontSize,
        fontWeight,
        hasHoverEffect,
    } = props;

    const Icon = Boolean(showIcon)
        ? <Box display='flex' sx={{ opacity: 0.7 }}>
            {
                cloneElement(icon,
                    {
                        style: {
                            fontSize: iconSize || 21,
                        },
                        color: theme.palette.primary.purpleDark,
                    }
                )
            }
        </Box>
        : <></>;

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
                {Icon}
                <Stack gap={0.5}>
                    <Typography variant="body2" fontSize={fontSize} fontWeight={fontWeight || 500}>
                        {title}
                    </Typography>
                </Stack>
            </MenuItemStack>
            {
                isOpen && 
                <SidePopup onClickCb={toggleIsOpen} title={title}>
                    { content }
                </SidePopup>
            }
        </>
    )
}

export default PopupButton;