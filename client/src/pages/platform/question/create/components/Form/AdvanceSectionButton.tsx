import type { FirstSectionContextProps } from "./FirstSection"

import SolidButton from "@components/buttons/Default/Solid"

export default function(): JSX.Element {
    const context: any = {}
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