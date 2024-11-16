import Box from '@mui/material/Box'

import { FaChevronLeft } from "react-icons/fa6";

import { styled } from '@mui/material';

import type { BoxProps } from '@mui/material/Box';

const CircularButton = styled(({ children, onClickCb, ...props }: BoxProps & { onClickCb: Function }) => <>
    {
        children 
        ? <Box {...props} onClick={onClickCb}>
            {children}
        </Box>
        : <Box {...props} onClick={onClickCb}>
            <FaChevronLeft/>
        </Box>
    }

</>)(({ theme }) => ({
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    
    backgroundColor: 'rgba(0,0,0,0)',
    width: 42,
    height: 42,
    border: `solid 2px ${theme.palette.primary.purpleExtraLight}`,
    borderRadius: '50%',
    
    '&:hover': {
        border: 'solid 2px rgba(235,235,235)',
        backgroundColor: 'rgba(235,235,235)',
    },
    transition: 'border 130ms linear, background 130ms linear',
    
    flexShrink: '0',

    cursor: 'pointer',
}))

export default CircularButton;