// Components
import { Divider, Stack, StackProps, useMediaQuery, useTheme } from '@mui/material'
import type { PopupButtonProps } from '../../buttons/Popup/Popup'
import type { ButtonGroupProps } from '../../buttons/ButtonGroup'
import type { LinkButtonProps } from '../../buttons/Link'

import ButtonGroup from '../../buttons/ButtonGroup'
import { Fragment } from 'react/jsx-runtime'
import SectionTitle from '@components/headers/sectionHeader/SectionTitle'
import LinkButton from '../../buttons/Link'
import PopupButton from '../../buttons/Popup/Popup'
import Button, { ButtonEffectsProps, ButtonsProps } from '@components/buttons/Default'

export interface SideMenuButtons extends ButtonGroupProps { }

export interface SideMenuButtonsProps extends StackProps {
    isOpen?: boolean,
    showDescription?: boolean
    menus: (LinkButtonProps | ButtonsProps | ButtonGroupProps | PopupButtonProps)[]
    effects?: ButtonEffectsProps
}

const SideMenuButtons = ({
    isOpen,
    title,
    menus,
    ...props
}: SideMenuButtonsProps): JSX.Element => {
    const theme = useTheme()
    const isMobile = useMediaQuery(theme.breakpoints.down("sm"))

    return (
        <Fragment>
            {Boolean(title) && !isOpen && <Divider />}
            <Stack
                component="ul"
                sx={{ WebkitUserSelect: 'none', userSelect: 'none', msUserSelect: 'none' }}
                gap={props.gap ? props.gap : 0.5}
            >
                {isOpen && title && <SectionTitle sx={{ fontSize: 14 }}>{title}</SectionTitle>}
                {
                    menus && menus.map((item) => {
                        item.isMobile = isMobile
                        switch (item.type) {
                            case "link":
                                return <LinkButton key={item.title} {...item as LinkButtonProps} />;
                            case "popup":
                                return <PopupButton key={item.title} {...item as PopupButtonProps} iconVariant="large" />;
                            case "group":
                                return <ButtonGroup key={item.title} isOpen={isOpen} {...item as ButtonGroupProps} />;
                            case "button":
                                return <Button key={item.title} {...item} />;
                            default:
                                return <div>Error rendering button, type does not exist in ISideMenuItemsProps</div>
                        }
                    })
                }
            </Stack>
        </Fragment>
    )
}

export default SideMenuButtons
