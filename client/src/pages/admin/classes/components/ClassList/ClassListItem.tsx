// Components
import Grid from "@mui/material/Grid2";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";

// Features
import styled from "@emotion/styled";

// Models
import type { ClassListItemProps } from "./types";

const ClassListItem = styled(({ children, class: { name, segment, series }, isFull, navegable, isSelected, ...props }: ClassListItemProps) => 
    <Grid { ...props }>
        <Typography variant="body1" color={isSelected ? "white" : "black"} fontWeight="bold">{name}</Typography>
        <Typography variant="body2" color={isSelected ? "rgb(210,210,210)" : "gray"}>{segment} {series}</Typography>
        <Box>
            <Typography variant="h6" component="p" fontWeight="bold" color="white">Ver mais</Typography>
        </Box>
        { children }
    </Grid>
)(({ theme, selectable, isSelected, isFull, navegable }) => ({
    ".MuiBox-root": {
        position: "absolute",
        right: "min(30%, -300px)",
        top: "0",

        display: navegable ? "flex" : "none",
        alignItems: "center",
        justifyContent: "flex-end",
        
        backgroundImage: `linear-gradient(90deg, transparent, ${theme.palette.primary.purpleExtraLight})`,
        width: "min(30%, 300px)",
        height: "100%",
        paddingRight: "0px",
        
        transition: "right ease-in-out 400ms, padding-right linear 500ms"
    },
    ".MuiTypography-root": {
        textTransform: "capitalize",
    },
    "&:hover": {
        ".MuiBox-root": {
            paddingRight: "20px",
            right: "0",
        }
    },
    position: "relative",
    flex: 1,
    
    backgroundColor: isSelected ? "primary.purpleLight" : isFull ? "rgb(220,220,220)" : "white",
    padding: "1rem",
    boxShadow: "1px 1px 4px 1px rgb(190,190,190,0.3)",
    borderRadius: "0.7rem",
    
    overflow: "hidden",
    transition: "all 150ms ease-in-out",
    cursor: selectable || navegable ? "pointer" : "initial",
}))

export default ClassListItem;