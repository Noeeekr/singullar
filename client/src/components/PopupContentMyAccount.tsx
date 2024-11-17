import { useNavigate } from 'react-router-dom'
import type { MouseEventHandler } from 'react'

import {
    Box,
    Stack,
} from '@mui/material'

import useSignout from '../hooks/useSignout'

// COMPONENTS
import UserProfile from './UserProfile'
import Divider from './Divider'
import SideMenuItems from './SideMenuButtons'
import Button from './Button'

import { ISideMenuLinkProps } from './ButtonLink'
import { ISideMenuButtonProps } from './Button'

const MyAccountPopupContent = (props: { items: (ISideMenuLinkProps | ISideMenuButtonProps)[] }) => {
    const { items } = props;

    const navigate = useNavigate()

    const { signout, error } = useSignout()
    const handleSignout = async () => {
        await signout()

        if (!error) {
            navigate("/auth")
        }
    }

    return (
        <Stack>
            <UserProfile structure="center" />
            <Divider sx={{ marginX: 2 }} />
            <Box padding={1}>
                <SideMenuItems
                    gap={1}
                    fontWeight={400}
                    fontSize={13.5}
                    showIcon={true}
                    hasHoverEffect={true}
                    items={items.slice(0, -1)}
                />
            </Box>
            <Divider />
            <Box padding={1}>
                <div onClick={handleSignout as MouseEventHandler<HTMLDivElement>}>
                    <Button 
                        fontWeight={400}
                        fontSize={13.5}
                        showIcon={true}
                        hasHoverEffect={true}
                        { ...items[items.length - 1] }
                        type="button"
                    />
                </div>
            </Box>
        </Stack>
    )
}

export default MyAccountPopupContent