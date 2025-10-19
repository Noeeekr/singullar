import { SideMenuButtons } from "../SideMenuButtons";
import { NavigationPopupProps } from "@components/layout/navbar/AppNavbar";

export const teacher_sidemenu_data: SideMenuButtons[] = [

];

export type TeacherNavbarPopupId = "deleteLater" | "";

export const teacher_navbar_popup_data: NavigationPopupProps[] = [
    {
        variant: "side",
        title: "deleteLater",
        id: "deleteLater",
        icon: {
            component: <span>D</span>
        },
        element: <div>Memeteur Vivem</div>,
        type: "popup"
    },
];