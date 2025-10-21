// Components
import Box from '@mui/material/Box'

// Icons 
import { FaChevronLeft } from 'react-icons/fa'

// Features
import { styled } from '@mui/material';

// Types
import type { BoxProps } from '@mui/material/Box';

const CircularButton = styled(({ children, fontSize, ...props }: BoxProps) => (
    <Box {...props}>
        {
            children
                ? children
                : <FaChevronLeft fontSize={fontSize ? Number(fontSize) : 13} />
        }
    </Box>
))(({ theme }) => ({
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