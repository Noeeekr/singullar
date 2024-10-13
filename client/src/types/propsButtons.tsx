export interface ISideMenuButtonBase {
    title: string,
    icon: JSX.Element,
}

export interface ISideMenuLinkButton extends ISideMenuButtonBase {
    href: string
    type: "link"
    description?: string 
}

export interface ISideMenuPopupButton extends ISideMenuButtonBase {
    type: "popup"
    content: JSX.Element
    notifications?: boolean
}

// type button will have onclick events
export interface ISideMenuButton extends ISideMenuButtonBase {
    type: "button"
}

export interface ISideMenuButtonGroup extends ISideMenuButtonBase {
    items: (ISideMenuLinkButton | ISideMenuButton)[]
    type: "group"
}

export interface ISideMenuItems {
    title: string,
    items: (ISideMenuLinkButton | ISideMenuButtonGroup | ISideMenuPopupButton | ISideMenuButton)[]
}