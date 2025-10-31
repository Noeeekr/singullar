import { useForm } from "react-hook-form"
import { createContext } from "react"

export interface SecondSectionContextProps {
    questionDescription?: string
}

export const SecondSectionContext = createContext<SecondSectionContextProps>({})

export default function({ 
    update 
}: { 
    update: (values: SecondSectionContextProps) => void 
}): JSX.Element {
    const { formState: { defaultValues } } = useForm<SecondSectionContextProps>({
        defaultValues: {
            questionDescription: undefined,
        }
    })      

    return(
        <div>
            {defaultValues ? "Valor padrão definido" : "Valor padrão indefinido"} -
            Segunda sessão
        </div>
    )
}