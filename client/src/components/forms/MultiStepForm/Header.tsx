import StepLabel from "@mui/material/StepLabel"
import Stepper, { StepperProps } from "@mui/material/Stepper"
import Check from '@mui/icons-material/Check';
import Step from "@mui/material/Step"

// Features
import { JSX } from "react"
import { styled } from "@mui/material"

export interface SectionHeader {
    label: string
    icon?: React.ElementType<any, keyof JSX.IntrinsicElements>
}
export interface MultiStepFormHeaderProps extends StepperProps {
    headers: SectionHeader[]
}

export interface IconProps {
    ownerState: { completed?: boolean, active?: boolean }
}

const Icon = styled(({ ownerState, ...props }: IconProps) => (
    <Check {...props} />
))(({ theme }) => ({
    backgroundColor: "rgba(0,0,0,0.1)",
    padding: "0.2rem",
    borderRadius: "50%",
    boxSizing: "content-box",

    fontSize: "1.2rem",
    color: "white",

    transition: "background-color 200ms ease-in-out",
    variants: [
        {
            props: ({ ownerState }) => ownerState.active,
            style: {
                backgroundColor: theme.palette.primary.purpleExtraLight,
            }
        },
        {
            props: ({ ownerState }) => ownerState.completed,
            style: {
                backgroundColor: theme.palette.primary.purpleLight,
            }
        }
    ]
}))

export default function ({ headers, ...props }: MultiStepFormHeaderProps): JSX.Element {
    return (
        <Stepper alternativeLabel {...props}>
            {
                headers.map(({ label, icon }) => (
                    <Step>
                        <StepLabel slots={{ stepIcon: icon ? icon : Icon }}>
                            {label}
                        </StepLabel>
                    </Step>
                ))
            }
        </Stepper>
    )
}