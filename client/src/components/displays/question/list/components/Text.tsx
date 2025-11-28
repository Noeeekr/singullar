// Components
import Typography from "@mui/material/Typography";

// Utilities
import { styled } from "@mui/material";

// Models
import type { TypographyProps } from "@mui/material/Typography"

export default styled(({ ...props }: TypographyProps) => (
    <Typography component="p" {...props} />
))(({
    textTransform: "capitalize",  
}))