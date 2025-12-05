import { PopupButtonProps } from '@components/buttons/Popup';
import { SideMenuButtons } from "../SideMenuButtons";

export const supervisor_sidemenu_data: SideMenuButtons[] = [

];

export type SupervisorNavbarPopupId = "deleteLater" | "";

export const supervisor_navbar_popup_data: PopupButtonProps[] = [
    {
        variant: "side",
        title: "deleteLater",
        id: "deleteLater",
        icon: {
            component: <span>D</span>,
        },
        element: <div>Memeteur Vivem</div>,
        type: "popup"
    },
];