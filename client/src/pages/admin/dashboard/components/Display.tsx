// Components
import Typography from "@mui/material/Typography"
import Grid from "@mui/material/Grid2";

// Models
import { Grid2Props } from "@mui/material"

export default ({ title, children, ...props }: Grid2Props): JSX.Element => {
    return (
        <Grid container spacing={1} {...props}>
            <Grid size={6}>
                <Typography variant="body2" fontWeight="bold" textTransform="capitalize">
                    {title}
                </Typography>
            </Grid>
            <Grid size={6}>
                <Typography variant="body2">
                    {children}
                </Typography>
            </Grid>
        </Grid>
    )
}