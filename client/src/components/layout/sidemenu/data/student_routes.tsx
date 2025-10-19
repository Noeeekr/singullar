// Icons;
import { BiDirections } from 'react-icons/bi';
import { LuScanFace } from 'react-icons/lu';
import { MdOutlineNotificationsNone } from 'react-icons/md';
import { IoExitOutline, IoPeopleOutline, IoBarcodeOutline } from 'react-icons/io5';
import { GoPerson, GoLock } from 'react-icons/go';
import { TbSmartHome, TbSpeakerphone, TbGridDots } from 'react-icons/tb';
import { IoIosHelpCircleOutline, IoIosNotificationsOutline } from 'react-icons/io';
import { VscAccount } from 'react-icons/vsc';

// components;
import NotificationPopupContent from '../../../popups/NotificationPopupContent'
import MyAccountPopupContent from '../../../popups/PopupContentMyAccount'
import QuickAccessPopupContent from '../../../popups/PopupContentQuickAccess'

// Data
import {
    sidemenu_data_classroom,
    sidemenu_data_quickaccess,
} from './any_role_routes';

import type { LinkButtonProps } from '@components/buttons/ButtonLink';
import { SideMenuButtons } from '../SideMenuButtons';
import { ButtonsProps } from '@components/buttons/Button';
import { NavigationPopupProps } from '@components/layout/navbar/AppNavbar';

const students_sidemenu_data_myaccount: (LinkButtonProps | ButtonsProps)[] = [
    {
        title: "Dados pessoais e acesso",
        icon: {
            component: <LuScanFace />,
        },
        type: "link",
        href: "/",
    },
    {
        title: "Responsáveis vinculados",
        icon: {
            component: <IoPeopleOutline />,
        },
        type: "link",
        href: "/",
    },
    {
        title: "Código de acesso",
        icon: {
            component: <IoBarcodeOutline />,
        },
        type: "link",
        href: "/",
    },
    {
        title: "Comunicações",
        icon: {
            component: <TbSpeakerphone />,
        },
        type: "link",
        href: "/",
    },
    {
        title: "Privacidade",
        icon: {
            component: <GoLock />,
        },
        type: "link",
        href: "/",
    },
    {
        title: "Sair",
        icon: {
            component: <IoExitOutline />,
        },
        type: "button",
    },
];

const students_sidemenu_data_main: SideMenuButtons = {
    title: "Principal",
    menus: [
        {
            title: "Início",
            icon: {
                component: <TbSmartHome />,
            },
            type: "link",
            href: "/home",
        },
        {
            title: "Notificações",
            icon: {
                component: <MdOutlineNotificationsNone />,
            },
            type: "popup",
            element: <NotificationPopupContent />,
            variant: "side",
            id: "notifications",
        },
        {
            title: "Ajuda",
            icon: {
                component: <BiDirections />,
            },
            type: "popup",
            element: <NotificationPopupContent />,
            variant: "side",
            id: "help",
        },
        {
            title: "Minha conta",
            menus: students_sidemenu_data_myaccount,
            icon: {
                component: <GoPerson />,
            },
            type: "group",
        },
    ],
    type: "group",
};

export const students_sidemenu_data: SideMenuButtons[] = [
    students_sidemenu_data_main,
    sidemenu_data_classroom,
    sidemenu_data_quickaccess,
]

export type StudentNavbarPopupIds = "help" | "notifications" | "myaccount" | "quickaccess" | "";

export const students_navbar_popup_data: NavigationPopupProps[] = [
    {
        variant: "side",
        title: "Central de ajuda",
        id: "help",
        icon: {
            component: <IoIosHelpCircleOutline color="white" fontSize={26} />,
        },
        element: <div>Lorem Ipsum</div>,
        type: "popup",
    },
    {
        variant: "side",
        title: "Notificações",
        id: "notifications",
        icon: {
            component: <IoIosNotificationsOutline color="white" fontSize={26} />,
        },
        element: <NotificationPopupContent />,
        type: "popup",
    },
    {
        variant: "bubble",
        title: "Minha conta",
        id: "myaccount",
        icon: {
            component: <VscAccount color="white" fontSize={22} />,
        },
        element: <MyAccountPopupContent menus={students_sidemenu_data_myaccount} />,
        type: "popup",
    },
    {
        variant: "bubble",
        title: "Acesso Rápido",
        id: "quickaccess",
        icon: {
            component: <TbGridDots color="white" fontSize={23} />,
        },
        element: <QuickAccessPopupContent menus={sidemenu_data_quickaccess.menus} />,
        type: "popup",
    }
]
