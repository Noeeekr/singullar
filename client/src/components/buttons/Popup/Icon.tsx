import {
    Box,
    Stack,
    StackProps,
    Typography,
    styled,
    useTheme,
} from '@mui/material'

import type { PopupIconLabelProps, PopupIconProps } from "../../popups"
import { cloneElement } from 'react'

const Label = styled(({ children, title, ...props }: PopupIconLabelProps & StackProps) => (
    <Box {...props}>
        <Typography
            color="primary.main"
            variant="body1"
        >
            {title}
        </Typography>
    </Box>
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

const Icon = styled(({ children, ...props }: StackProps) => (
    <Stack {...props}>
        {children}
    </Stack>
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
}))

export default function ({
    children,
    isOpen,
    iconVariant,
    hasNotifications,
    title,
    isCorner,
    icon: { display = true, component, size, style } = { display: true, component: <></>, size: 14, style: {} },
    ...props
}: PopupIconProps): JSX.Element {
    const theme = useTheme()

    return (
        <Icon sx={{
            '&:hover': iconVariant =="large" ? {} : {
                backgroundColor: `rgb(150, 205, 245,0.3)`,
            },
        }} {...props}>
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
                    }} />
                }
                {
                    display ?
                        <Box
                            sx={{
                                display: "flex",
                                alignItems: "center",
                                justifyContent: "center",

                                opacity: 0.7,
                                padding: "0.4rem",
                            }}
                        >
                            {
                                cloneElement(component, {
                                    style: style || { 
                                        fontSize: size || "1.5rem",
                                        color: iconVariant == "large" ? theme.palette.primary.purpleDark : "white",
                                    },
                                })
                            }
                        </Box>
                        : <></>
                }
            </div>
            {
                !isOpen && title
                    ? <Label title={title} isCorner={isCorner} />
                    : <></>
            }
        </Icon>
    )
}

// const SidePopupWithIcon = (props: PopupProps) => {
//     const { notifications } = useContext(NotificationContext);

//     const id = props.id ? props.id : "_"

//     const hasNotifications = Object.keys(notifications).includes(id);

//     return (
//         <div>
//             <PopupIconButton
//                 id={props.id}
//                 isCorner={Boolean(props.isCorner)}
//                 title={props.title}
//                 isOpen={Boolean(props.isOpen)}
//                 hasNotification={hasNotifications}
//                 onClick={props.onClick}
//             >
//                 {props.icon ? props.icon : <></>}
//             </PopupIconButton>
//             {
//                 <SidePopup {...props} isOpen={props.isOpen} />
//             }
//         </div>
//     )
// }
