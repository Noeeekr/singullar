// Components
import Box from "@mui/material/Box"

// Utilities
import { styled } from '@mui/material';

// Models
import type { BoxProps } from "@mui/material/Box"

export interface BannerProps extends BoxProps {
    // From 0 - 100 (number), to be used as (%)
    offset?: number
    // Css background color property, to be set as fallback for the banner
    background: string
}
export default styled(({ children, ...props }: BannerProps) => (
    <Box {...props}>{children}</Box>
))(({ background, offset = 0 }) => ({
    position: 'relative',
    left: offset + "%",

    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    flexShrink: 0,

    textTransform: "uppercase",
    fontWeight: 600,
    fontSize: 'min(9vw,60px)',
    letterSpacing: 2,
    color: "rgb(0,0,0,0.3)",

    backgroundColor: background,
    width: '100%',
    height: '100%',

    transition: 'left 200ms linear',
}))