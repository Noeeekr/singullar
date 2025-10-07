import { createSlice, type PayloadAction } from '@reduxjs/toolkit'

import type { Institution } from '../models/server';

type InstitutionPayload = Institution | null

const InstitutionsSlice = createSlice({
    name: "institution",
    initialState: null as InstitutionPayload,
    reducers: {
        setInstitution: (_state, action: PayloadAction<InstitutionPayload>) => {
            return action.payload
        },
        clearInstitutions: () => {
            return null;
        },
    }
})

export default InstitutionsSlice.reducer;

export const { 
    clearInstitutions, 
    setInstitution,
} = InstitutionsSlice.actions; 