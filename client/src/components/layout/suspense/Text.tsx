import Typography from "@mui/material/Typography"

import { styled } from "@mui/material"

export default styled(Typography)(({ theme }) => ({
    fontSize: 29,
    fontWeight: 600,
    fontFamily: "system-ui",
    color: theme.palette.primary.purpleLight,
}))