// Features
import useContextAwareFetch from '@hooks/useContextAwareFetch';
import { useAppSelector } from '@slices/store';
import { createContext } from "react"

// Components
import SearchFilters from "./components/SearchFilters"
import ClassBubble from "./components/ClassBubble"
import SectionHeader from "@components/SectionHeader"
import LinkButton from "@components/ButtonLink"
import Typography from '@mui/material/Typography';

// Models
import type { SearchClassFilters } from './components/SearchFilters';
import type { Class } from '@models/server';
import { SERVER_ADDR } from '../../../configs';
import ErrorBubble from '@components/ErrorBubble';

const FormContext = createContext<(body: SearchClassFilters[]) => void>((_) => { })

export default function Search(): JSX.Element {
    const institution = useAppSelector((store) => store.institution)

    const [classes, isLoading, error, send] = useContextAwareFetch<Class[], SearchClassFilters[]>(
        `${SERVER_ADDR}/api/classes`,
        {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json"
            }
        }
    )

    return (
        <FormContext.Provider value={send}>
            <SectionHeader
                title="Turmas"
                subtitle={institution?.name || "Nome desconhecido"}
            >
                <div style={{ margin: '0px 0px 0px auto' }}>
                    <LinkButton
                        title=""
                        icon={<></>}
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
            <ClassBubble classes={classes} />
        </FormContext.Provider>
    )
}