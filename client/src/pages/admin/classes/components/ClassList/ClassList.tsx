// Components
import Grid from "@mui/material/Grid2";
import ClassListItem from "./ClassListItem"

// Features
import { useCallback, useEffect, useState } from "react";

// Types
import type { ClassListProps } from "./types";
import type { Class } from "@models/server/server";
import ErrorBubble from "@components/bubbles/ErrorBubble/ErrorBubble";
import { useNavigate } from "react-router-dom";

export default function ClassBubble(
    { classes, selectable, onSelect, navegable }: ClassListProps
): JSX.Element {
    if (classes == null) return <ErrorBubble err={"Nenhuma turma encontrada"} />
    const [selectedClasses, setSelectedClasses] = useState<Class[]>([])
    const navigate = useNavigate();
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

    const handleInteraction = (c: Class) => {
        if (selectable) handleSelect(c);
        if (navegable) navigate(`${c.id}`); 
    }

    return (
        <Grid container spacing={2} direction="column" paddingY={2}>
            {
                classes.map(c => {
                    const isSelected = selectedClasses.some(cl => cl.id == c.id)
                    return (
                        <ClassListItem 
                            class={c} 
                            key={c.id} 
                            onClick={() => handleInteraction(c)} 
                            navegable={navegable}
                            isSelected={isSelected} 
                            isFull={selectable == selectedClasses.length}
                        />
                    )
                })
        }
        </Grid>
    )
}