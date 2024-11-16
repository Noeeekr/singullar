// Features
import {
    useState,
} from 'react'

// Components
import {
    Box,
    Stack,
    styled,
} from '@mui/material';

// Types 
import type {
    BoxProps,
    StackProps,
} from '@mui/material';

const SwitchBannerIcon = styled(({ children, onClick, ...props }: BoxProps) =>
    <Box onClick={onClick} {...props}>{ children }</Box>
)(() => ({
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',

    backgroundColor: 'rgba(250,250,250)',
    width: 35,
    height: 35,
    padding: 5,
    borderRadius: '50%',

    cursor: 'pointer',

    overflow: 'hidden',

    transition: 'transform 150ms ease-in-out',
    opacity: 0.8,

    '&:hover': {
        opacity: 0.9,
        transform: 'scale(1.05)'
    }
}))

const SliderWrapper = styled(({ children, ...props }: StackProps) => 
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

const BannerWrapper = styled(({ children, ...props }: BoxProps) => (
    <Box {...props}>{ children }</Box>
))(() => ({
    position: 'relative',

    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    flexShrink: 0,

    width: '100%',
    height: '100%',

    transition: 'left 200ms linear',
}))

const BannerSlider = (props: StackProps): JSX.Element => {
    const [BoxLeftPosition, setBoxLeftPosition] = useState(0);

    const bannerImages = ["rgb(200,0,60)", "rgb(50,190,90)", "rgb(190,169,3)"];

    const toggleBoxPosition = (side: "left" | "right") => {
        if (side === "left") {
            if (BoxLeftPosition >= 0) return;
            setBoxLeftPosition(side => side + 100);
        } else {
            if (BoxLeftPosition <= -((bannerImages.length - 1) * 100)) return;
            setBoxLeftPosition(side => side - 100);
        }
    }

    return (
        <SliderWrapper {...props}>
            {
                bannerImages.map((content, i) => (
                    <BannerWrapper
                        key={i}
                        sx={{
                            backgroundColor: content,
                            left: BoxLeftPosition + "%"
                        }}
                    >

                        Banner image
                    </BannerWrapper>
                ))
            }
            <Stack
                direction="row"
                sx={{
                    position: 'absolute',
                    left: '-12%',

                    justifyContent: "space-between",

                    width: '124%',

                    transition: 'all 180ms ease-in-out'
                }}
            >
                <SwitchBannerIcon onClick={() => (toggleBoxPosition("left"))}>
                    &lt;
                </SwitchBannerIcon>
                <SwitchBannerIcon onClick={() => (toggleBoxPosition("right"))}>
                    &gt;
                </SwitchBannerIcon>
            </Stack>
        </SliderWrapper>
    )
}

export default BannerSlider