// Components
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Fade from "@mui/material/Fade";

// Features
import { styled } from "@mui/material";

// Types
import type { BoxProps } from "@mui/material/Box";

const ErrorHelperText = styled(
  ({ show, children, ...props }: BoxProps & { show: boolean }) => (
    <Fade in={show}>
      <div style={{ position: "relative", marginBottom: show ? 20 : 1 }}>
        <Box {...props}>
          <Typography variant="body1" color="error.dark">
            {children}
          </Typography>
        </Box>
      </div>
    </Fade>
  )
)(({ theme }) => ({
  position: "absolute",
  bottom: -25,
  backgroundColor: theme.palette.error.whiteHigh,
  paddingX: 2,
  paddingTop: 2,
  paddingBottom: 0.5,
  borderEndEndRadius: 10,
  borderEndStartRadius: 10,
  zIndex: 1,
  width: "100%",
}));

export default ErrorHelperText;
