import Stack from "@mui/material/Stack"

import { keyframes, styled } from "@mui/material"

export default styled(Stack)(() => ({
    position: "absolute",
    top: 0,
    left: 0,

    width: "100%",
    height: "100%",

    animation: `${keyframes`
        from {
            transform: rotate(0deg);
        }
        to {
            transform: rotate(360deg);    
        }
    `} 2s linear infinite`,
}))
