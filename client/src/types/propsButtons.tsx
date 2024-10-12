export interface ISideMenuItemBase {
    title: string,
    icon: JSX.Element,
}

export interface ISideMenuLinkButton extends ISideMenuItemBase {
    href: string
    type: "link"
    description?: string 
}

export interface ISideMenuLinkButtonGroup extends ISideMenuItemBase {
    items: ISideMenuLinkButton[]
    type: "group"
}

export interface ISideMenuPopupButton extends ISideMenuItemBase {
    type: "popup"
    content: JSX.Element
    notifications?: boolean
}

export interface ISideMenuItems {
    title: string,
    items: (ISideMenuLinkButton | ISideMenuLinkButtonGroup | ISideMenuPopupButton)[]
}