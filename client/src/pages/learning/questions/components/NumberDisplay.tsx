import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

import { styled } from "@mui/material";

import type { BoxProps } from "@mui/material"

export interface NumberDisplay extends BoxProps {
    level: number
}

function colorMapper(level: number): string {
    switch(level) {
        case 3:
            return "linear-gradient(rgb(240,100,100), rgb(130,40,40))"
        case 2:
            return "linear-gradient(rgb(200,160,100), rgb(140,110,60))"
        case 1:
            return "linear-gradient(rgb(120,200,120), rgb(90,140,90))"
        case 0:
        default:
            return "linear-gradient(rgb(190,240,180), rgb(130,190,130))"
    }
}
const NumberDisplay = styled(({ children, ...props}: NumberDisplay) => (
    <Box {...props}>
        <Typography variant="body2" color="rgb(255,255,255,1)" fontWeight="bold">
            { children }
        </Typography>
    </Box>
))(({ level }) => ({
    position: "relative",

    display: "flex",
    alignItems: "center",
    justifyContent: "center",

    backgroundImage: colorMapper(level),
    width: "40px",
    height: "50px",
    borderRadius: "0.5rem",
    userSelect: "none",

    overflow: "hidden",
}))

export default NumberDisplay;