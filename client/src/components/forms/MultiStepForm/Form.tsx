// Components
import SectionContainer from "./SectionContainer"
import TitleContainer from "./TitleContainer"
import Typography from "@mui/material/Typography"
import Stack from "@mui/material/Stack"
import Header from "./Header"
import { Fragment } from "react"

// Features
import { useMemo, useState } from "react"

// Models
import type { JSX, Context } from "react"
import SolidButton from "@components/buttons/Default/Solid"

export interface Updater<Value> {
    update: (values: Value) => void
}
export type MultiStepFormSectionContent<SectionValue> = ({ update }: Updater<SectionValue>) => JSX.Element
export interface MultiStepFormSectionProps<Section> {
    // Components
    content: MultiStepFormSectionContent<Section>,

    // Static values
    title: string
    context: Context<Section>
    defaultValues: Section
}
export type MultiStepFormSections<Sections extends unknown[]> = MultiStepFormSectionProps<Sections[number]>[]
/**

    []Sections => [](A | B | C | D)
    [
        A,
        B,
        C,
        D,
    ]

*/
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
export type SectionUnionArray<T> = MultiStepFormSectionProps<T>[];

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
    sections 
}: { 
    sections: Sections
} ): JSX.Element {
    const [activeStep, setActiveStep] = useState(0)
    const [form, setForm] = useState<Sections>({} as Sections)
    const [disabled, setDisabled] = useState(true)

    const {
        context: SectionContext,
        defaultValues,
        content,
        title,
    }: MultiStepFormSectionProps<Sections> = useMemo(() => {
        return sections[activeStep]
    }, [activeStep])

    const updater = (values: Sections) => {
        for (const key in values) {
            if (values[key] === undefined) {
                setDisabled(true)
                return
            }
        }
        setDisabled(false)
        setForm(formState => ({ ...formState, ...values }))
    }

    return (
        <Stack marginY={2}>
            <Header headers={[]} activeStep={activeStep} />
            <SectionContext.Provider value={defaultValues as Sections}>
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
                                        onClick={() => console.log("submit", form) /* onSubmit(form) */ }
                                    />
                                    : <SolidButton
                                        title="Continuar"
                                        disabled={disabled}
                                        onClick={() => { setDisabled(true), setActiveStep(prev => prev+1) }}
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