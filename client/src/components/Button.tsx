// Features
import { cloneElement } from 'react'
import { useMediaQuery } from '@mui/material'
import { useLocation } from 'react-router-dom'
import { useTheme } from '@mui/material/styles'

// Components
import {
    Box,
    Paper,
    Stack,
    styled,
    Typography,
} from '@mui/material'
import { MenuItemStack } from './SideMenuButtons'

// Types
import type { ISideMenuButton } from '../types/buttonProps';

export interface IButtonBaseProps {
    href?: string,
    fontSize?: number,
    iconSize?: number,
    showIcon?: boolean,
    fontWeight?: number,
    description?: string,
    hasHoverEffect?: boolean,
    showDescription?: boolean,
    hasNotifications?: boolean,
    variant?: "button" | "paper",
}

export type ISideMenuButtonProps = ISideMenuButton & IButtonBaseProps;

const PaperWrapper = styled(Paper)(() => ({
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'space-between',
    alignItems: 'flex-start',

    height: '120px',
    borderRadius: '17px',
    boxShadow: 'rgba(114, 119, 128, 0.09) 0px 1px 0px 0px,rgba(114, 119, 128, 0.09) 0px 2px 4px 0px, rgba(114, 119, 128, 0.09) 0px 4px 8px 0px',
    padding: '20px',
    border: 'solid 1px rgb(230,230,230)',

    transition: 'boxShadow 0ms linear',
    '&:hover': {
        boxShadow: 'rgba(114, 119, 128, 0.09) 0px 1px 0px 0px, rgba(114, 119, 128, 0.09) 0px 2px 4px 0px, rgba(114, 119, 128, 0.09) 0px 4px 8px 0px, rgba(114, 119, 128, 0.09) 0px 8px 16px 0px, rgba(114, 119, 128, 0.09) 0px 12px 24px 0px',

        cursor: 'pointer',
    }
}));

const DefaultButton = (props: ISideMenuButtonProps): JSX.Element => {
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

    const Icon = showIcon
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

const PaperButton = (props: ISideMenuButtonProps): JSX.Element => {
    const {
        fontWeight,

        showDescription,

        icon,
        iconSize,
        showIcon,

        title,
        description,
    } = props;

    const theme = useTheme();

    const Icon = showIcon
        ? <Box display='flex' sx={{ opacity: 0.7 }}>
            {
                cloneElement(icon,
                    {
                        style: {
                            fontSize: iconSize || 25,
                        },
                        color: theme.palette.primary.purpleDark,
                    }
                )
            }
        </Box>
        : <></>;


    return (
        <PaperWrapper>
            {Icon}
            <Stack gap={0.5}>
                <Typography variant="body2" fontWeight={fontWeight || 500}>
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
        </PaperWrapper>
    )
}

const Button = (props: ISideMenuButtonProps) => {
    switch(props.variant) {
        case "button":
            return <DefaultButton {...props}/>
        case "paper":
            return <PaperButton {...props}/>
        default: 
            return <DefaultButton {...props}/>
    }
}

export default Button;