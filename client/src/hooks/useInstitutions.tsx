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

const useInstitutions = (): ([InstitutionsState[] | [], Function]) => {
    const insts = useAppSelector((store) => store.institutions);
    const [institutions, setInstitutions] = useState<InstitutionsState[]>(insts)

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
                    dispatch(actionSetInstitutions([])).payload
                );
                return;
            }         
            setInstitutions(
                dispatch(actionSetInstitutions(res.data)).payload
            );
        });
    },[])

    return [institutions, getInstitutions];
};

export default useInstitutions;