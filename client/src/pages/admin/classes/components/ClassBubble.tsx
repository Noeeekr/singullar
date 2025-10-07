import Typography from "@mui/material/Typography";

// Models
import { Class } from "@models/server";

import Grid from "@mui/material/Grid2";
import { useCallback, useEffect, useState } from "react";

export interface SelectableComponentProps<OnSelectProps = never> {
    selectable?: number,
    onSelect?: (props?: OnSelectProps) => void,
}
export interface ClassBubbleProps extends SelectableComponentProps<Class[]> {
    classes: Class[] | null
}
export default function ClassBubble(
    { classes, selectable, onSelect }: ClassBubbleProps
): JSX.Element {
    if (classes == null) {
        return (
            <Typography
                marginY={2}
                textAlign="center"
                fontWeight="bold"
                variant="body2"
                component="p"
            >
                Nenhuma turma encontrada
            </Typography>
        )
    }

    const [selectedClasses, setSelectedClasses] = useState<Class[]>([])
    const handleSelect = useCallback((cl: Class) => {
        if (!selectable) return;
        let classes = selectedClasses.filter((cls) => cls.id != cl.id)
        if (classes.length == selectedClasses.length) {
            classes.push(cl)
            if (selectedClasses.length >= selectable) {
                return
            }
        }
        setSelectedClasses(classes)
    }, [selectedClasses, selectable])

    useEffect(() => {
        if (onSelect) onSelect(selectedClasses);
    }, [selectedClasses, onSelect])

    return (
        <Grid container spacing={2} flexDirection="column" paddingY={2}>
            {
                classes.map(c => {
                    const isSelected = selectedClasses.some(cl => cl.id == c.id)
                    return (
                        <Grid
                            key={c.id}
                            sx={{
                                cursor: selectable ? "pointer" : "initial",
                                flex: 1,
                                bgcolor: isSelected ? "primary.purpleLight" : selectable == selectedClasses.length ? "rgb(220,220,220)" : "white",
                                transition: "all 150ms ease-in-out",
                                padding: 2,
                                borderRadius: 2,
                                boxShadow: "1px 1px 4px 1px rgb(190,190,190,0.3)",
                            }}
                            onClick={() => { handleSelect(c) }}
                        >
                            <Typography variant="body1" fontWeight="bold" color={isSelected ? "white" : "black"}>{c.name}</Typography>
                            <Typography variant="body2" color={isSelected ? "rgb(210,210,210)" : "gray"}>{c.segment} {c.series}</Typography>
                        </Grid>
                    )
                })
        }
        </Grid>
    )
}