import MultiStepForm from "@components/forms/MultiStepForm"
import FirstSection, { FirstSectionContext, FirstSectionContextProps } from "./FirstSection"
import SecondSection, { SecondSectionContext, SecondSectionContextProps } from "./SecondSection"
import useContextAwareFetch, { defaultResponse } from "@hooks/useContextAwareFetch"
import { SERVER_ADDR } from "../../../../../../configs"
import { useCallback } from "react"

const sections = [
    {
        content: FirstSection,

        title: "Informações Básicas",

        context: FirstSectionContext,
        defaultValues: {
            question_title: undefined as string | undefined,
            question_difficulty_level: undefined as number | undefined,
        }
    },
    {
        content: SecondSection,

        title: "Conteúdo da questão",

        context: SecondSectionContext,
        defaultValues: {
            question_short_description: undefined as undefined | string,
            question_description: undefined as undefined | string,
            alternatives: [] as string[] | undefined
        }
    },
]

type FormData = FirstSectionContextProps & SecondSectionContextProps

export default function (): JSX.Element {
    const { response, send, isLoading, error } = useContextAwareFetch<null, FormData>(
        `${SERVER_ADDR}/api/question/create`,
        { 
            ...defaultResponse,
            method: "POST",
        },
    )

    const onSubmit = useCallback((formData: FirstSectionContextProps & SecondSectionContextProps) => {
        send(formData)
    }, [send])  

    return (
        <>
            <MultiStepForm sections={sections} onSubmit={onSubmit} />
            { response }
            Isloading { isLoading }
            error { error }
        </>
    )
}