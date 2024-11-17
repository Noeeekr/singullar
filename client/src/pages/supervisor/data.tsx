import type { 
    ISideMenuItems, 
    // ISideMenuLinkButton, 
    // ISideMenuButton
} from '../../types/buttonProps'

export const supervisor_sidemenu_data: ISideMenuItems[] = [

];

export type SupervisorNavbarPopupId = "deleteLater" | "";

export interface INavbarPopupData {
    structure: "side" | "bubble",
    title: string,
    id: SupervisorNavbarPopupId,
    icon: JSX.Element,
    content: JSX.Element,
}

export const supervisor_navbar_popup_data: INavbarPopupData[] = [
    {
        structure: "side",
        title: "deleteLater",
        id: "deleteLater",
        icon: <span>D</span>,
        content: <div>Memeteur Vivem</div>,
    },
];