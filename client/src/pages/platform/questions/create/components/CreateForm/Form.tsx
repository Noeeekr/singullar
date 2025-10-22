import FirstSection, { FirstSectionFormState } from "./FirstSection"
import SecondSection, { SecondSectionFormState } from "./SecondSection"
import ThirdSection, { ThirdSectionFormState } from "./ThirdSection"
import Typography from "@mui/material/Typography"
import Header from "./Header"
import Stack from "@mui/material/Stack"

import { styled } from "@mui/material/styles"

import { useMemo, createContext, useState, useCallback, JSX } from "react"
import { useForm, UseFormSetValue } from "react-hook-form"
import SolidButton from "@components/buttons/Button/Solid"

export type FormStateSectionProps<OverridableProps = {}> = OverridableProps & {
    complete?: boolean
}
export type FormSectionKeys = keyof FormContext["formState"]
export interface FormState {
    "1"?: FormStateSectionProps<FirstSectionFormState>
    "2"?: FormStateSectionProps<SecondSectionFormState>
    "3"?: FormStateSectionProps<ThirdSectionFormState>
}
export type CallbackRegister = (cb: () => void) => void

export interface FormContext {
    setFormState: UseFormSetValue<FormContext["formState"]>,
    setActiveSection: (section: FormSectionKeys) => void
    formState: FormState
}
export const FormContext = createContext<FormContext>({} as FormContext)

const defaultValues: FormContext["formState"] = {}

const TitleContainer = styled(Stack)(() => ({
    alignItems: "center",
    justifyContent: "space-between",

    backgroundColor: "rgba(130, 130, 130, 0.1)",
    padding: 18,
    border: "solid 1px rgb(190,190,190)",
    borderBottom: "none",
    borderTopLeftRadius: 16,
    borderTopRightRadius: 16,
}))
const SectionContainer = styled(Stack)(() => ({
    backgroundColor: "rgba(190,190,190,0.1)",
    padding: 18,
    border: "solid 1px rgb(190,190,190)",
    borderBottomLeftRadius: 6,
    borderBottomRightRadius: 6,
    borderTopLeftRadius: 0,
    borderTopRightRadius: 0,
}))

export default function (): JSX.Element {
    const [activeSection, setActiveSection] = useState<FormSectionKeys>("2")

    // **Bad code warning** The way this is designed makes it harder to have a filter button, but I wanted it anyway 
    const [childrenFunc, onParentClick] = useState<(() => void) | null>(null)

    const { watch, setValue } = useForm<FormContext["formState"]>({ defaultValues: defaultValues })
    const section = useMemo(() => {
        // **Bad code warning** Resets the callback so other sections don't get the button because of the previous value.
        onParentClick(null)
        switch (activeSection) {
            case "1":
                return <FirstSection title="Informações Básicas" />
            case "2":
                return <SecondSection title="Adicionar questões" parentButtonTitle="Pesquisar" onParentClick={onParentClick} />
            case "3":
                return <ThirdSection />
        }
    }, [activeSection])

    const handleActiveSection = useCallback(() => setActiveSection(prev => {
        const activeSection = Number(prev)
        if (activeSection == 3) return prev
        return `${activeSection + 1}` as FormSectionKeys
    }), [])

    return (
        <Stack gap={2}>
            <Header activeStep={Number(activeSection)} />
            <FormContext.Provider value={{
                formState: watch(),
                setActiveSection: (section: FormSectionKeys) => setActiveSection(section),
                setFormState: setValue,
            }}>
                <Stack>
                    <TitleContainer direction="row">
                        <Typography variant="h5" fontWeight="bold" component="h6">{section.props?.title}</Typography>
                        <Stack direction="row" gap={2}>
                            {
                                childrenFunc
                                    ? <SolidButton
                                        title={section.props.parentButtonTitle || ""}
                                        onClick={() => childrenFunc()}
                                    />
                                    : <></>
                            }
                            <SolidButton
                                title={activeSection == "3" ? "Criar lista" : "Continuar"}
                                disabled={watch("1")?.complete != true}
                                onClick={handleActiveSection}
                            />
                        </Stack>
                    </TitleContainer>
                    <SectionContainer>
                        {section}
                    </SectionContainer>
                </Stack>
            </FormContext.Provider>
        </Stack>
    )
}