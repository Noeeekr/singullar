import { useState } from 'react'
import { useTheme } from '@mui/material/styles'

import { Suspense, lazy } from 'react'

const CloseRounded = lazy(() => import('@mui/icons-material/CloseRounded'))

/**
 * 
 * @param notifications Shows a small dot on icon if true
 * @param onClick a callback triggered by click events
 * 
 */
const IconButton = (
    { notifications, onClick }:
    { notifications?: boolean, onClick: () => void }
): JSX.Element => {
    const theme = useTheme()

    const [isOpen, setIsOpen] = useState(false)
    const [isHovered, setIsHovered] = useState(false)

    const toggleButtonOpen = () => (setIsOpen(prevState => !prevState))

    return (
        <div style={{
            position: 'relative',

            display: 'flex',
            alignItems: isOpen ? 'center' : 'initial' ,
            flexDirection: 'column',
            justifyContent: 'space-around',
            gap: 3,

            backgroundColor: isHovered ? 'rgba(155,215,255,0.3)' : 'rgba(0,0,0,0)',
            borderRadius: 5,

            width: 36,
            height: 36,

            boxSizing: 'border-box',
            padding: isOpen ? '0' : '10px 8px',

            flex: '0 0 36px',

            cursor: 'pointer',
        }}
            onClick={() => {
                toggleButtonOpen()
                onClick()
            }}
            onMouseEnter={() => (setIsHovered(true))}
            onMouseLeave={() => (setIsHovered(false))}
        >
            {   // Hand made Icons 
                isOpen
                    ? <Suspense fallback={<></>}>
                        <CloseRounded 
                            sx={{
                                fontSize: 21,
                                color: 'white'
                            }}
                        />
                    </Suspense>
                    : (
                        <>
                        <div style={{
                            backgroundColor: 'rgb(250,250,250,0.7)',
                            width: '100%',
                            flex: '1',
                            borderRadius: 5,
                        }}></div>   
                        <div style={{
                            backgroundColor: 'rgb(250,250,250,0.7)',
                            width: '100%',
                            flex: '1',
                            borderRadius: 5,
                        }}></div>
                        <div style={{
                            backgroundColor: 'rgb(250,250,250,0.7)',
                            width: '80%',
                            flex: '1',
                            borderRadius: 5,
                        }}></div>

            
                        {notifications
                            ? (<div style={{
                                position: 'absolute',
                                top: '6px',
                                right: '6px',
                            
                                content: '',
                                backgroundColor: theme.palette.primary.contrast,
                                width: 7.6,
                                height: 7.6,
                                borderRadius: 20
                            }}/>)
                            : <></>
                        }
                        </>
                    )
            }
        </div>
    )

}

export default IconButton