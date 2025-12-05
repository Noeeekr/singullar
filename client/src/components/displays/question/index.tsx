// Models
import type { Question } from "@models/server"
import type { PaperProps } from "@mui/material"

export type DisplayerVariants = { type?: DisplayerTypes }
export type DisplayerTypes = typeof SMALL | typeof MEDIUM | typeof LARGE
export type DisplayerConfigurationProps = {
    question: Question,
}

// Question displayer props
export type DisplayerProps<IsSelected extends boolean> =
    DisplayerConfigurationProps
    & DisplayerStylingProps<IsSelected>

export type DisplayerStylingProps<IsSelected extends boolean> =
    DisplayerContainerStylingProps
    & NavigableDisplayerContainerStylingProps
    & SelectableDisplayerProps<IsSelected>

export type DisplayerContainerStylingProps = PaperProps

// Navigable Displayer
export type NavigableDisplayerContainerStylingProps = {
    navigable?: boolean,
}
// Selectable Displayer
export type OptionallySelectableDisplayerProps = {
    selectable?: true,
    isSelected?: boolean,

    // Optional prop that may become necessary if selectable == true
    onSelected?: (id: number, question: Question) => void
}
export type SelectableDisplayer = {
    selectable: true,
    isSelected: boolean,
    // Called when a question is selected or de-selected
    onSelected: (id: number, question: Question) => void
}
export type SelectableDisplayerProps<IsSelectable extends boolean> =
    IsSelectable extends true
    ? SelectableDisplayer & OptionallySelectableDisplayerProps
    : OptionallySelectableDisplayerProps

// Displayer Variants 
export const SMALL = "sm"
export const MEDIUM = "md"
export const LARGE = "lg"

export { default } from "./NavigableDisplay"