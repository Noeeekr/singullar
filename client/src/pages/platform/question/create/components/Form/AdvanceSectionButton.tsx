import type { FirstSectionContextProps } from "./FirstSection"
import type { FormContextSectionValues } from "@components/forms/MultiStepForm/Form"
import type { FormContextProps } from "./Form"

import SolidButton from "@components/buttons/Default/Solid"

export default function({ context }: { context: FormContextSectionValues<FormContextProps, FirstSectionContextProps>}): JSX.Element {
    return(
        <SolidButton 
            title="Continuar"
            disabled={context?.isLastSection == true}
            onClick={() => {
                if (!context?.setActiveSection) return;
                context.setActiveSection((prev) => ++prev)
            }}
        />
    )
}