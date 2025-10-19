import { LinkButtonProps } from '@components/buttons/ButtonLink';
import SectionTitle from '../headers/sectionHeader/SectionTitle'
import SideMenuItems, { SideMenuButtons } from '../layout/sidemenu/SideMenuButtons'

import {
    Box,
    Stack,
} from '@mui/material'

const QuickAccessPopupContent = (props: { menus: (LinkButtonProps | SideMenuButtons)[] }) => {
    const { menus } = props;

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
                    title=""
                    showIcon={true}
                    showDescription={true}
                    menus={menus}
                />
            </Stack>
        </Box>
    )
}

export default QuickAccessPopupContent;