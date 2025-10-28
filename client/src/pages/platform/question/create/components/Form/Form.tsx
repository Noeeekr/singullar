import AdvanceSectionButton from "./AdvanceSectionButton"
import MultiStepForm from "@components/forms/MultiStepForm"
import FirstSection, { FirstSectionContext } from "./FirstSection"

import type { FormSection } from "@components/forms/MultiStepForm/Form"
import type { FirstSectionContextProps } from "./FirstSection"
export type FormContextProps = FirstSectionContextProps

export default function(): JSX.Element {
    const sections: FormSection<FormContextProps, FirstSectionContextProps>[] = [
        {
            content: FirstSection,
            header: (context) => AdvanceSectionButton({ context: context }),
            
            title: "Informações Básicas",
            
            SectionContext: FirstSectionContext,
            defaultValues: { title: "Titulo da sessão 1" }
        }
    ]

    return(
        <div>
            <MultiStepForm sections={sections} />
        </div>
    )
}