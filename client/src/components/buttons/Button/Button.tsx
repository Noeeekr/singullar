import { StackProps } from "@mui/material";

import DefaultButton from "./Default";
import PaperButton from "./Paper";
import SolidButton from "./Solid";

export interface ButtonIconProps {
    component: JSX.Element
    display?: boolean,
    size?: number,
}

export interface ButtonEffectsProps {
    enableSelectEffect?: boolean,
    enableHoverEffect?: boolean,
    enableNotificationEffect?: boolean,
}
export interface ButtonLayoutProps extends StackProps {
    isMobile?: boolean
    effects?: ButtonEffectsProps
}
export interface ButtonProps extends ButtonLayoutProps {
    title: string,
    isMobile?: boolean

    showDescription?: boolean,
    description?: string,

    icon?: ButtonIconProps
    type?: string,
}
export interface ButtonsProps extends ButtonProps {
    variant?: "paper" | "button" | "solid";
}

function Buttons(props: ButtonsProps): JSX.Element {
    switch (props.variant) {
        case "button":
            return <DefaultButton {...props} />
        case "paper":
            return <PaperButton {...props} />
        case "solid":
            return <SolidButton {...props} />
        default:
            return <DefaultButton {...props} />
    }
}

export default Buttons;