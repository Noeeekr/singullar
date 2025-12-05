import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

import { User } from '@models/server';
import { UserState } from '@models/data'

const initialState: UserState = {
    user: null,
    mostVisitedUrls: {},
}

const userSlice = createSlice({
    name: "user",
    initialState: initialState,
    reducers: {
        updateUser: (state, action: PayloadAction<User | null>) => {
            state.user = action.payload
        },
        incrementUrlVisitedCount: (state, action: PayloadAction<string>) => {
            if (action.payload in state.mostVisitedUrls) {
                state.mostVisitedUrls[action.payload] = state.mostVisitedUrls[action.payload] + 1;
            } else {
                state.mostVisitedUrls[action.payload] = 1;
            }
        },
        resetUrlVisitedCount: (state) => {
            state.mostVisitedUrls = {};
        },
    },
})

export default userSlice.reducer;

export const { 
    updateUser,
    resetUrlVisitedCount,
    incrementUrlVisitedCount,
} = userSlice.actions