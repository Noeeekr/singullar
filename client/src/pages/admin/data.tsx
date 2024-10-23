// Icons

import { LuScanFace,LuBarChart3 } from 'react-icons/lu'
import { VscAccount } from "react-icons/vsc";
import { BiDirections } from "react-icons/bi";
import { GoPerson, GoLock } from "react-icons/go";
import { MdOutlineNotificationsNone } from "react-icons/md";
import { TbSmartHome, TbSpeakerphone } from "react-icons/tb";
import { FaChalkboardTeacher } from "react-icons/fa";
import { IoSchool, IoPeopleSharp, IoExitOutline } from 'react-icons/io5';
import { IoIosHelpCircleOutline, IoIosNotificationsOutline } from "react-icons/io";

// components;

import NotificationPopupContent from '../../components/NotificationPopupContent'
import MyAccountPopupContent from '../../components/MyAccountPopupContent'

// Types

import type { 
    ISideMenuItems, 
    ISideMenuLinkButton, 
    ISideMenuButton,
} from '../../types/propsButtons'

export const admin_sidemenu_data: ISideMenuItems[] = [
    {
        title: "Principal",
        items: [
            {
                title: "Inicio",
                type: "link",
                href: "/admin",
                icon: <TbSmartHome/>,
            },
            {
                title: "Notificações",
                icon: <MdOutlineNotificationsNone />, // might need the other version for hover effect
                type: "popup",
                content: <NotificationPopupContent />,
            },
            {
                title: "Ajuda",
                icon: <BiDirections />,
                type: "popup",
                content: <NotificationPopupContent />,
            },
            {
                title: "Minha conta",
                icon: <GoPerson />,
                type: "group",
                items: []
            },
        ]
    },
    {
        title: "Gerenciar",
        items: [
            {
                title: "Minhas turmas",
                type: "link",
                href: "/admin/classes",
                icon: <IoSchool/>,
            },
            {
                title: "Meus professores",
                type: "link",
                href: "/admin/teachers",
                icon: <FaChalkboardTeacher/>,
            },
            {
                title: "Meus alunos",
                type: "link",
                href: "/admin/students",
                icon: <IoPeopleSharp/>,
            }
        ]
    }
];

const admin_sidemenu_data_myaccount: (ISideMenuLinkButton | ISideMenuButton)[] = [
    {
        title: "Dados pessoais e acesso",
        icon: <LuScanFace />,
        type: "link",
        href: "/",
    },
    {
        title: "Comunicações",
        icon: <TbSpeakerphone />,
        type: "link",
        href: "/",
    },
    {
        title: "Privacidade",
        icon: <GoLock />,
        type: "link",
        href: "/",
    },
    {
        title: "Sair",
        icon: <IoExitOutline />,
        type: "button",
    },
];
/**
 * This must reflect all the id's fields present in navbar pop-up data 
 **/ 
export type AdminNavbarPopupId = "notifications" | "schooldata" | "myaccount" | "" | "help";

export interface INavbarPopupData {
    structure: "side" | "bubble",
    title: string,
    id: AdminNavbarPopupId,
    icon: JSX.Element,
    content: JSX.Element,
}

export const admin_navbar_popup_data: INavbarPopupData[] = [
    {
        structure: "side",
        title: "Dados escolares",
        id: "schooldata",
        icon: <LuBarChart3 color="white" fontSize={22}/>,
        content: <div>Memeteur Vivem</div>,
    },
    {
        structure:"side",
        title:"Notificações",
        id:"notifications",
        icon: <IoIosNotificationsOutline color="white" fontSize={26} />,
        content: <NotificationPopupContent/>,
    },
    {
        structure: "bubble",
        title:"Minha conta",
        id:"myaccount",
        icon: <VscAccount color="white" fontSize={22} />,
        content: <MyAccountPopupContent items={admin_sidemenu_data_myaccount} />,
    },
    {
        structure:"side",
        title:"Central de ajuda",
        id:"help",
        icon: <IoIosHelpCircleOutline color="white" fontSize={26} />,
        content: <div>Lorem Ipsum</div>,
    },
];