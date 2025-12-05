import Stack from "@mui/material/Stack"

import { styled } from "@mui/material"

export default styled(Stack)(({ theme }) => ({
    backgroundColor: theme.palette.primary.paperLight,
    width: "100%",
    height: "100%",
    alignItems: "center",
    gap: 2,
    justifyContent: "center",
}))