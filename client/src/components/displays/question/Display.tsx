// Components
import LargeQuestionDisplayer from "./components/LargeDisplayer"
import MediumQuestionDisplayer from "./components/MediumDisplayer"
import SmallQuestionDisplayer from "./components/SmallDisplayer"

// Utilities
import { LARGE, MEDIUM } from "."

// Models
import type { DisplayerVariants, DisplayerProps } from "."

export default <IsSelected extends boolean>({ type, ...props }: DisplayerProps<IsSelected> & DisplayerVariants): JSX.Element => {
    switch (type) {
        case LARGE:
            return <LargeQuestionDisplayer {...props} />
        case MEDIUM:
            return <MediumQuestionDisplayer {...props} />
        default:
            return <SmallQuestionDisplayer {...props} />
    }
}