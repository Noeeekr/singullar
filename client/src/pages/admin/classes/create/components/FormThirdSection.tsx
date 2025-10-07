import { useCallback, useContext } from "react"
import SearchTeachers from "../../../teachers/components/Search"
import { FormContext } from "./Form"

export interface FormThirdSectionData {
    "teacher_id": number | null
}

export default function FormThirdSection(): JSX.Element {
    const { setFormData } = useContext(FormContext);
    const handleSelect = useCallback((ids: number[]) => {
        setFormData("thirdSection.teacher_id", ids.length ? ids[0] : null)
    },[])

    return (
        <SearchTeachers selectable={1} onSelect={handleSelect}/>
    )
}