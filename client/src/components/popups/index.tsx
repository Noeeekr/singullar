import Side from "./Side" 
import Bubble from "./Bubble" 
import { StackProps } from "@mui/material"
import { ButtonIconProps, ButtonProps } from "@components/buttons/Default"

export interface PopupIconLabelProps extends StackProps {
    title: string
    isCorner?: boolean,
}
export interface PopupIconProps extends PopupIconLabelProps, ButtonIconProps  {
    hasNotifications?: boolean,
    isOpen?: boolean,
    iconVariant?: "small" | "large"
}
export interface PopupProps extends PopupIconProps, ButtonProps, StackProps  {
    title: string,
    isOpen?: boolean,

    // Triggered on close actions
    onClose?: () => void,
    element?: JSX.Element,
}
export interface PopupsProps extends PopupProps {
    variant: "bubble" | "side"
}

export default function Popup({ variant, ...props }: PopupsProps): JSX.Element {
    const Popup = (() => {
        switch (variant) {
            case "bubble":
                return <Bubble {...props} />
            default:
                return <Side {...props} />
        }
    })()

    return Popup
}