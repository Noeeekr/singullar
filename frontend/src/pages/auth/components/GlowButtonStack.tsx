import {
    Box,
    Stack,
    Divider,
    ButtonBase,
    Typography,
} from '@mui/material'

import { ReactNode } from 'react'

const TextGlowButton = (
    { children, isActive, onClick, label }: 
    { isActive: Boolean, children: ReactNode, onClick: Function, label: string }
): JSX.Element => {
    return (
        <Box
            component="div"
        >
            <ButtonBase 
                onClick={() => (onClick(label))}
                disableRipple={true}
            >
                <Typography variant="body2" component="p" sx={(theme) => {
                    let colors = theme.palette.primary;
                    return {
                        ":hover": {
                            color: colors.lightPurple
                        },
                        fontSize: '0.9rem',
                        color: isActive ? colors.lightPurple : colors.lightGray
                    }
                }}>
                    {children}
                </Typography>
            </ButtonBase>

            {isActive
                ? <Box
                    sx={(theme) => {
                        return {
                            backgroundColor: theme.palette.primary.darkPurple,
                            width: "100%",
                            height: '3.4px',
                            translate: "0 6px",
                            borderRadius: 3
                        }
                    }}
                ></Box>
                : <></>
            }
        </Box>
    )
}
const GlowButtonStack = (
    { active, labels, sx, onClick }:
    { active: ReactNode, labels: string[], sx: Object, onClick: Function },
): JSX.Element => {
    return (
        <Box sx={sx}>
            <Stack
                component="nav"
                direction="row"
                spacing={3}
                sx={(theme) => {
                    return {
                        color: theme.palette.primary.lightGray,
                    }
                }}
                marginBottom={0.5}
            >
                {labels.map((label) => (
                    <TextGlowButton
                        isActive={label === active}
                        onClick={onClick}
                        label={label}
                        key={label}
                    >
                        {label}
                    </TextGlowButton>
                ))}
            </Stack>

            <Divider
                sx={{
                    width: '120%',
                    translate: '-10% 0px',
                    marginBottom: 3
                }}
            />
        </Box>

    ) // Might become a component under customButton folder or smt
}

export default GlowButtonStack