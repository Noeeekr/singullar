import SolidButton from "@components/buttons/Default/Solid"
import Typography from "@mui/material/Typography"
import Stack, { StackProps } from "@mui/material/Stack"

import { JSX, createContext } from "react"
import CircularButton from "@components/buttons/Circular"

export interface SecondSectionContextProps {
    question_description?: string
    question_short_description?: string
    question_correct_alternative?: string
    alternatives?: string[]
}
export interface SecondSectionProps {
    update: (values: SecondSectionContextProps | null) => void
}

export const SecondSectionContext = createContext<SecondSectionContextProps>({})

export type SelectableProps<Selectable extends boolean> = Selectable extends true ? {
    onSelection: (title: string) => void
    isSelected: boolean
} : {
    onSelection?: (title: string) => void
    isSelected?: boolean
}

export type RemovableProps<Removable extends boolean> = Removable extends true ? {
    onDelete: (title: string) => void

} : {
    onDelete?: (title: string) => void
}

export type RequiredProps<Selectable extends boolean> = {
    title: string
    decorativeIconIndex: number

    removable?: boolean,
    selectable?: Selectable,
}

export type AlternativeDisplayerProps<
    Removable extends boolean = false,
    Selectable extends boolean = false,
> =
    StackProps
    & RequiredProps<Selectable>
    & SelectableProps<Selectable>
    & RemovableProps<Removable>

export default <
    Removable extends boolean,
    Selectable extends boolean,
>({
    title,
    onClick,
    onDelete,
    removable,
    selectable,
    isSelected,
    onSelection,
    decorativeIconIndex,
    ...props
}: AlternativeDisplayerProps<Removable, Selectable>): JSX.Element => {
    return (
        <Stack {...props} direction="row" alignItems="center" justifyContent="space-between">
            <Stack direction="row" alignItems="center" justifyContent="space-between" gap={2}>
                <CircularButton
                    sx={(theme) => ({
                        width: 40,
                        height: 40,
                        paddingTop: 0.2,
                        borderColor: isSelected ? "white" : selectable ? "rgb(170,170,170)" : 'solid 2px rgba(235,235,235)',
                        backgroundColor: isSelected ? theme.palette.primary.purpleDark : selectable ? "rgb(249,249,249)" : 'rgba(235,235,235)',
                        color: isSelected ? "white" : "rgb(90,90,90)",
                        "&:hover": selectable ? {
                            borderColor: theme.palette.primary.purpleLight,
                            backgroundColor: theme.palette.primary.purpleLight,
                            color: "white",
                        } : {
                            borderColor: isSelected ? "white" : selectable ? "rgb(170,170,170)" : 'solid 2px rgba(235,235,235)',
                            backgroundColor: isSelected ? theme.palette.primary.purpleDark : selectable ? "rgb(249,249,249)" : 'rgba(235,235,235)',
                            color: isSelected ? "white" : "rgb(90,90,90)",
                        },
                    })}
                    disabled={!isSelected && !selectable}
                    onClick={(e) => {
                        if (onClick) onClick(e)
                        if (selectable && onSelection) onSelection(title)
                    }}
                >
                    {String.fromCharCode(('a'.charCodeAt(0) + decorativeIconIndex)).toUpperCase()}
                </CircularButton>
                <Typography
                    variant="body2"
                    component="p"
                >
                    {title}
                </Typography>
            </Stack>
            {
                removable
                    ? <SolidButton
                        sx={{
                            backgroundColor: "rgb(240,170,170)",
                            border: "solid 1px rgb(200,50,50)",

                            transition: "transform 200ms ease-in-out, background-color 200ms ease-in-out, border 200ms ease-in-out",
                            "&:hover": {
                                backgroundColor: "rgb(210,130,130)",
                                border: "solid 1px rgb(150,20,20)",
                            }
                        }}
                        title="Remover alternativa"
                        onClick={() => {
                            if (removable && onDelete) onDelete(title)
                        }}
                    />
                    : <></>
            }
        </Stack>
    )
}