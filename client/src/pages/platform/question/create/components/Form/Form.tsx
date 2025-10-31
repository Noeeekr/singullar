import MultiStepForm from "@components/forms/MultiStepForm"
import FirstSection, { FirstSectionContext } from "./FirstSection"
import SecondSection, { SecondSectionContext } from "./SecondSection"

export default function(): JSX.Element {

    const sections = [
        {
            content: FirstSection,
            
            title: "Informações Básicas",
            
            context: FirstSectionContext,
            defaultValues: { 
                questionTitle: "Contexto não iniciado",
                difficultyLevel: undefined,
            }
        },
        {
            content:SecondSection,
            
            title: "Conteúdo da questão",

            context: SecondSectionContext,
            defaultValues: { 
                questionDescription: "Contexto não iniciado",
            }
        },
    ]

    return(
        <div>
            <MultiStepForm sections={sections} />
        </div>
    )
}