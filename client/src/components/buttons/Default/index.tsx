import { StackProps } from "@mui/material";

import DefaultButton from "./Button";
import PaperButton from "./Paper";
import SolidButton from "./Solid";

export interface ButtonIconProps {
    icon?: {
        component: JSX.Element
        display?: boolean,
        size?: number,
        style?: object
    }
}
export interface ButtonEffectsProps {
    effects?: {
        enableSelectEffect?: boolean,
        enableHoverEffect?: boolean,
        enableNotificationEffect?: boolean,
    }
}
export interface ButtonLayoutProps extends ButtonEffectsProps {
    isMobile?: boolean
}
export interface ButtonProps extends ButtonLayoutProps, ButtonIconProps, StackProps {
    title: string,

    showDescription?: boolean,
    description?: string,

    type?: string,
}
export interface ButtonsProps extends ButtonProps {
    variant?: "paper" | "button" | "solid";
}

function Button(props: ButtonsProps): JSX.Element {
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

export default Button;