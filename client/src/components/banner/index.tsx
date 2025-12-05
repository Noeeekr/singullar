// Components
import IconContainer from "./IconContainer";
import Banner from "./Banner";
import Slider from "./Slider"
import Box from "@mui/material/Box"

// Utilities
import { useState } from 'react'
import { styled } from '@mui/material';

// Models
import type { BoxProps } from "@mui/material/Box"
import type { StackProps } from "@mui/material/Stack"

const SlideIconWrapper = styled(({ children, onClick, ...props }: BoxProps) =>
    <Box onClick={onClick} {...props}>{children}</Box>
)(() => ({
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',

    backgroundColor: 'rgba(250,250,250)',
    width: 50,
    height: 50,
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

const BannerSlider = (props: StackProps): JSX.Element => {
    const [offset, setOffset] = useState(0);

    const images = ["rgba(212, 69, 69, 0.9)", "rgb(50,190,90)", "rgb(190,169,3)"];
    const colors = ['rgba(242, 109, 109, 0.9)', 'rgba(50, 190, 90, 0.9)', 'rgba(240, 200, 60, 0.9)'];

    const slideTo = (side: "left" | "right") => {
        switch (side) {
            case "left":
                if (offset >= 0) return;
                setOffset(side => side + 100);
                return
            default:
                if (offset <= -((images.length - 1) * 100)) return;
                setOffset(side => side - 100);
                return
        }
    }

    return (
        <Slider {...props}>
            {
                images.map((color, i) => (
                    <Banner
                        key={i}
                        offset={offset}
                        color={colors[i]}
                        background={color}
                    >
                        Banner Title
                    </Banner>
                ))
            }
            <IconContainer>
                <SlideIconWrapper onClick={() => (slideTo("left"))}>
                    &lt;
                </SlideIconWrapper>
                <SlideIconWrapper onClick={() => (slideTo("right"))}>
                    &gt;
                </SlideIconWrapper>
            </IconContainer>
        </Slider>
    )
}

export default BannerSlider