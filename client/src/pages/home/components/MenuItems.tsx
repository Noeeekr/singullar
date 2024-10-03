import {
    useState,
    cloneElement
} from 'react'

import {
    Link,
} from 'react-router-dom'

import {
    Box,
    Stack,
    Typography,
    useMediaQuery,
    styled,
    StackProps,
} from '@mui/material'
import {
    useTheme,
} from '@mui/material/styles'

import {
    ISideMenuLink,
    ISideMenuGroup,
} from '../../../types/sideMenu'

import {
    FaCaretUp,
} from "react-icons/fa";

interface IMenuItemsProps {
    items: (ISideMenuLink | ISideMenuGroup)[]
    
    showIcon?: boolean
    isCompacted?: boolean
    fontSize?: number
    fontWeight?: number
}
interface ISideMenuItemGroupProps {
    title: string,
    type: "group",
    icon: JSX.Element,
    items: ISideMenuLink[],

    fontSize?: number,
}
interface ISideMenuItemLinkProps {
    title: string,
    type: "link",
    icon: JSX.Element,
    href: string,

    showIcon?: boolean,
    fontSize?: number,
    iconSize?: number,
    fontWeight?: number,
}

const MenuItemStack = styled(({ children, onClick, ...props }: StackProps) => (
    <Stack direction="row" onClick={onClick} {...props}>{ children }</Stack>
))(({ theme }) => ({
    alignItems: "center",
    overflow: 'hidden',
    gap: 10,
    textWrap: 'nowrap',
    cursor: "pointer",
    color: "black",
    "&:hover > .MuiBox-root": {
        opacity: 1,
        transition: 'opacity 200ms ease-in-out'
    },
    "&:active": {
        borderRadius: 2,
        backgroundColor: theme.palette.primary.purpleLightInv
    },
}))

const MenuItems = ({ items, ...props}: IMenuItemsProps): JSX.Element => {

    return (
        <Stack component="ul" gap={0.5}>
            {
                items.map((item) => {
                    if (item.type === "link")
                        return <MenuItemLink key={item.title} {...props} {...item} />;
                    return <MenuItemGroup key={item.title} {...props} {...item} />
                })
            }
        </Stack>
    )
}

const MenuItemGroup = (props: ISideMenuItemGroupProps): JSX.Element => {
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))

    const { icon, title, items, fontSize } = props;
    
    const Icon = cloneElement(
        icon,
        {
            style: {
                fontSize: fontSize || 21,
            },
            color: theme.palette.primary.purpleDark,
        }
    )

    const [isOpen, setIsOpen] = useState(false);
    const toggleDrawer = () => (setIsOpen(prevState => !prevState));

    return (
        <Stack
            component="ul"
            gap={0.5}
        >
            <MenuItemStack
                paddingX={1}
                paddingY={isMobile ? 1.5 : 1}
                gap={1}

                onClick={toggleDrawer}
            >
                {/* ICON, TEXT CONTENT, EXPAND MENU ARROW */}
                <Box display='flex' sx={{ opacity: 0.7 }}>
                    {Icon}
                </Box>
                <Typography variant="body1" fontWeight={500}>
                    {title}
                </Typography>
                <FaCaretUp
                    style={{
                        height: '100%',
                        margin: 'auto 5px auto auto',
                        transform: isOpen ? 'rotate(180deg)' : 'rotate(0deg)',
                        transition: 'transform 200ms ease-in-out'
                    }}
                />
            </MenuItemStack>
            {
                isOpen
                    ? <Box sx={{ paddingLeft: 3.8 }}>
                        <MenuItems items={items} fontWeight={350} fontSize={13.5} />
                    </Box>
                    : <></>
            }
        </Stack>
    )
}

const MenuItemLink = (props: ISideMenuItemLinkProps): JSX.Element => {
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('xs'))

    const { title, icon, href, showIcon, iconSize, fontSize, fontWeight } = props;

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
        <Link to={href} style={{ textDecoration: 'none'}}>
            <MenuItemStack
                component="li"

                paddingX={1}
                paddingY={isMobile ? 1.5 : 1}
            >
                {Icon}
                <Typography variant="body1" fontSize={fontSize} fontWeight={fontWeight || 500}>
                    {title}
                </Typography>
            </MenuItemStack>
        </Link>
    )
}

export default MenuItems
