import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import { styled } from '@mui/material';
import type { BoxProps }  from '@mui/material'

const SolidButton = styled(({ children, disabled, onClick, ...props }: BoxProps & { disabled?: boolean }) => (
    <Box {...props} onClick={disabled ? () => {}: onClick}>
        <Typography variant="subtitle2" component="p" color="white" sx={{ margin: 0, padding: 0}}>
            { children }
        </Typography>
    </Box>  
))(({ theme, disabled }) => ({
    display: 'flex',
    justifyContent: "center",
    textWrap: "nowrap",

    backgroundColor: disabled === true ? 'rgb(220,220,220)' : theme.palette.primary.purpleLight,
    border: disabled === true ? 'solid 1px rgb(180,180,180)' : 'none',
    borderRadius: '22px',
    padding: '8px 22px',
    
    fontSize: '14px',
    fontWeight: 'bold',
    color: 'white',

    flex: '0 0 auto',

    '&:hover': disabled === true ? {} : {
        transform: "scale(1.03)",    
    },
    transition: "transform 150ms linear",
        
    cursor: disabled === true ? 'not-allowed' : "pointer",
    userSelect: "none",
}))

export default SolidButton;