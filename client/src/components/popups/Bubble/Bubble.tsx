import Paper from "@mui/material/Paper"

import { PopupIconProps, PopupProps } from ".."

export interface BubbleProps extends PopupProps, PopupIconProps { }

/**
 * 
            <Icon isOpen={isOpen} {...props} />
 */
const BubbleVariant = ({ isOpen, children }: BubbleProps) => {
    return (
        <Paper
            elevation={3}
            sx={{
                position: "absolute",
                top: 55,
                right: 0,

                backgroundColor: 'white',
                height: 'auto',
                maxWidth: 500,
                padding: 1,
                borderRadius: 5,

                transition: "all 200ms ease-in-out",
                pointerEvents: isOpen ? "all" : "none",
                transform: isOpen ? "scale(1)" : "scale(0.97)",
                opacity: isOpen ? 1 : 0,

                margin: 2,
            }}
        >
            {children}
        </Paper>
    )
}

export default BubbleVariant;
