import {
    Box,
    Stack,
    Typography,
} from '@mui/material'
import { CloseRounded } from '@mui/icons-material'

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

    return(
            <>
                    <Box
                        width='100vw'
                        height='100vh'

                        sx={{
                            backgroundColor: 'rgba(20,20,20,0.4)',
                            overflow: 'hidden',

                            position: 'absolute',
                            top: '0',
                            left: '0',

                            opacity: isOpen ? '1' : '0',
                            pointerEvents: isOpen ? 'all' : 'none',
                        }}

                        onClick={togglePopup}
                    >
                    </Box>
                    <Box
                        width='40vw'
                        height='100vh'

                        sx={{
                            position: 'absolute',
                            top: 0,
                            right: 0,

                            backgroundColor: 'white',
                            minWidth: '500px',
                            paddingX: 4,
                            paddingY: 2,

                            zIndex: 2,

                            opacity: isOpen ? '1' : '0',
                            pointerEvents: isOpen ? 'all' : 'none',
                        }}
                    >
                        <Stack
                            direction="row"
                            justifyContent="space-between"
                            alignItems='center'
                        >
                            <Typography
                                variant="h4"
                                fontWeight={600}
                            >
                                { title }
                            </Typography>
                            <Box
                                sx={{
                                    display: 'flex',
                                    alignItems: 'center',
                                    justifyContent: 'center',

                                    backgroundColor: 'rgba(0,0,0,0)',
                                    width: 42,
                                    height: 42,
                                    borderRadius: '50%',
                                    border: (theme) => `solid 2px ${theme.palette.primary.purpleExtraLight}`,
                                    '&:hover': {
                                        border: 'solid 2px rgba(235,235,235)',
                                        backgroundColor: 'rgba(235,235,235)',
                                    },

                                    transition: 'border 130ms linear, background 130ms linear',

                                    cursor: 'pointer',
                                }}
                                onClick={togglePopup}
                            >
                                <CloseRounded style={{ fontSize: 21 }}/>
                            </Box>
                        </Stack>
                        {children}
                    </Box>
                </>
    )
}

export default SidePopup