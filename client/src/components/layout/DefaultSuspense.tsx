import { styled, Typography } from "@mui/material"
import Stack from "@mui/material/Stack"

import { keyframes } from "@mui/system"
import { Suspense, SuspenseProps } from "react"

const animation = keyframes`
    from {
        transform: rotate(0deg);
    }
    to {
        transform: rotate(360deg);    
    }
`

const Ball = styled(Stack)<{ thickness: number | string }>(({ theme, thickness }) => ({
    position: "absolute",
    top: "1px",
    left: "40%",

    backgroundColor: theme.palette.primary.purpleLight,

    // Prevents small scale overflow
    border: "solid 1% white",

    width: `calc(${thickness} - 1px)`,
    height: `calc(${thickness} - 1px)`,
    borderRadius: "50%",
}))

const BallContainer = styled(Stack)(() => ({
    position: "absolute",
    top: 0,
    left: 0,

    width: "100%",
    height: "100%",

    animation: `${animation} 2s linear infinite`,
}))

const CContainer = styled(Stack)(() => ({
    transform: "rotate(180deg)",
}))

const CShape = styled(Stack)<{ thickness: number | string }>(({ theme, thickness }) => ({
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
    animation: `${animation} 3s linear infinite`,
}))

export default ({ size = 80, thickness = '20px', ...props }: SuspenseProps & { size?: number | string, thickness?: number | string }) => {
    return (
        <Suspense {...props} fallback={
            <Stack
                sx={(theme) => ({
                    backgroundColor: theme.palette.primary.paperLight,
                    width: "100%",
                    height: "100%",
                    alignItems: "center",
                    gap: 2,
                    justifyContent: "center",
                })}
            >
                <Stack
                    sx={() => ({
                        padding: 6,
                        alignItems: "center",
                        justifyContent: "center",
                        borderRadius: 5,
                        gap: 2,
                    })}
                >
                    <Typography variant="subtitle1" fontWeight="bold" sx={(theme) => ({
                        fontSize: 29,
                        fontWeight: 600,
                        fontFamily: "system-ui",
                        color: theme.palette.primary.purpleLight,
                    })}>
                        Carregando...
                    </Typography>
                    <Stack sx={{ position: "relative", aspectRatio: 1, width: size }}>
                        <BallContainer sx={{ animationDelay: '0ms' }}>
                            <Ball thickness={thickness} />
                        </BallContainer>
                        <BallContainer sx={{ animationDelay: '300ms' }}>
                            <Ball thickness={thickness} />
                        </BallContainer>
                        <BallContainer sx={{ animationDelay: '600ms' }}>
                            <Ball thickness={thickness} />
                        </BallContainer>
                        <BallContainer sx={{ animationDelay: '900ms' }}>
                            <Ball thickness={thickness} />
                        </BallContainer>
                        <CContainer>
                            <CShape thickness={thickness} />
                        </CContainer>
                    </Stack>
                </Stack>
            </Stack>
        }/>
    );
};