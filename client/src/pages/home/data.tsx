// Icons
import { LuScanFace } from "react-icons/lu";
import { VscAccount } from "react-icons/vsc";
import { BiDirections } from "react-icons/bi";
import { MdOutlineNotificationsNone } from "react-icons/md";
import { GoPerson, GoLock } from "react-icons/go";
import { TbSmartHome,  TbSpeakerphone, TbGridDots } from "react-icons/tb";
import { IoIosHelpCircleOutline, IoIosNotificationsOutline } from "react-icons/io";
import { IoPeopleOutline, IoExitOutline, IoBarcodeOutline } from "react-icons/io5";

// components;
import NotificationPopupContent from '../../components/NotificationPopupContent'
import MyAccountPopupContent from '../../components/MyAccountPopupContent'
import QuickAccessPopupContent from '../../components/QuickAccessPopupContent'

// Types
import { 
    ISideMenuItems, 
    ISideMenuLinkButton, 
    ISideMenuButton 
} from '../../types/propsButtons'

// Data
import { 
    sidemenu_data_classroom,
    sidemenu_data_quickaccess,
} from '../data';

const students_sidemenu_data_myaccount: (ISideMenuLinkButton | ISideMenuButton)[] = [
    {
        title: "Dados pessoais e acesso",
        icon: <LuScanFace />,
        type: "link",
        href: "/",
    },
    {
        title: "Responsáveis vinculados",
        icon: <IoPeopleOutline />,
        type: "link",
        href: "/",
    },
    {
        title: "Código de acesso",
        icon: <IoBarcodeOutline />,
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
const students_sidemenu_data_main: ISideMenuItems = {
    title: "Principal",
    items: [
        {
            title: "Início",
            icon: <TbSmartHome />,
            type: "link",
            href: "/home",
        },
        {
            title: "Notificações",
            icon: <MdOutlineNotificationsNone />,
            type: "popup",
            content: <NotificationPopupContent />,
            structure: "side",
            id: "notifications",
        },
        {
            title: "Ajuda",
            icon: <BiDirections />,
            type: "popup",
            content: <NotificationPopupContent />,
            structure: "side",
            id: "help",
        },
        {
            title: "Minha conta",
            icon: <GoPerson />,
            type: "group",
            items: students_sidemenu_data_myaccount
        },
    ],
};

export const students_sidemenu_data: ISideMenuItems[] = [
    students_sidemenu_data_main,
    sidemenu_data_classroom,
    sidemenu_data_quickaccess,
]

export type StudentNavbarPopupId = "help" | "notifications" | "myaccount" | "quickaccess" | "";

export interface INavbarPopupData {
    structure: "side" | "bubble",
    title: string,
    id: StudentNavbarPopupId,
    icon: JSX.Element,
    content: JSX.Element,
}

export const students_navbar_popup_data: INavbarPopupData[] = [
    {
        structure:"side",
        title:"Central de ajuda",
        id:"help",
        icon: <IoIosHelpCircleOutline color="white" fontSize={26} />,
        content: <div>Lorem Ipsum</div>,
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
        content: <MyAccountPopupContent items={students_sidemenu_data_myaccount} />,
    },
    {
        structure: "bubble",
        title:"Acesso Rápido",
        id:"quickaccess",
        icon: <TbGridDots color="white" fontSize={23} />,
        content: <QuickAccessPopupContent items={sidemenu_data_quickaccess.items} />,
    }
]
