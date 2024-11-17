import type {
    AdminNavbarPopupId
} from '../pages/admin/data';
import type {
    StudentNavbarPopupId
} from '../pages/home/data';

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
    id: AdminNavbarPopupId | StudentNavbarPopupId
    structure: "side" | "bubble"
}

// Type button will have onclick events
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