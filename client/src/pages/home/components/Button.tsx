import {
    cloneElement,
} from 'react'

import {
    useLocation
} from 'react-router-dom'

import {
    Box,
    Stack,
    Typography,
    useMediaQuery,
} from '@mui/material'
import {
    useTheme,
} from '@mui/material/styles'

import { MenuItemStack } from './SideMenuItems'
import { ISideMenuButton } from '../../../types/propsButtons'

export interface IButtonBaseProps {
    href?: string,
    showDescription?: boolean,
    showIcon?: boolean,
    fontSize?: number,
    iconSize?: number,
    fontWeight?: number,
    description?: string,
    hasHoverEffect?: boolean,
    hasNotifications?: boolean,
}

export type ISideMenuButtonProps = ISideMenuButton & IButtonBaseProps;

const Button = (props: ISideMenuButtonProps): JSX.Element => {
    const {
        fontWeight,

        hasHoverEffect,
        showDescription,

        href,
        icon,
        iconSize,
        showIcon,

        title,
        description,
    } = props;

    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))
    const pathname = useLocation().pathname

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
        <MenuItemStack
            paddingX={1}
            paddingY={isMobile ? 1.5 : 1}

            sx={{
                "&:hover": hasHoverEffect ? {
                    backgroundColor: theme.palette.primary.purpleLightInv,
                } : {},
                backgroundColor: pathname == href ? theme.palette.primary.purpleLightInv : 'none',
            }}
        >
            {Icon}
            <Stack gap={0.5}>
                <Typography variant="body2" fontWeight={fontWeight || 500} sx={{ paddingX: 1 }}>
                    {title}
                </Typography>
                {
                    showDescription && description
                        ? (
                            <Typography variant="body1" sx={{ textWrap: 'wrap', minWidth: 150, paddingX: 1 }}>
                                {description}
                            </Typography>
                        )
                        : <></>
                }
            </Stack>
        </MenuItemStack>
    )
}

export default Button