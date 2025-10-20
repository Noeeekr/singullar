// Features
import { cloneElement } from 'react'
import { useTheme } from '@mui/material/styles'
import { styled } from "@mui/material"

// Components
import Box from "@mui/material/Box"
import Stack from "@mui/material/Stack"
import Typography from "@mui/material/Typography"

import type { ButtonLayoutProps, ButtonProps } from './Button'

/**
 * Holds the icon, the text, and the arrow of each Link / Group
 */
export const DefaultButtonLayout = styled(({ children, ...props }: ButtonLayoutProps) => (
    <Stack direction="row" component="li" {...props} >
        { children }
    </Stack>
))(({ theme, isMobile, effects }) => ({
    alignItems: "center",
    
    backgroundColor: effects?.enableSelectEffect ? theme.palette.primary.purpleLightInv : 'transparent',
    paddingY: isMobile ? "2.5rem" : "2rem",
    paddingX: "10px",
    borderRadius: 6,
    boxSizing: "content-box",

    textWrap: 'nowrap',
    color: "black",
    
    cursor: "pointer",
    overflow: 'hidden',
    "&:hover > .MuiBox-root": {
        opacity: 1,
        transition: 'opacity 200ms ease-in-out',
    },
    "&:hover > .MuiBox-root > .MuiBox-root": {
        opacity: 1,
        transition: 'opacity 200ms ease-in-out',
    },
    "&:active": {
        backgroundColor: theme.palette.primary.purpleLightInv
    },
    "&:hover": effects?.enableHoverEffect ? {
        backgroundColor: theme.palette.primary.purpleLightInv,
    } : {},
}))




// Note: Change locals that use href to use Selected={true}
// Note: Change whatever is loading buttons to use useMediaQuery(theme.breakpoints.down('xs')) to get isMobile value
const Button = ({
    icon: { component: IconComponent, size: iconSize, display: displayIcon = true } = { component: <></> },
    title,
    description, showDescription,
    fontWeight,
    ...props
}: ButtonProps): JSX.Element => {
    const theme = useTheme();
    return (
        <DefaultButtonLayout {...props}>
            {
                displayIcon ? 
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
                        cloneElement(IconComponent, {
                            style: { fontSize: iconSize || "1.5rem" },
                            color: theme.palette.primary.purpleDark,
                        })
                    }
                </Box>
                : <></>
            }
            <Stack gap={0.2} paddingY="0.4rem">
            {
                Boolean(title) ?
                <Typography variant="body2" fontWeight={fontWeight || 500} fontSize={15} sx={{ paddingX: 0.5 }}>
                    {title}
                </Typography>
                : <></>
            }
            {
                showDescription && Boolean(description)
                    ? (
                        <Typography variant="body1" sx={{ textWrap: 'wrap', minWidth: 150, paddingX: 1 }}>
                            {description}
                        </Typography>
                    )
                    : <></>
            }
            </Stack>
        </DefaultButtonLayout>
    )
}

export default Button;