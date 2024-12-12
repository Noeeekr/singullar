import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import { styled } from '@mui/material';
import type { BoxProps }  from '@mui/material'

const SolidButton = styled(({ children, ...props }: BoxProps) => (
    <Box {...props}>
        <Typography variant="subtitle2" component="p" color="white" sx={{ margin: 0, padding: 0}}>
            { children }
        </Typography>
    </Box>  
))(({ theme }) => ({
    display: 'flex',
    justifyContent: "center",
    
    backgroundColor: theme.palette.primary.purpleLight,
    borderRadius: '22px',
    padding: '8px 22px',
    
    fontSize: '14px',
    fontWeight: 'bold',
    color: 'white',

    '&:hover': {
        transform: "scale(1.03)",    
    },
    transition: "transform 150ms linear",
        
    cursor: "pointer",
    userSelect: "none",
}))

export default SolidButton;