import {
    useState,
    useEffect,
    useCallback,
    cloneElement,
} from 'react'

import {
    Box,
    Stack,
    Typography,
    useMediaQuery,
} from '@mui/material'
import {
    useTheme,
} from '@mui/material/styles'

import {
    FaCaretUp,
} from "react-icons/fa";

import {
    ISideMenuLinkButton,
    ISideMenuButton
} from '../models/buttonProps'

import MenuItems, { MenuItemStack } from './SideMenuButtons'

interface ILinkButtonGroupProps {
    title: string,
    icon: JSX.Element,
    items: (ISideMenuLinkButton | ISideMenuButton)[],

    fontSize?: number,
    isOpen?: boolean,
}

const MenuItemGroup = (props: ILinkButtonGroupProps): JSX.Element => {
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))

    const { icon, title, items, fontSize, isOpen } = props;

    const Icon = cloneElement(
        icon,
        {
            fontSize: fontSize || 21,
            color: theme.palette.primary.purpleDark,
        }
    )

    const [isWardrobeOpen, setIsWardrobeOpen] = useState(false);
    const toggleDrawer = useCallback(() => { 
        if (isOpen) setIsWardrobeOpen(prevState => !prevState);
    },[isOpen]);

    useEffect(() => {
        if (!isOpen && isWardrobeOpen) {
            setIsWardrobeOpen(false)
        }
    }, [isOpen])

    return (
        <Stack component="ul" gap={0.5}>
            <MenuItemStack
                paddingX={1}
                paddingY={isMobile ? 1.5 : 1}

                onClick={toggleDrawer}
            >
                {/* ICON, TEXT CONTENT, EXPAND MENU ARROW */}
                <Box display='flex' sx={{ opacity: 0.7 }}>
                    {Icon}
                </Box>
                <Typography variant="body2" fontWeight={500} sx={{ paddingX: 1 }}>
                    {title}
                </Typography>
                <FaCaretUp
                    fontSize={10}
                    style={{
                        height: '100%',
                        margin: 'auto 5px auto auto',
                        transform: isWardrobeOpen ? 'rotate(180deg)' : 'rotate(0deg)',
                        transition: 'transform 200ms ease-in-out'
                    }}
                />
            </MenuItemStack>
            {
                isWardrobeOpen &&
                    <Box sx={{ paddingLeft: 2.6 }}>
                        <MenuItems items={items} fontWeight={350} fontSize={13.5} />
                    </Box>
            }
        </Stack>
    )
}

export default MenuItemGroup