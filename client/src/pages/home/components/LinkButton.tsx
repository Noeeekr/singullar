import {
    cloneElement
} from 'react'

import {
    Link,
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
import { ISideMenuLinkButton } from '../../../types/sideMenu'

export interface ILinkButtonBaseProps {
    showDescription?: boolean,
    showIcon?: boolean,
    fontSize?: number,
    iconSize?: number,
    fontWeight?: number,
    hasHoverEffect?: boolean,
}

export interface ISideMenuLinkProps extends ISideMenuLinkButton, ILinkButtonBaseProps {}

const MenuItemLink = (props: ISideMenuLinkProps): JSX.Element => {
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))
    const pathname = useLocation().pathname

    const { title, icon, hasHoverEffect, href, showIcon, description, showDescription, iconSize, fontSize, fontWeight } = props;

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
        <Link to={href} style={{
            textDecoration: 'none',
        }}>
            <MenuItemStack
                component="li"

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
                    <Typography variant="body2" fontSize={fontSize} fontWeight={fontWeight || 500}>
                        {title}
                    </Typography>
                    {
                        showDescription && description
                            ?(
                                <Typography variant="body1" sx={{ textWrap: 'wrap', minWidth: 150}}>
                                    { description }
                                </Typography>
                            )
                            : <></>
                    }
                </Stack>
            </MenuItemStack>
        </Link>
    )
}

export default MenuItemLink