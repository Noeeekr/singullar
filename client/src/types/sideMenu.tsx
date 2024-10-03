export interface ISideMenuItemBase {
    title: string,
    icon: JSX.Element,
}

export interface ISideMenuLink extends ISideMenuItemBase {
    href: string
    type: "link"
}

export interface ISideMenuGroup extends ISideMenuItemBase {
    items: ISideMenuLink[]
    type: "group"
}

export interface ISideMenuItems {
    title: string,
    items: (ISideMenuLink | ISideMenuGroup)[]
}