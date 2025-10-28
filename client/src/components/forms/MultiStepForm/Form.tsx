// Components
import SectionContainer from "./SectionContainer"
import TitleContainer from "./TitleContainer"
import Typography from "@mui/material/Typography"
import Stack from "@mui/material/Stack"
import Header from "./Header"
import { Fragment } from "react"

// Features
import { useMemo, useState, useCallback, createContext } from "react"

// Models
import type { StackProps } from "@mui/material/Stack"
import type { JSX, Context } from "react"


/**
 * @type "FormState" refers to the type which contains all underlying sections types.
 */
export interface FormContextDefaultValues<FormState> {
    isLastSection: boolean,
    setActiveSection: (modifier: (n: number) => number) => void,
    formState: FormState,
}

export type FormContextSectionValues<FormState extends SectionState, SectionState> = (FormContextDefaultValues<FormState> & SectionState) | undefined

/**
 * @param SectionContext A context that enables passing values to external components like in "header" component.
 */
export interface FormSection<FormState extends SectionState, SectionState> {
    title: string,

    content: () => JSX.Element,
    header?: (context: FormContextSectionValues<FormState, SectionState>) => JSX.Element,

    SectionContext?: Context<FormContextSectionValues<FormState, SectionState>>,

    defaultValues?: SectionState,
}

/**
 * @type "FormState" is used to define the type of values passed through context in the field "FormState" and its default values.
 * @type "Section" is used to allow each section to have its own underlying type
 */
export interface FormProps<Section extends FormSection<FormState, any>, FormState> extends StackProps {
    sections: Section[]
}

export const staticValues: FormContextSectionValues<{}, any> = {
    setActiveSection: (_: (_: number) => number) => {},
    isLastSection: false,
    formState: {},
}

function RenderDelayer<FormState extends SectionsState, SectionsState = any>({ render, context }: { context?: FormContextSectionValues<FormState, SectionsState>, render: ((context: FormContextSectionValues<FormState, SectionsState>) => JSX.Element) | undefined }): JSX.Element {
    return (
        <Fragment>
            {
                render ?
                    render(context)
                    : <></>
            }
        </Fragment>
    )
}

/**
 * Generic Form Interface
 *   @summary 
 *   Can be used to create N sections that get N independent contexts that all merge into one giant context.  
 *   Sections can provide components to the interface outside it's scope and still talk with the main section component through context without rerendering everything.
 *  
 *  Note: I cannot even call this complex-atrocity-fault-intolerant-unchecked-enormous-multi-state-machine "bad code", this is apex coding right here. There is even room for improvement. May create the millionest react lib just from the premise of this.
 *
 *  @type "SectionsState" refers to the union type with all the sections types, used to specify all allowed section types. 
 *  @type "FormState" is meant to be used to specify the union of all types into a shared state.
 * 
 *  @param sections refers to the sections the form will load.
 */
export default function <SectionsState, FormState extends SectionsState>({ sections, ...props }: FormProps<FormSection<FormState, SectionsState>, FormState>): JSX.Element {
    const [activeSection, setActiveSection] = useState(0)

    const { SectionContext = createContext<SectionsState | undefined>(undefined), defaultValues, header, title, content } = useMemo(() => {
        return sections[activeSection]
    }, [activeSection, sections])

    const addValidation = useCallback((setter: (n: number) => void) => {
        return (modifier: (n: number) => number) => {
            const num = modifier(activeSection)
            if (num == sections.length) return;
            if (num == 0) return;
            setter(num)
        }
    }, [])

    const providerValues: FormContextSectionValues<FormState, SectionsState> = useMemo(() => {
        if (defaultValues) return {
            ...defaultValues,
            isLastSection: activeSection == sections.length - 1,
            formState: {} as FormState,
            setActiveSection: addValidation(setActiveSection)
        }
        return undefined
    }, [])

    return (
        <Stack gap={2} {...props}>
            <Header headers={[]} activeStep={activeSection} />
            <SectionContext.Provider value={providerValues}>
                <Stack>
                    <TitleContainer direction="row">
                        <Typography variant="h5" fontWeight="bold" content="h6">{title}</Typography>
                        <Stack direction="row" gap={2}>
                            <RenderDelayer context={providerValues} render={header} />
                        </Stack>
                    </TitleContainer>
                    <SectionContainer>
                        <RenderDelayer render={content} />
                    </SectionContainer>
                </Stack>
            </SectionContext.Provider>
        </Stack>
    )
}