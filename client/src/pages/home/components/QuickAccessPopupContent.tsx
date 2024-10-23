import SectionTitle from '../../../components/SectionTitle'
import SideMenuItems from '../../../components/SideMenuItems'

import type { ISideMenuLinkProps } from '../../../components/LinkButton'
import type { ISideMenuButtonProps } from '../../../components/Button'

import {
    Box,
    Stack,
} from '@mui/material'

const QuickAccessPopupContent = (props: { items: (ISideMenuLinkProps | ISideMenuButtonProps)[] }) => {
    const { items } = props;

    return (
        <Box
            sx={{
                backgroundColor: 'rgb(220,220,220,0.2)',
                borderRadius: 4,
                padding: 2,
            }}
        >
            <Stack gap={1}>
                <SectionTitle>
                    Ir para
                </SectionTitle>
                <SideMenuItems
                    showIcon={true}
                    showDescription={true}
                    items={items}
                />
            </Stack>
        </Box>
    )
}

export default QuickAccessPopupContent;