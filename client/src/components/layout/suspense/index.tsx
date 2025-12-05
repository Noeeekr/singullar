// Components
import InnerContainer from "./InnerContainer"
import OuterContainer from "./OuterContainer"
import SnakeContainer from "./Snake/Container"
import { Suspense } from "react"
import Snake from "./Snake/Snake"
import Stack from "@mui/material/Stack"
import Text from "./Text"

// Models
import type { ReactNode, SuspenseProps as ReactSuspenseProps } from "react"
import MovingBallContainer from "./Snake/BallContainer"
import MovingBall from "./Snake/Ball"

export interface SuspenseProps extends ReactSuspenseProps {
    /** Number or css width */
    size?: number | string,
    /** Css pixel size */
    thickness?: string,
    /** Element to be displayed when suspense ends, can also be done with children */
    to?: ReactNode
}

export default ({
    to,
    size = 80,
    thickness = '20px',
    children,
    ...props
}: SuspenseProps) => {
    return (
        <Suspense {...props} fallback={
            <OuterContainer>
                <InnerContainer>
                    <Text variant="subtitle1">
                        Carregando...
                    </Text>
                    <Stack sx={{ position: "relative", aspectRatio: 1, width: size }}>
                        <MovingBallContainer sx={{ animationDelay: '0ms' }}>
                            <MovingBall thickness={thickness} />
                        </MovingBallContainer>
                        <MovingBallContainer sx={{ animationDelay: '300ms' }}>
                            <MovingBall thickness={thickness} />
                        </MovingBallContainer>
                        <MovingBallContainer sx={{ animationDelay: '600ms' }}>
                            <MovingBall thickness={thickness} />
                        </MovingBallContainer>
                        <MovingBallContainer sx={{ animationDelay: '900ms' }}>
                            <MovingBall thickness={thickness} />
                        </MovingBallContainer>
                        <SnakeContainer>
                            <Snake thickness={thickness} />
                        </SnakeContainer>
                    </Stack>
                </InnerContainer>
            </OuterContainer>
        }>
            {to || children}
        </Suspense>
    );
};