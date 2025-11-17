// Components
import Box, { BoxProps } from '@mui/material/Box'
import Stack from '@mui/material/Stack'
import Typography from '@mui/material/Typography'
import CircularButton from '../buttons/Circular'

// Features
import { styled } from '@mui/material'
import { lazy, Suspense } from 'react'

// Types
import { PopupProps } from '.'

// Icons
const CloseRounded = lazy(() => import("@mui/icons-material/CloseRounded"))

export interface ShadowBackgroundProps extends BoxProps {
    isOpen?: boolean
}

const ShadowBackground = styled(({ isOpen, ...props }: ShadowBackgroundProps) => <Box {...props}></Box>)(({ isOpen }: { isOpen: boolean }) => ({
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

export interface SidePopupWrapperProps extends BoxProps {
    isOpen?: boolean
}

const SidePopupWrapper = styled(({ children, isOpen, ...props }:SidePopupWrapperProps) => <>
    <Box {...props}>
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

const SidePopup = ({
    isOpen,
    onClose,
    title,
    element
}: PopupProps & BoxProps) => {
    return (
        <>
            <ShadowBackground onClick={onClose} isOpen={isOpen ? isOpen : false}/>
            <SidePopupWrapper isOpen={isOpen}>
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
                    <CircularButton onClick={onClose}>
                        <Suspense fallback={<></>}>
                            <CloseRounded style={{ fontSize: 21 }} />
                        </Suspense>
                    </CircularButton>
                </Stack>
                {element}
            </SidePopupWrapper>
        </>
    )
}

export default SidePopup