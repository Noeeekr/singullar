// Components
import {
    Stack,
    styled,
    StackProps,
    Divider,
} from '@mui/material'

import PopupButton from '../../buttons/ButtonSidePopup'
import LinkButtonGroup from '../../buttons/ButtonGroup'
import LinkButton from '../../buttons/ButtonLink'
import Button from '@components/buttons/Button'

// types
import type {
    ISideMenuLinkButton,
    ISideMenuButtonGroup,
    ISideMenuPopupButton,
    ISideMenuButton,
} from '../../../models/buttonProps'
import type {
    ButtonCoreProps
} from '@components/buttons/Button'
import SectionTitle from '@components/headers/sectionHeader/SectionTitle'
import { Fragment } from 'react/jsx-runtime'

interface ISideMenuItemsProps extends ButtonCoreProps {
    title: string

    items: (ISideMenuLinkButton | ISideMenuButtonGroup | ISideMenuPopupButton | ISideMenuButton)[],

    isOpen?: boolean,
    gap?: number,
}

/**
 * Holds the icon, the text, and the arrow of each Link / Group
 */
const MenuItemStack = styled(({ children, onClick, ...props }: StackProps) => (
    <Stack direction="row" component="li" onClick={onClick} {...props}>{children}</Stack>
))(({ theme }) => ({
    alignItems: "center",
    overflow: 'hidden',
    textWrap: 'nowrap',
    cursor: "pointer",
    color: "black",
    borderRadius: 6,
    "&:hover > .MuiBox-root": {
        opacity: 1,
        transition: 'opacity 200ms ease-in-out',
    },
    "&:hover > .MuiBox-root > .MuiBox-root": {
        opacity: 1,
        transition: 'opacity 200ms ease-in-out',
    },
    "&:active": {
        backgroundColor: theme.palette.primary.purpleLightInv
    },
}))

const SideMenuItems = ({ items, isOpen, title, ...props }: ISideMenuItemsProps): JSX.Element => {
    return (
        <Fragment>
            { Boolean(title) && !isOpen && <Divider /> }
            <Stack gap={1}>
                { isOpen && title && <SectionTitle>{title}</SectionTitle> }
                <Stack component="ul" sx={{ WebkitUserSelect: 'none', userSelect: 'none', msUserSelect: 'none' }} gap={props.gap ? props.gap : 0.5}>
                    {
                        items.map((item) => {
                            switch (item.type) {
                                case "link":
                                    return <LinkButton key={item.title} {...props} {...item} variant="button" />;
                                case "popup":
                                    return <PopupButton key={item.title} {...props} {...item} />;
                                case "group":
                                    return <LinkButtonGroup key={item.title} {...props} {...item} />;
                                case "button":
                                    return <Button key={item.title} {...props} {...item} />;
                                default:
                                    return <div>Error rendering button, type does not exist in ISideMenuItemsProps</div>
                            }
                        })
                    }
                </Stack>
            </Stack>
        </Fragment>

    )
}

export { MenuItemStack }

export default SideMenuItems
