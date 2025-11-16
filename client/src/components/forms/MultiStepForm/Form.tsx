// Components
import SectionContainer from "./SectionContainer"
import TitleContainer from "./TitleContainer"
import Typography from "@mui/material/Typography"
import Stack from "@mui/material/Stack"
import Header, { SectionHeader } from "./Header"
import { Fragment, useCallback } from "react"

// Features
import { useMemo, useState } from "react"

// Models
import type { JSX, Context } from "react"
import SolidButton from "@components/buttons/Default/Solid"

export interface Updater<Value> {
    update: (values: Value) => void
}
export type MultiStepFormSectionContent<SectionValue> = ({ update }: Updater<SectionValue>) => JSX.Element
export interface MultiStepFormSectionProps<SectionValue> {
    // Components
    content: MultiStepFormSectionContent<SectionValue>,

    // Static values
    title: string
    context: Context<SectionValue>
    defaultValues: SectionValue
}
export type MultiStepFormSections<Sections extends unknown[]> = MultiStepFormSectionProps<Sections[number]>[]

export interface MultiStepFormProps<Sections extends {}> {
    sections: MultiStepFormSectionProps<Sections>[],
    onSubmit: () => void,
}

export interface RenderDelayerProps {
    element: () => JSX.Element
}
export interface AdvanceStepButtonProps<SectionContextValue> {
    setActiveStep?: (value: number | ((prev: number) => number)) => void
    action: (context: SectionContextValue) => boolean
    context: Context<SectionContextValue>
}

// // Helper to compute the Intersection (A & B & C) from a tuple of types [A, B, C, ...]
// type Intersect<U> = (U extends any ? (k: U) => void : never) extends ((k: infer I) => void) ? I : never;

// // Helper to extract the context types from the array of section props (gets [A, B, C])
// type ContextsTuple<T extends MultiStepFormSectionProps<any>[]> = {
//     [K in keyof T]: T[K] extends MultiStepFormSectionProps<infer C> ? C : never
// };

// // The type for the final combined state object (A & B & C)
// type CombinedState<T extends MultiStepFormSectionProps<any>[]> = Intersect<ContextsTuple<T>[number]>;

function RenderDelayer({ element }: RenderDelayerProps): JSX.Element {
    return (
        <Fragment>
            {element()}
        </Fragment>
    )
}

export default function<Sections extends MultiStepFormSectionProps<any>[]>({ 
    headers = [],
    sections, 
    onSubmit,
    initialSection = 0,
}: { 
    headers?: SectionHeader[],
    sections: Sections,
    onSubmit: (formData: Sections[number]["defaultValues"]) => void
    initialSection?: number
} ): JSX.Element {
    type SectionValues = Sections[number]["defaultValues"]
    const [activeStep, setActiveStep] = useState(sections.length <= initialSection || 0 > initialSection ? 0 : initialSection)
    const [form, setForm] = useState<SectionValues>({} as SectionValues)
    const [disabled, setDisabled] = useState(true)

    const { context: SectionContext, defaultValues, content, title }: MultiStepFormSectionProps<Sections> = useMemo(() => {
        return sections[activeStep]
    }, [activeStep])

    /**
     * Updater function enables advancing for next section if given a value different than null.
     * If it gets null then it sets the form values of this section to its default values to prevent form to be sending the previous updates. Which is preferred to be zero values.
     * @param values The values to be appended to form final state
     * @returns void
     */
    const updater = (values: Sections | null) => {
        if (values == null) {
            setForm(formState => {
                return ({ ...formState, ...defaultValues })
            })
            setDisabled(true)
            return
        } 
        setDisabled(false)
        setForm(formState => ({ ...formState, ...values }))
    }

    const handleSubmit = useCallback(() => {
        onSubmit(form)
        return
        setDisabled(true)
        setForm({})
        setActiveStep(0)
    }, [onSubmit, form])

    const handleContinue = useCallback(() => {
        setDisabled(true)
        setActiveStep(prev => ++prev)
    }, [setDisabled, setActiveStep])

    return (
        <Stack marginY={2} component="section" gap={2}>
            <Header headers={headers} activeStep={activeStep} />
            <SectionContext.Provider value={defaultValues as SectionValues}>
                <Stack>
                    <TitleContainer direction="row">
                        <Typography variant="h5" fontWeight="bold" content="h6">{title}</Typography>
                        <Stack direction="row" gap={2}>
                            {
                                // Add + 1 (delete later (comment))
                                activeStep + 1 == sections.length
                                    ? <SolidButton
                                        title="Criar questão"   
                                        disabled={disabled}
                                        onClick={handleSubmit}
                                    />
                                    : <SolidButton
                                        title="Continuar"
                                        disabled={disabled}
                                        onClick={handleContinue}
                                    />
                            }
                        </Stack>
                    </TitleContainer>
                    <SectionContainer>
                        <RenderDelayer element={() => content({ update: updater })} />
                    </SectionContainer>
                </Stack>
            </SectionContext.Provider>
        </Stack>
    )
}