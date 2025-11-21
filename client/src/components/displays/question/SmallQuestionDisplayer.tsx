import { Divider, Paper, Stack, styled, Typography, TypographyProps } from "@mui/material"

import type { JSX } from "react"
import type { 
    DisplayerProps, 
    ToggleSelectableDisplayerProps, 
    NonSelectableDisplayerProps
} from "@components/displays/question"

const ContainerText = styled(({ children, isSelected, selectable, onSelected, ...props }: TypographyProps & ToggleSelectableDisplayerProps<boolean>) => (
    <Typography {...props}>
        {children}
    </Typography>
))<NonSelectableDisplayerProps>(({ isSelected }) => ({
    whiteSpace: "nowrap",
    overflow: "hidden",
    textOverflow: "ellipsis",
    msTextOverflow: "ellipsis",
    lineClamp: 4,
    WebkitLineClamp: 4,
    color: isSelected ? "white" : "black",
}))

const Container = styled(Paper)<NonSelectableDisplayerProps>(({ theme, selectable, isSelected }) => ({
    position: "relative",
    top: 0,

    backgroundColor: isSelected ? theme.palette.primary.purpleLight : "white",
    minWidth: 340,
    minHeight: 120,
    padding: 10,
    borderRadius: 10,

    cursor: selectable ? "pointer" : "initial",

    transition: "all 100ms ease",
    transform: "scale(1)",

    "&:hover": {
        boxShadow: "rgba(0,0,0, 0.2) 0px 2px 4px 2px",
        transform: "scale(1.01)",
        top: "-2px",
    }
}))

export default <IsSelected extends boolean>({
    selectable,
    onSelected,
    isSelected,
    onClick,
    question,
    ...props
}: DisplayerProps<IsSelected>): JSX.Element => {
    return (
        <Container elevation={2} isSelected={isSelected} onClick={(e) => {
            if (onSelected) onSelected(question.id, question)
            if (onClick) onClick(e)
        }} selectable={selectable}>
            <Stack gap={1} {...props}>
                <Stack direction="row" justifyContent="space-between">
                    <ContainerText width="50%" isSelected={isSelected}>
                        {question.question_title}
                    </ContainerText>
                    <ContainerText isSelected={isSelected}>
                        {question.question_difficulty_level}
                    </ContainerText>
                </Stack>
                <Divider color={isSelected ? "white" : ""} />
                <ContainerText isSelected={isSelected}>
                    {question.question_description}
                </ContainerText>
            </Stack>
        </Container >
    )
}