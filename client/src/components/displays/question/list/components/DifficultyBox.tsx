import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

import { styled } from "@mui/material";

import type { BoxProps } from "@mui/material"

export interface NumberDisplayProps extends BoxProps {
    level: number
}

function colorMapper(level: number, opacity?: number): string {
    if (level < 4) {
        return `rgb(100,170,100, ${opacity ? opacity : 1})`
    }
    if (level < 7) {
        return "rgb(70,100,70, 0.7)"
    }
    if (level < 10) {
        return "rgb(140,110,60)"
    }
    return "rgb(130,40,40)"
}
function gradientMapper(level: number): string {
    if (level < 4) {
        return "linear-gradient(rgb(100,200,100, 0.2), rgb(90,190,90, 0.4))"
    }
    if (level < 7) {
        return "linear-gradient(rgb(120,200,120, 0.7), rgb(70,110,70, 0.7))"
    }
    if (level < 10) {
        return "linear-gradient(rgb(200,160,100, 0.4), rgb(140,110,60, 0.6))"
    }
    return "linear-gradient(rgb(240,100,100, 0.4), rgb(130,40,40 ,0.6))"
}
export const NumberDisplay = styled(({ children, ...props }: NumberDisplayProps) => (
    <Box {...props}>
        <Typography variant="body2" color={colorMapper(props.level)} fontWeight="bold">
            {children}
        </Typography>
    </Box>
))(({ level }) => ({
    position: "relative",

    display: "flex",
    alignItems: "center",
    justifyContent: "center",

    backgroundImage: gradientMapper(level),
    width: "44px",
    height: "50px",
    borderRadius: "0.6rem",
    userSelect: "none",
    border: `solid 1px ${colorMapper(level, 0.4)}`,

    overflow: "hidden",
}))

export default ({ amount, level }: NumberDisplayProps & { amount: number }) => {
    return (
        <>
            {
                new Array(amount).map((_, i) => (
                    <NumberDisplay key={i} level={level}>{i + 1}</NumberDisplay>
                ))
            }
        </>
    )
};