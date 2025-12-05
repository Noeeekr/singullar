import { SideMenuButtons } from "../SideMenuButtons";
import { PopupButtonProps } from '@components/buttons/Popup';

export const teacher_sidemenu_data: SideMenuButtons[] = [

];

export type TeacherNavbarPopupId = "deleteLater" | "";

export const teacher_navbar_popup_data: PopupButtonProps[] = [
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