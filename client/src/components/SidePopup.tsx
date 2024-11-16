// Components
import Box from '@mui/material/Box'
import Stack from '@mui/material/Stack'
import Typography from '@mui/material/Typography'
import CircularButton from './ButtonCircular'

// Icons
import { CloseRounded } from '@mui/icons-material'

// Features
import { styled } from '@mui/material'

// Types
import type { BoxProps } from '@mui/material'

const ShadowBackground = styled(Box)(({ isOpen }: { isOpen: boolean }) => ({
    position: 'absolute',
    top: '0',
    left: '0',
    
    backgroundColor: 'rgba(20,20,20,0.4)',
    width: '100vw',
    height: '100vh',
    
    overflow: 'hidden',
    zIndex: 10,

    opacity: isOpen ? '1' : '0',
    pointerEvents: isOpen ? 'all' : 'none',
}))

const SidePopupWrapper = styled(({ children, onClick, ...props }: BoxProps & { isOpen?: boolean }) => <>
    <Box onClick={onClick} {...props}>
        {children}
    </Box>
</>)(({ isOpen }) => ({
    position: 'absolute',
    top: 0,
    right: 0,

    backgroundColor: 'white',
    width: '40vw',
    minWidth: '500px',
    height: '100vh',
    padding: '30px',

    zIndex: 10,

    opacity: isOpen ? '1' : '0',
    pointerEvents: isOpen ? 'all' : 'none',
}))

interface INavbarItemGroupProps {
    // for toggle menu open click event handling
    isOpen: boolean,
    onClickCb?: Function,
    id?: string,

    children: JSX.Element,
    title: string, // for small title popup

    structure?: "side" | "bubble"
}

const SidePopup = (props: INavbarItemGroupProps) => {
    const { children, isOpen, onClickCb, title, id } = props;

    const togglePopup = () => { if (onClickCb) onClickCb(id); }

    return (
        <>
            <ShadowBackground
                isOpen={isOpen}
                onClick={togglePopup}
            />
            <SidePopupWrapper
                isOpen={isOpen}
            >
                <Stack
                    direction="row"
                    justifyContent="space-between"
                    alignItems='center'
                    paddingBottom={2}
                >
                    <Typography
                        variant="h4"
                        fontWeight={600}
                    >
                        {title}
                    </Typography>
                    <CircularButton
                        onClickCb={togglePopup}
                    >
                        <CloseRounded style={{ fontSize: 21 }} />
                    </CircularButton>
                </Stack>
                {children}
            </SidePopupWrapper>
        </>
    )
}

export default SidePopup