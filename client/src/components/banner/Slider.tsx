// Components
import Stack from "@mui/material/Stack"

// Utilities
import { styled } from '@mui/material';

// Models
import type { StackProps } from "@mui/material/Stack"

export default styled(({ children, ...props }: StackProps) => 
    <Stack direction="row" {...props}>{ children }</Stack>
)(({ theme }) => ({
    position: 'relative',
    alignItems: 'center',

    backgroundColor: theme.palette.primary.purpleDark,
    width: '100vw',
    maxWidth: '100%',
    
    height: 'min(30vw,300px)',
    borderRadius: 26,

    cursor: 'pointer',
    userSelect: 'none',
    
    marginRight: 0,
    
    overflow: 'hidden',
    transition: 'all 500ms ease-in-out',
    float: 'none',

    zIndex: 1,
    
    '&:hover .MuiStack-root': {
        left: '5%',
        width: '90%',
    }
}))