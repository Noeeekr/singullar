// Icons
import { BiDirections } from "react-icons/bi";
import { LuScanFace,LuBarChart3 } from 'react-icons/lu'
import { MdOutlineNotificationsNone } from "react-icons/md";
import { FaChalkboardTeacher, FaEdit } from "react-icons/fa";
import { IoSchool, IoPeopleSharp, IoExitOutline } from 'react-icons/io5';
import { GoPerson, GoLock, GoDiscussionDuplicate } from "react-icons/go";
import { TbSmartHome,  TbSpeakerphone, TbGridDots } from "react-icons/tb";

// components;
import NotificationPopupContent from '../../components/NotificationPopupContent'
import MyAccountPopupContent from '../../components/MyAccountPopupContent'
import QuickAccessPopupContent from '../../components/QuickAccessPopupContent'

// Types
import type { 
    ISideMenuItems, 
    ISideMenuButton,
    ISideMenuLinkButton, 
    ISideMenuPopupButton,
} from '../../types/propsButtons'

// Data
import { 
    sidemenu_data_classroom,
    sidemenu_data_quickaccess,
    sidemenu_data_utilities,
} from '../data'
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

export const admin_navbar_popup_data: ISideMenuPopupButton[] = [
    {
        id:"notifications",
        title:"Notificações",
        structure:"side",
        type: "popup",
        icon: <MdOutlineNotificationsNone color="white" fontSize={22} />,
        content: <NotificationPopupContent/>,
    },
    {
        id:"myaccount",
        title:"Minha conta",
        structure: "bubble",
        type: "popup",
        icon: <GoPerson color="white" fontSize={22} />,
        content: <MyAccountPopupContent items={admin_sidemenu_data_myaccount} />,
    },
    {
        id:"help",
        type: "popup",
        structure:"side",
        title:"Central de ajuda",
        content: <div>Lorem Ipsum</div>,
        icon: <BiDirections color="white" fontSize={20}/>,
    },
    {
        id: "quickaccess",
        title: "Acesso rápido",
        structure: "bubble",
        type: "popup",
        icon: <TbGridDots color="white" fontSize={23} />,
        content: <QuickAccessPopupContent items={sidemenu_data_quickaccess.items}/>,
    },
];

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
            ...admin_navbar_popup_data,
        ]
    },
    {
        title: "Ferramentas",
        items: [
            {
                title: "Dados escolares",
                type: "link",
                icon: <LuBarChart3 color="white" fontSize={22}/>,
                href: "/",
            },
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
            },
        ]
    },
    {
        title: "Plataforma",
        items: [
            {
                title: "Banco de questões",
                type: "link",
                href: "/admin/students",
                icon: <FaEdit/>,
            },
            {
                title: "Estudo orientado",
                type: "link",
                href: "/admin/students",
                icon: <GoDiscussionDuplicate/>,
            },
            ...sidemenu_data_classroom.items,
        ]
    },
    sidemenu_data_utilities,
];

