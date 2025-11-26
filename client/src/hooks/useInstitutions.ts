import {
    useState,
    useCallback,
} from 'react';
import {
    useNavigate
} from 'react-router-dom'
import {
    useAppDispatch,
    useAppSelector,
} from '../slices/store';
import { setInstitution as ActionSetInstitution } from '../slices/institutionsSlice'

import type { Institution } from '../models/server';

import { SERVER_ADDR } from '../configs'

const useInstitutions = (): [Institution | null, () => void] => {
    const i = useAppSelector((store) => store.institution);
    const [institution, setInstitution] = useState<Institution | null>(i)
    const navigate = useNavigate()
    const dispatch = useAppDispatch();

    const getInstitution = useCallback(() => {
        fetch(`${SERVER_ADDR}/api/institution`,{
            "headers": {
                "Content-Type": "application/json"
            },
            "credentials": "include",
        })
        .then(res => {
            if (res.status == 401) navigate("/auth");
            return res
        })
        .then(res => (res.json()))
        .then(res => {
            if (res.error) {
                setInstitution(
                    dispatch(ActionSetInstitution(null)).payload
                );
                return;
            }         
            setInstitution(
                dispatch(ActionSetInstitution(null)).payload
            );
        });
    },[dispatch, navigate])

    return [institution, getInstitution];
};

export default useInstitutions;