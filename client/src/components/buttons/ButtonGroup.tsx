
// Components
import Box from "@mui/material/Box"
import Stack from "@mui/material/Stack"
import Typography from "@mui/material/Typography"
import { DefaultButtonLayout } from '@components/buttons/Button/Default'

// Icons
import { FaCaretUp } from "react-icons/fa";

// Features
import { useTheme } from '@mui/material/styles'
import { useState } from "react"
import { useEffect } from "react"
import { useCallback } from "react"
import { cloneElement } from "react"

// Types
import type { PopupButtonProps } from "./ButtonSidePopup"
import type { ButtonProps, ButtonsProps } from './Button/Button';
import type { LinkButtonProps } from './ButtonLink';
import SideMenuButtons from "@components/layout/sidemenu/SideMenuButtons"

export interface ButtonGroupProps extends ButtonProps {
    menus: (LinkButtonProps | ButtonsProps | ButtonGroupProps | PopupButtonProps)[]
    type: "group"

    isOpen?: boolean
}

const ButtonGroup = ({
    icon: { component: IconComponent, size: iconSize, display: displayIcon = true } = { component: <></> },
    title,
    isOpen,
    menus,
    ...props
}: ButtonGroupProps): JSX.Element => {
    const theme = useTheme();

    const [isWardrobeOpen, setIsWardrobeOpen] = useState(false);
    const toggleDrawer = useCallback(() => {
        if (isOpen) setIsWardrobeOpen(prevState => !prevState);
    }, [isOpen]);

    useEffect(() => {
        if (!isOpen && isWardrobeOpen) {
            setIsWardrobeOpen(false)
        }
    }, [isOpen])

    return (
        <Stack component="ul" gap={0.5}>
            <DefaultButtonLayout {...props} onClick={toggleDrawer}>
                {
                    displayIcon &&
                    <Box sx={{
                        display: "flex",
                        alignItems: "center",
                        justifyContent: "center",

                        opacity: 0.7,
                        padding: "0.4rem",
                    }}>
                        {
                            cloneElement(IconComponent, {
                                style: { fontSize: iconSize || 21 },
                                color: theme.palette.primary.purpleDark,
                            })
                        }
                    </Box>
                }
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
            </DefaultButtonLayout>
            {
                isWardrobeOpen &&
                <Box sx={{ paddingLeft: 2.6 }}>
                    <SideMenuButtons 
                        gap={2}
                        isOpen={isOpen} 
                        title={""}
                        showIcon={false} 
                        menus={menus} 
                        fontWeight={350} 
                    />
                </Box>
            }
        </Stack>
    )
}

export default ButtonGroup;