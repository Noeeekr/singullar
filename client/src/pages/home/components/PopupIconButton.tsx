import {
    Box,
    Paper,
} from '@mui/material'
import {
    MouseEventHandler,
} from 'react'

import SidePopup from './SidePopup'

interface INavbarItemGroupProps {
    // for toggle menu open click event handling
    id: string,
    isOpen?: string,
    onClickCb: Function,

    icon: JSX.Element,
    children: JSX.Element,
    title: string, // for small title popup

    structure?: "side" | "popup"
}

const SidePopupWithIcon = (props: INavbarItemGroupProps) => {
    const { icon, isOpen, onClickCb } = props;
    const id = props.id ? props.id : "_"
    

    return(
        <Box>
            <Box
                onClick={() => (onClickCb(id) as MouseEventHandler<HTMLDivElement>)}
                sx={{
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',

                    width: 40,
                    height: 40,
                    borderRadius: 2,

                    '&:hover': {
                        backgroundColor: `rgb(140, 195, 255,0.4)`,
                    },

                    cursor: 'pointer',
                }}
            >
                {icon}
            </Box>
            {
                (isOpen === id) &&
                <SidePopup {...props}/>
            }
        </Box>
    )
}

const BubblePopupWithIcon = (props: INavbarItemGroupProps) => {
    const { icon, children, isOpen, onClickCb, id } = props;

    return (
        <Box sx={{
            position: 'relative'
        }}>
            <Box
                onClick={() => (onClickCb(id))}
                display="flex"
                alignItems="center"
                justifyContent="center"
                sx={{
                    width: 40,
                    height: 40,
                    borderRadius: 2,

                    '&:hover': {
                        backgroundColor: `rgb(140, 195, 255,0.4)`,
                    },

                    cursor: 'pointer',
                }}
            >
                {icon}
            </Box>
            {
                (isOpen === id) &&
                <Paper
                    elevation={1}
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
        </Box>
    )
}
const NavbarItemPopup = (props: INavbarItemGroupProps) => {
    switch(props.structure) {
        case "side": 
            return <SidePopupWithIcon {...props}/>
        default: 
            return <BubblePopupWithIcon {...props}/>
    }
}

//     const { title, icon, href, showIcon, iconSize, fontSize, fontWeight } = props;

export default NavbarItemPopup;