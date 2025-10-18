import type { ISideMenuItems } from '@models/buttonProps'

export const teacher_sidemenu_data: ISideMenuItems[] = [

];

export type TeacherNavbarPopupId = "deleteLater" | "";

export interface INavbarPopupData {
    structure: "side" | "bubble",
    title: string,
    id: TeacherNavbarPopupId,
    icon: JSX.Element,
    content: JSX.Element,
}

export const teacher_navbar_popup_data: INavbarPopupData[] = [
    {
        structure: "side",
        title: "deleteLater",
        id: "deleteLater",
        icon: <span>D</span>,
        content: <div>Memeteur Vivem</div>,
    },
];