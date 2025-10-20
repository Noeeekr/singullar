// Components
import FormSection from "./FormSection"
import ErrorBubble from "@components/bubbles/ErrorBubble/ErrorBubble";

// Hooks
import { useState, createContext } from 'react';
import { useForm } from 'react-hook-form';
import useContextAwareFetch from "@hooks/useContextAwareFetch";

// Models
import { SERVER_ADDR } from "../../../../../configs";

import type { DefaultResponse } from "@models/server/server";
import type { UseFormSetValue } from 'react-hook-form';
import type { Dispatch, SetStateAction } from 'react';
import type { FormFirstSectionData } from "./FormFirstSection";
import type { FormSecondSectionData } from "./FormSecondSection";
import type { FormThirdSectionData } from "./FormThirdSection";
import type { FormFourthSectionData } from "./FormFourthSection";
import { Typography } from "@mui/material";

export interface FormSections {
    firstSection: FormFirstSectionData,
    secondSection: FormSecondSectionData,
    thirdSection: FormThirdSectionData,
    fourthSection: FormFourthSectionData,
}
export interface FormContext {
    send: (body: FormRequest) => void,
    setFormData: UseFormSetValue<FormSections>,
    setActiveSection: Dispatch<SetStateAction<1 | 2 | 3 | 4>>,
    formSections: FormSections,
}
export type FormRequest = FormFirstSectionData & FormSecondSectionData & FormThirdSectionData & FormFourthSectionData

export const FormContext = createContext<FormContext>({} as FormContext);

/**
 * Form to create classes. 
 * 
 * This component also has a context provided by the same file that can help manage its inputs, their values, and their state.
 */
const FormCreateClass = (): JSX.Element => {
    const { response, isLoading, error, send } = useContextAwareFetch<DefaultResponse<string>, FormRequest>(
        `${SERVER_ADDR}/api/class/create`,
        {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            cache: "no-cache"
        }
    )
    const { setValue, watch } = useForm<FormSections>({
        defaultValues: {
            firstSection: {
                className: "",
                series: "",
                segment: "",
            },
            secondSection: {
                students: [],
            },
            fourthSection: {
                materials: []
            }
        }
    });
    const [activeSection, setActiveSection] = useState<1 | 2 | 3 | 4>(1)

    return (
        <>
        <form>
            <FormContext.Provider value={{
                setActiveSection: setActiveSection,
                setFormData: setValue,
                formSections: watch(),
                send: send,
            }}>
                <FormSection activeSection={activeSection}/>
            </FormContext.Provider>
        </form>
        {
            isLoading 
            ? <Typography textAlign="center" variant="body1" fontWeight="bold">Criando turma...</Typography>
            : <></>
        }
        {
            error
            ? <ErrorBubble err={error}/>
            : <></>
        }
        {
            response
            ? <Typography textAlign="center" variant="body1" fontWeight="bold">Turma criada</Typography>
            : <></>
        }
        </>
    )
}

export default FormCreateClass;