import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

import { IUser } from '../types/user';
import { AuthState } from '../types/general'

const initialState: AuthState = {
    user: null
}

const authSlice = createSlice({
    name: "auth",
    initialState,
    reducers: {
        updateUser: (state, action: PayloadAction<IUser | null>) => {
            state.user = action.payload
        }
    },
})

export default authSlice.reducer;

export const { updateUser } = authSlice.actions