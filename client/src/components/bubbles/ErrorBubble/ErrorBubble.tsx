// Components 
import Box from '@mui/material/Box';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import {
    HighlightOff
} from '@mui/icons-material'

import { styled } from '@mui/material';

const BubbleContainer = styled(Box)<BubbleContainerPrimaryColor>(({ theme, primary }) => ({
    backgroundColor: primary ? primary : `${theme.palette.error.light}`,
    padding: 16,
    borderRadius: 12,

    marginY: 2,
}))

const BubbleInnerContainer = styled(Stack)({
    alignItems: 'center',
})

const Bubble = styled(Box)<BubbleContainerSecondaryColor>(({ theme, secondary }) => ({
    position: 'relative',
    display: 'flex',
    '&:before': {
        position: 'absolute',
        top: -16,
        content: '""',
        backgroundColor: secondary ? secondary : theme.palette.error.main,
        width: '80%',
        height: 5,
        marginLeft: '2px',
        borderRadius: 2,
    }
}))

export interface BubbleContainerPrimaryColor {
    primary?: string
}
export interface BubbleContainerSecondaryColor {
    secondary?: string
}
export interface ErrorBubbleProps extends BubbleContainerPrimaryColor, BubbleContainerSecondaryColor {
    err?: string
    // Literally the same as error, but I'm not into debugging every err message so I'll leave both 
    message: string
} 

export default ({ err, message, primary, secondary }: ErrorBubbleProps): JSX.Element => {
    if (message) err = message
    if (!err) return <></>
    return (
        <BubbleContainer primary={primary}>
            <BubbleInnerContainer direction="row" spacing={1}>
                <Bubble secondary={secondary}>
                    <HighlightOff htmlColor={secondary ? secondary : "error"} />
                </Bubble>
                <Typography color={secondary ? secondary : "error.dark"} variant="body1" fontWeight="bold">
                    {err}
                </Typography>
            </BubbleInnerContainer>
        </BubbleContainer>
    )
}