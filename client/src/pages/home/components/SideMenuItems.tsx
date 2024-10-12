import {
    Stack,
    styled,
    StackProps,
} from '@mui/material'

import {
    ISideMenuLinkButton,
    ISideMenuPopupButton,
    ISideMenuLinkButtonGroup,
} from '../../../types/propsButtons'

import PopupButton from './SidePopupButton'
import LinkButtonGroup from './LinkButtonGroup'
import LinkButton, { ILinkButtonBaseProps } from './LinkButton'

interface ISideMenuItemsProps extends ILinkButtonBaseProps {
    items: (ISideMenuLinkButton | ISideMenuLinkButtonGroup | ISideMenuPopupButton)[],

    isCompacted?: boolean,
    gap?: number,
}

/**
 * Holds the icon, the text, and the arrow of each Link / Group
 */
const MenuItemStack = styled(({ children, onClick, ...props }: StackProps) => (
    <Stack direction="row" component="li" onClick={onClick} {...props}>{ children }</Stack>
))(({ theme }) => ({
    alignItems: "center",
    overflow: 'hidden',
    textWrap: 'nowrap',
    cursor: "pointer",
    color: "black",
    borderRadius: 6,
    "&:hover > .MuiBox-root": {
        opacity: 1,
        transition: 'opacity 200ms ease-in-out'
    },
    "&:active": {
        backgroundColor: theme.palette.primary.purpleLightInv
    },
}))

const SideMenuItems = ({ items, ...props}: ISideMenuItemsProps): JSX.Element => {
    return (
        <Stack component="ul" sx={{ WebkitUserSelect: 'none', userSelect: 'none', msUserSelect: 'none' }} gap={props.gap ? props.gap : 0.5}>
            {
                items.map((item) => {
                    switch(item.type) {
                        case "link":
                            return <LinkButton key={item.title} {...props} {...item} />;
                        case "popup":
                            return <PopupButton key={item.title} hasNotifications={item?.notifications ? item.notifications : false} {...props} {...item} />;
                        case "group":
                            return <LinkButtonGroup key={item.title} {...props} {...item} />;
                    }
                })
            }
        </Stack>
    )
}

export { MenuItemStack }

export default SideMenuItems
