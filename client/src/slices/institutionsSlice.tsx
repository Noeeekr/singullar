import { createSlice } from '@reduxjs/toolkit'

import type { PayloadAction } from '@reduxjs/toolkit'
import type { InstitutionsState } from '../types/general';

const inst: InstitutionsState[] = [];

const InstitutionsSlice = createSlice({
    name: "institutions",
    initialState: inst,
    reducers: {
        setInstitutions: (state, action: PayloadAction<InstitutionsState[]>) => {
            return action.payload;
        },
        addInstitutions: (state, action: PayloadAction<InstitutionsState[]>) => {
            state.push(...action.payload)
        },
        clearInstitutions: (state) => {
            return []
        },
    }
})

export default InstitutionsSlice.reducer;

export const { 
    clearInstitutions, 
    addInstitutions,
    setInstitutions,
} = InstitutionsSlice.actions; 