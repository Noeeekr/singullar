// Components
import LargeQuestionDisplayer from "./LargeQuestionDisplayer"
import MediumQuestionDisplayer from "./MediumQuestionDisplayer"
import SmallQuestionDisplayer from "./SmallQuestionDisplayer"

// Models
import type { Question } from "@models/server"
import type { JSX } from 'react'
import type { PaperProps } from "@mui/material"

export type DisplayerOptions = { type?: DisplayerTypes }
export type DisplayerTypes = typeof SMALL | typeof MEDIUM | typeof LARGE
export type DisplayerConfigurationProps = {
    question: Question,
}

// Question displayer props
export type DisplayerProps<IsSelected extends boolean> =
    DisplayerConfigurationProps
    & DisplayerStylingProps
    & ToggleSelectableDisplayerProps<IsSelected>

export type DisplayerStylingProps = DisplayerContainerStylingProps
export type DisplayerContainerStylingProps = PaperProps

export type NonSelectableDisplayerProps = {
    selectable?: true,
    isSelected?: boolean,

    // Optional prop that may become necessary if selectable == true
    onSelected?: (id: number, question: Question) => void
}
export type SelectableDisplayerProps = {
    selectable: true,
    isSelected: boolean,
    // Called when a question is selected or de-selected
    onSelected: (id: number, question: Question) => void
}

export type ToggleSelectableDisplayerProps<IsSelectable extends boolean> =
    IsSelectable extends true
    ? SelectableDisplayerProps & NonSelectableDisplayerProps
    : NonSelectableDisplayerProps

export const SMALL = "sm"
export const MEDIUM = "md"
export const LARGE = "lg"

export default <IsSelected extends boolean>({ type, ...props }: DisplayerProps<IsSelected> & DisplayerOptions): JSX.Element => {
    switch (type) {
        case LARGE:
            return <LargeQuestionDisplayer {...props} />
        case MEDIUM:
            return <MediumQuestionDisplayer {...props} />
        default:
            return <SmallQuestionDisplayer {...props} />
    }
}