// Components
import Grid from "@mui/material/Grid2"

// Utilities
import { styled } from "@mui/material";

// Models
import type { QuestionListDisplayStylingProps } from "..";

export default styled(({ children, navigable, ...props }: QuestionListDisplayStylingProps) => (
    <Grid size={4} {...props}>{children}</Grid>
))(({ theme, selectable, navigable }) => ({
    position: "relative",

    backgroundColor: "rgba(255,255,255,0.5)",
    height: "100%",
    minHeight: "200px",
    width: "100%",
    minWidth: "350px",
    padding: "20px",
    border: `solid 1px ${theme.palette.grey[400]}`,
    borderRadius: "1rem",

    cursor: "pointer",

    transition: "all 150ms ease-in-out",
    translate: "0px 0px",
    transform: "scale(1)",

    "&:hover": selectable ? {
        backgroundColor: `rgba(138, 59, 217, 0.19)`,
        border: `solid 1px ${theme.palette.primary.purpleLight}`,
    } : navigable ? {
        translate: "0px -5px",
        transform: "scale(1.01)",
    } : {}
}))