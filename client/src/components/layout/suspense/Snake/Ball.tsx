import Stack from "@mui/material/Stack"

import { styled } from "@mui/material"

export default styled(Stack)<{ thickness: string }>(({ theme, thickness }) => ({
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