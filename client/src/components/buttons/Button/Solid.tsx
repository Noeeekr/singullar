import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import { styled } from '@mui/material';
import { ButtonProps } from './Button';

interface SolidButtonProps extends ButtonProps {
    disabled?: boolean
}

const SolidButton = styled(({ title, disabled, onClick, ...props }: SolidButtonProps) => (
    <Stack {...props} onClick={disabled ? () => {}: onClick}>
        <Typography variant="subtitle2" component="p" color={props.color ? String(props.color) : "white"} sx={{ margin: 0, padding: 0 }}>
            { title }
        </Typography>
    </Stack>  
))(({ theme, disabled }) => ({
    display: 'flex',
    justifyContent: "center",
    alignItems: "center",
    textWrap: "nowrap",

    backgroundColor: disabled === true ? 'rgb(220,220,220)' : theme.palette.primary.purpleLight,
    border: disabled === true ? 'solid 1px rgb(180,180,180)' : 'none',
    borderRadius: '22px',
    padding: '8px 22px',
    height: "min-content",

    fontSize: '14px',
    fontWeight: 'bold',
    color: 'white',

    flex: '0 0 auto',

    transition: "transform 150ms linear",
    '&:hover': disabled === true ? {} : {
        "-moz-osx-font-smoothing": "grayscale",
        "-webkit-font-smoothing": "antialiased",

        transform: "scale(1.1) translateZ(0)",    
    },

    cursor: disabled === true ? 'not-allowed' : "pointer",
    userSelect: "none",
}))

export default SolidButton;