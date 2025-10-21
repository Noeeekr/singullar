import StepLabel from "@mui/material/StepLabel"
import Stepper, { StepperProps } from "@mui/material/Stepper"
import Check from '@mui/icons-material/Check';
import Step from "@mui/material/Step"

// Features
import { styled } from "@mui/material"

const Icon = styled(Check)<{ ownerState: { completed?: boolean, active?: boolean } }>(({ theme }) => ({
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

export default function ({ ...props }: StepperProps): JSX.Element {
    return (
        <Stepper alternativeLabel { ...props }>
            <Step>
                <StepLabel slots={{ stepIcon: Icon }}>
                    Informações básicas
                </StepLabel>
            </Step>
            <Step>
                <StepLabel slots={{ stepIcon: Icon }}>
                    Adicionar questões
                </StepLabel>
            </Step>
            <Step>
                <StepLabel slots={{ stepIcon: Icon }}>
                    Finalizar
                </StepLabel>
            </Step>
        </Stepper>
    )
}