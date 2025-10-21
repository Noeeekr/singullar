import FirstSection, { FirstSectionProps } from "./FirstSection"
import SecondSection, { SecondSectionProps } from "./SecondSection"
import ThirdSection, { ThirdSectionProps } from "./ThirdSection"
import Typography from "@mui/material/Typography"
import Stack from "@mui/material/Stack"

import { styled } from "@mui/material/styles"

import { useMemo, createContext, useState, useCallback } from "react"
import { useForm, UseFormSetValue } from "react-hook-form"
import SolidButton from "@components/buttons/Button/Solid"

export type FormSectionProps<OverridableProps = {}> = OverridableProps & {
    complete?: boolean
}
export type FormSectionKeys = keyof FormContext["formState"]
export interface FormState {
    "1"?: FormSectionProps<FirstSectionProps>
    "2"?: FormSectionProps<SecondSectionProps>
    "3"?: FormSectionProps<ThirdSectionProps>
}
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
    const [activeSection, setActiveSection] = useState<FormSectionKeys>("1")
    const { watch, setValue } = useForm<FormContext["formState"]>({ defaultValues: defaultValues })
    const section: JSX.Element = useMemo(() => {
        switch (activeSection) {
            case "1":
                return <FirstSection title="Informações Básicas" />
            case "2":
                return <SecondSection />
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
        <FormContext.Provider value={{
            setActiveSection: (section: FormSectionKeys) => setActiveSection(section),
            formState: watch(),
            setFormState: setValue,
        }}>
            <Stack>
                <TitleContainer direction="row">
                    <Typography variant="h5" fontWeight="bold" component="h6">{section.props.title}</Typography>
                    <SolidButton
                        title={activeSection == "3" ? "Criar lista" : "Continuar"}
                        disabled={watch("1")?.complete != true}
                        onClick={handleActiveSection}
                    />
                </TitleContainer>
                <SectionContainer>
                    {section}
                </SectionContainer>
            </Stack>
        </FormContext.Provider>
    )
}