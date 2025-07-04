import {
    useState,
    useCallback,
} from 'react';

import {
    useAppDispatch,
    useAppSelector,
} from '../slices/store';
import {
    setInstitutions as actionSetInstitutions,
} from '../slices/institutionsSlice'

import type { InstitutionsState } from '../types/data';

const useInstitutions = (): [InstitutionsState[], () => void] => {
    const insts = useAppSelector((store) => store.institutions);
    const [institutions, setInstitutions] = useState<InstitutionsState[]>(insts.institutions)

    const dispatch = useAppDispatch();

    const getInstitutions = useCallback(() => {
        fetch("http://localhost:8000/api/user/institutions",{
            "headers": {
                "Content-Type": "application/json"
            },
            "credentials": "include",
        })
        .then(res => (res.json()))
        .then(res => {
            if (res.error) {
                setInstitutions(
                    dispatch(actionSetInstitutions({ current: null, institutions: [] })).payload.institutions
                );
                return;
            }         
            setInstitutions(
                dispatch(actionSetInstitutions({ current: null, institutions: res.data })).payload.institutions
            );
        });
    },[dispatch])

    return [institutions, getInstitutions];
};

export default useInstitutions;