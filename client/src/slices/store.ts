import { useSelector, useDispatch } from 'react-redux'
import { configureStore } from '@reduxjs/toolkit'

// REDUCERS

import authReducer from './authSlice'

const store = configureStore({
    reducer: {
        "auth": authReducer
    },
})

export default store;

export const useAppDispatch = useDispatch.withTypes<AppDispatch>()
export const useAppSelector = useSelector.withTypes<RootState>()

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch