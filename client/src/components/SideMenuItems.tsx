import {
    useMemo,
} from 'react'

// components
import {
    Stack,
    styled,
    StackProps,
} from '@mui/material'

import PopupButton from './SidePopupButton'
import LinkButtonGroup from './ButtonGroup'
import LinkButton from './LinkButton'
import Button from './Button'

// types
import type {
    ISideMenuLinkButton,
    ISideMenuButtonGroup,
    ISideMenuPopupButton,
    ISideMenuButton,
} from '../types/propsButtons'
import type {
    IButtonBaseProps
} from './Button'

interface ISideMenuItemsProps extends IButtonBaseProps {
    items: (ISideMenuLinkButton | ISideMenuButtonGroup | ISideMenuPopupButton | ISideMenuButton)[],

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
                            return <PopupButton key={item.title} notifications={item?.notifications ? item.notifications : false} {...props} {...item} />;
                        case "group":
                            return <LinkButtonGroup key={item.title} {...props} {...item} />;
                        case "button":
                            return <Button key={item.title} {...props} {...item}/>;
                        default: 
                            return <div>Error rendering button, type does not exist in ISideMenuItemsProps</div>
                    }
                })
            }
        </Stack>
    )
}

export { MenuItemStack }

export default SideMenuItems
