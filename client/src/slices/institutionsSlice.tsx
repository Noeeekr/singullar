import { createSlice } from '@reduxjs/toolkit'

import type { PayloadAction } from '@reduxjs/toolkit'
import type { InstitutionsState } from '../types/data';

interface IInstitutionSlice {
    institutions:    InstitutionsState[]
    current:       InstitutionsState | null
}

const inst: IInstitutionSlice = {
    institutions: [],
    current: null
};

const InstitutionsSlice = createSlice({
    name: "institutions",
    initialState: inst,
    reducers: {
        setInstitutions: (_state, action: PayloadAction<IInstitutionSlice>) => {
            return action.payload
        },
        addInstitutions: (state, action: PayloadAction<InstitutionsState[]>) => {
            state = { ...state, institutions: [ ...state.institutions, ...action.payload ] }
        },
        clearInstitutions: () => {
            return {
                institutions: [],
                current: null
            };
        },
    }
})

export default InstitutionsSlice.reducer;

export const { 
    clearInstitutions, 
    addInstitutions,
    setInstitutions,
} = InstitutionsSlice.actions; 