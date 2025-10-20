// Features
import { useAppSelector } from '@slices/store';
import { createContext, useEffect } from "react"

// Components
import SearchFilters from "./components/SearchFilters"
import ClassBubble from "./components/ClassList"
import SectionHeader from "@components/headers/sectionHeader/SectionHeader"
import LinkButton from "@components/buttons/Link"
import Typography from '@mui/material/Typography';
import ErrorBubble from '@components/bubbles/ErrorBubble/ErrorBubble';

// Models
import type { SearchClassFilters } from '@hooks/useFetchClasses';
import useFetchClasses from '@hooks/useFetchClasses';

const FormContext = createContext<(body: SearchClassFilters[]) => void>((_) => { })

export default function Search(): JSX.Element {
    const institution = useAppSelector((store) => store.institution)

    const { response, isLoading, error, send } = useFetchClasses()


    useEffect(() => {
        send([])
    },[])

    return (
        <FormContext.Provider value={send}>
            <SectionHeader
                title="Turmas"
                subtitle={institution?.name || "Nome desconhecido"}
            >
                <div style={{ margin: '0px 0px 0px auto' }}>
                    <LinkButton
                        title=""
                        icon={{ component: <></> }}
                        type="link"
                        href="/admin/classes/create"
                    >
                        Cadastrar turmas
                    </LinkButton>
                </div>
            </SectionHeader>
            <SearchFilters send={send} />
            {
                error
                    ? <ErrorBubble err={error} />
                    : <></>
            }
            {
                isLoading
                    ? <Typography variant="body2" fontWeight="bold">Carregando turmas...</Typography>
                    : <></>
            }
            <ClassBubble classes={response} navegable={true}/>
        </FormContext.Provider>
    )
}