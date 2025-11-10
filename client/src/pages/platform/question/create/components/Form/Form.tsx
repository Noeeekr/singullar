import MultiStepForm from "@components/forms/MultiStepForm"
import FirstSection, { FirstSectionContext, FirstSectionContextProps } from "./FirstSection"
import SecondSection, { SecondSectionContext, SecondSectionContextProps } from "./SecondSection"

    const sections = [
        {
            content: FirstSection,
            
            title: "Informações Básicas",
            
            context: FirstSectionContext,
            defaultValues: { 
                questionTitle: undefined as string | undefined,
                difficultyLevel: undefined as number | undefined,
            }
        },
        {
            content:SecondSection,
            
            title: "Conteúdo da questão",

            context: SecondSectionContext,
            defaultValues: { 
                questionDescription: undefined as undefined | string,
                alternatives: [] as string[] | undefined
            }
        },
    ]

export default function(): JSX.Element {
    const onSubmit = (formData: FirstSectionContextProps & SecondSectionContextProps) => {
        console.log(formData)
    }
    return(
        <MultiStepForm sections={sections} onSubmit={onSubmit}/>
    )
}