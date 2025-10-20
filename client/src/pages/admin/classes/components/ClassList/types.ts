import type { Grid2Props } from "@mui/material/Grid2"
import { Class } from "@models/server/server";

export interface SelectableComponentProps<OnSelectProps = never> {
    selectable?: number,
    onSelect?: (props?: OnSelectProps) => void,
}
export interface ClassListProps extends SelectableComponentProps<Class[]> {
    classes: Class[] | null
    navegable?: boolean
}
export type ClassListItemProps = SelectableComponentProps<Class[]> & Grid2Props & {
    class: Class
    isFull?: boolean
    isSelected?: boolean
    navegable?: boolean
}