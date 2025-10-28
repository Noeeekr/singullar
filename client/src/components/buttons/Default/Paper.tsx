import Typography from "@mui/material/Typography"
import Stack from "@mui/material/Stack"
import Box from "@mui/material/Box"

import { styled, useTheme } from "@mui/material";
import { cloneElement } from "react";

import type { ButtonProps } from ".";

const PaperButtonLayout = styled(Stack)(() => ({
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'space-between',
    alignItems: 'flex-start',

    backgroundColor: "white",
    height: '120px',
    borderRadius: '17px',
    boxShadow: 'rgba(114, 119, 128, 0.09) 0px 1px 0px 0px,rgba(114, 119, 128, 0.09) 0px 2px 4px 0px, rgba(114, 119, 128, 0.09) 0px 4px 8px 0px',
    padding: '20px',
    border: 'solid 1px rgb(230,230,230)',

    transition: 'boxShadow 0ms linear',
    '&:hover': {
        boxShadow: 'rgba(114, 119, 128, 0.09) 0px 1px 0px 0px, rgba(114, 119, 128, 0.09) 0px 2px 4px 0px, rgba(114, 119, 128, 0.09) 0px 4px 8px 0px, rgba(114, 119, 128, 0.09) 0px 8px 16px 0px, rgba(114, 119, 128, 0.09) 0px 12px 24px 0px',

        cursor: 'pointer',
    }
}));


const PaperButton = ({
    icon: { display: displayIcon = true, component: IconComponent, size: iconSize } = { component: <></> },
    fontWeight,
    showDescription,
    description,
    title,
    ...props
}: ButtonProps): JSX.Element => {
    const theme = useTheme();

    return (
        <PaperButtonLayout {...props}>
            {
                displayIcon ? 
                <Box
                    sx={{
                        display: "flex",
                        alignItems: "center",
                        justifyContent: "center",

                        opacity: 0.7,
                        padding: "0.4rem",
                    }}
                >
                    {
                        cloneElement(IconComponent, {
                            style: { fontSize: iconSize || "1.5rem" },
                            color: theme.palette.primary.purpleDark,
                        })
                    }
                </Box>
                : <></>
            }
            <Stack gap={0.5}>
                <Typography variant="body2" fontWeight={fontWeight || 500}>
                    {title}
                </Typography>
                {
                    showDescription && description
                        ? (
                            <Typography variant="body1" sx={{ textWrap: 'wrap', minWidth: 150, paddingX: 1 }}>
                                {description}
                            </Typography>
                        )
                        : <></>
                }
            </Stack>
        </PaperButtonLayout>
    )
}

export default PaperButton;