import Box, { BoxProps } from "@mui/material/Box"

import { keyframes } from "@mui/material"

const animation = keyframes`
    0% {
        left: -300%;
    }
    100% {
        left: 0%;
    }
`

export interface FallbackStylingProps extends BoxProps {
    width?: number,
    height?: number,
}

export interface FallbackProps extends FallbackStylingProps {
    on: boolean
}

const Fallback = ({ width = 1, height = 0.8, ...props }: FallbackStylingProps): JSX.Element => {
    props.children = <></>

    return (
        <Box sx={{
            position: "relative",

            backgroundColor: 'rgb(230, 230, 230, 0.8)',
            width: 100 * width,
            height: 24 * height,
            borderRadius: 1.5,

            overflow: 'hidden',
        }} {...props}>
            <Box sx={{
                position: "absolute",
                left: 0,
                top: 0,

                backgroundImage: 'linear-gradient(110deg, rgba(255,255,25,0) 20%, rgba(255,255,255,0.5) 40%, rgba(255,255,255,0.2), rgba(255,255,255,0) 80%)',
                width: "400%",
                height: "100%",

                animation: `${animation} 2200ms ease-in-out infinite`,
            }} />
        </Box>
    )
}
export default ({
    on,
    children,
    ...props
}: FallbackProps) => {
    if (on) return <Fallback {...props} />
    return children
}