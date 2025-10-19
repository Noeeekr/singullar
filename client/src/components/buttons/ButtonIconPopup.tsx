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
} from '../../context/notificationsContext'

import SidePopup from '../popups/SidePopup'

interface NavigationPopupsProps {
    // for toggle menu open click event handling
    id: string,
    isOpen?: string,
    isCorner?: boolean,
    onClickCb: (key: string) => void,

    icon?: JSX.Element,
    children: JSX.Element,
    title: string, // for small title popup

    sx?: object,

    variant?: "side" | "bubble"
}

const PopupIconLabelBox = styled(({ children, ...props }: { isCorner?: boolean, children: JSX.Element }) => (
    <Box {...props}>{children}</Box>
))(({ isCorner }) => ({
    position: "absolute",
    top: 55,
    right: isCorner ? "0" : "initial",

    display: "flex",
    alignItems: 'center',
    justifyContent: "center",

    backgroundColor: 'rgba(80,80,80,1)',
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
        right: isCorner ? "20%" : "initial",

        width: '0px',
        height: '5px',
        borderLeft: '5px solid transparent',
        borderTop: '5px solid transparent',
        borderBottom: '10px solid rgba(80,80,80,1)',
        borderRight: '5px solid transparent',

        content: '""',
    },
}))

const PopupIconButton = styled((
    { popupIsOpen, isCorner, hasNotifications, children, onClickCb, title, ...props }: { isCorner?: boolean, hasNotifications?: boolean, popupIsOpen?: string, title: string, children: JSX.Element, onClickCb: MouseEventHandler<HTMLDivElement> }
) => (
    <Box onClick={onClickCb || undefined} {...props}>
        <div style={{ position: 'relative', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
            {
                hasNotifications &&
                <Box sx={{
                    position: 'absolute',
                    right: '0px',

                    content: '""',
                    backgroundColor: (theme) => theme.palette.primary.contrast,
                    width: 7.6,
                    height: 7.6,
                    borderRadius: 20,
                    alignSelf: 'flex-start',
                }}/>
            }
            {children}
        </div>
        {
            !popupIsOpen &&
            <PopupIconLabelBox isCorner={isCorner}>
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

const SidePopupWithIcon = (props: NavigationPopupsProps) => {
    const { title, icon, isOpen, isCorner, onClickCb } = props;
    const { notifications } = useContext(NotificationContext);

    const id = props.id ? props.id : "_"

    const hasNotifications = Object.keys(notifications).includes(id);

    return (
        <div>
            <PopupIconButton
                isCorner={isCorner}
                title={title}
                popupIsOpen={isOpen}
                hasNotifications={hasNotifications}
                onClickCb={() => (onClickCb(id))}
            >
                {icon ? icon : <></>}
            </PopupIconButton>
            {
                (isOpen === id) &&
                <SidePopup {...props} isOpen={true}/>
            }
        </div>
    )
}

const BubblePopupWithIcon = (props: NavigationPopupsProps) => {
    const { title, isCorner, icon, children, isOpen, onClickCb, id, sx } = props;

    return (
        <div style={{ position: 'relative' }}>
            <PopupIconButton
                title={title}
                popupIsOpen={isOpen}
                isCorner={isCorner}

                sx={sx ? { ...sx } : {}}
                onClickCb={() => (onClickCb(id))}
            >
                {icon ? icon : <></>}
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
const NavigationPopup = (props: NavigationPopupsProps) => {
    switch (props.variant) {
        case "side":
            return <SidePopupWithIcon {...props} />
        case "bubble":
            return <BubblePopupWithIcon {...props} />
        default:
            return <BubblePopupWithIcon {...props} />
    }
}

export default NavigationPopup;