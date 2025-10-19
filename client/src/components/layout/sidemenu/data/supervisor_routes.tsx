import { NavigationPopupProps } from "@components/layout/navbar/AppNavbar";
import { SideMenuButtons } from "../SideMenuButtons";

export const supervisor_sidemenu_data: SideMenuButtons[] = [

];

export type SupervisorNavbarPopupId = "deleteLater" | "";

export const supervisor_navbar_popup_data: NavigationPopupProps[] = [
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