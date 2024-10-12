import {
    Box,
    Paper,
    Typography,
    styled,
} from '@mui/material'
import {
    useContext,
    MouseEventHandler,
} from 'react'
import {
    NotificationContext
} from '../../../context/notificationsContext'

import SidePopup from './SidePopup'

interface INavbarItemGroupProps {
    // for toggle menu open click event handling
    id: string,
    isOpen?: string,
    onClickCb: Function,

    icon: JSX.Element,
    children: JSX.Element,
    title: string, // for small title popup

    sx?: object,

    structure?: "side" | "popup"
}

const PopupIconLabelBox = styled(({ children, ...props }: { children: JSX.Element }) => (
    <Box {...props}>{children}</Box>
))(() => ({
    position: "absolute",
    top: 55,

    display: "flex",
    alignItems: 'center',
    justifyContent: "center",

    backgroundColor: 'rgba(50,50,50,0.8)',
    width: 'auto',
    padding: '2px 6px',
    borderRadius: '5px',
    pointerEvents: 'none',

    textWrap: 'nowrap',

    opacity: 0,
    transform: "scale(0.8)",
    transition: 'all 250ms ease-in-out',

    '&:before': {
        position: 'absolute',
        top: '-12px',

        width: '0px',
        height: '5px',
        borderLeft: '5px solid transparent',
        borderTop: '5px solid transparent',
        borderBottom: '7px solid rgba(50,50,50,0.8)',
        borderRight: '5px solid transparent',

        content: '""',
    },
}))

const PopupIconButton = styled((
    { popupIsOpen, hasNotifications, children, onClickCb, title, ...props }: { hasNotifications?: boolean, popupIsOpen?: string, title: string, children: JSX.Element, onClickCb: MouseEventHandler<HTMLDivElement> }
) => (
    <Box onClick={onClickCb || undefined} {...props}>
        <div style={{ position: 'relative' }}>
            {
                hasNotifications &&
                <Box sx={{
                    position: 'absolute',
                    right: '0px',

                    content: '""',
                    backgroundColor: (theme) => theme.palette.primary.contrast,
                    width: 7.6,
                    height: 7.6,
                    borderRadius: 20
                }}></Box>
            }
            {children}
        </div>
        {
            !popupIsOpen &&
            <PopupIconLabelBox>
                <Typography
                    color="primary.main"
                    variant="body1"
                >
                    {title}
                </Typography>
            </PopupIconLabelBox>
        }
    </Box>
))(() => ({
    position: 'relative',

    display: "flex",
    alignItems: "center",
    justifyContent: "center",

    width: 40,
    height: 40,
    borderRadius: 8,

    cursor: 'pointer',

    '&:hover > .MuiBox-root': {
        opacity: 1,
        top: 55,
        transform: "scale(1)",
    },
    '&:hover': {
        backgroundColor: `rgb(150, 205, 245,0.3)`,
    },
}))

const SidePopupWithIcon = (props: INavbarItemGroupProps) => {
    const { title, icon, isOpen, onClickCb } = props;
    const { hasNotifications } = useContext(NotificationContext);
    const id = props.id ? props.id : "_"

    return (
        <div>
            <PopupIconButton
                title={title}
                popupIsOpen={isOpen}
                hasNotifications={hasNotifications}
                onClickCb={() => (onClickCb(id) as MouseEventHandler<HTMLDivElement>)}
            >
                {icon}
            </PopupIconButton>
            {
                (isOpen === id) &&
                <SidePopup {...props} />
            }
        </div>
    )
}

const BubblePopupWithIcon = (props: INavbarItemGroupProps) => {
    const { title, icon, children, isOpen, onClickCb, id, sx } = props;

    return (
        <div style={{ position: 'relative' }}>
            <PopupIconButton
                title={title}
                popupIsOpen={isOpen}

                sx={sx ? { ...sx } : {}}
                onClickCb={() => (onClickCb(id) as MouseEventHandler<HTMLDivElement>)}
            >
                {icon}
            </PopupIconButton>
            {
                (isOpen === id) &&
                <Paper
                    elevation={3}
                    sx={{
                        position: "absolute",
                        top: 55,
                        right: 0,

                        backgroundColor: 'white',
                        height: 'auto',
                        maxWidth: 500,
                        padding: 1,
                        borderRadius: 5,
                    }}
                >
                    {children}
                </Paper>
            }
        </div>
    )
}
const NavbarItemPopup = (props: INavbarItemGroupProps) => {
    switch (props.structure) {
        case "side":
            return <SidePopupWithIcon {...props} />
        default:
            return <BubblePopupWithIcon {...props} />
    }
}

//     const { title, icon, href, showIcon, iconSize, fontSize, fontWeight } = props;

export default NavbarItemPopup;