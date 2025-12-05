import { keyframes, styled } from "@mui/material"
import Stack from "@mui/material/Stack"

export default styled(Stack)<{ thickness: number | string }>(({ theme, thickness }) => ({
    '--v': 50,    // [0 100]
    '--b': thickness, // thickness

    '--_v': 'min(99.99,var(--v))',
    '--_f': 'round(down,var(--_v),50)',

    width: "100%",
    aspectRatio: '1',
    display: 'grid',
    containerType: 'size',

    '&::before': {
        content: '""',
        background: theme.palette.primary.purpleLight,

        clipPath: `shape(from top,
                            arc to calc(50% + 50%*sin(var(--_v)*3.6deg)) 
                            calc(50% - 50%*cos(var(--_v)*3.6deg)) of 50% cw var(--_c,large),
                            arc to calc(50% + (50% - var(--b))*sin(var(--_v)*3.6deg)) 
                            calc(50% - (50% - var(--b))*cos(var(--_v)*3.6deg)) of 1% cw,
                            arc to 50% var(--b) of calc(50% - var(--b)) var(--_c,large),
                            arc to top of 1% cw
                        )`,

        '@container style(--_f: 0)': {
            '--_c': 'small',
        }
    },
    animation: `${keyframes`
        from {
            transform: rotate(0deg);
        }
        to {
            transform: rotate(360deg);    
        }`
        } 3s linear infinite`,
}))
